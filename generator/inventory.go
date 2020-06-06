/*
 * Copyright 2020 zpxio (Jeff Sharpe)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import (
	"github.com/apex/log"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Inventory struct {
	dictionary  map[string][]Token
	selectRange map[string]float64
}

func CreateInventory() *Inventory {
	i := Inventory{
		dictionary:  make(map[string][]Token),
		selectRange: make(map[string]float64),
	}

	return &i
}

func (i *Inventory) AddToken(id string, content string, rarity float64, tags map[string]string) *Token {
	x := BuildToken(id, content, rarity, tags)

	i.Add(x)

	return &x
}

func (i *Inventory) Add(t Token) *Token {
	i.dictionary[t.Id] = append(i.dictionary[t.Id], t)
	i.selectRange[t.Id] += t.Rarity

	return &t
}

func (i *Inventory) getTokens(selector *Selector) ([]Token, float64) {
	idList, idFound := i.dictionary[selector.Id]

	if !idFound {
		return []Token{}, 0.0
	}

	var taggedList []Token
	selectRange := 0.0

	if selector.IsSimple() {
		return idList, i.selectRange[selector.Id]
	}

	for _, x := range idList {
		if selector.MatchesToken(&x) {
			taggedList = append(taggedList, x)
			selectRange += x.Rarity
		}
	}

	return taggedList, selectRange
}

func (i *Inventory) Pick(selector *Selector, offset float64) *Token {
	taggedList, selectRange := i.getTokens(selector)
	log.Infof("Selected tokens: %#v", taggedList)

	// Pick the first value that exceeds the offset value
	selectValue := offset * selectRange
	log.Infof("Selecting Weight Index: %f", selectValue)
	var lastToken *Token = nil

	for len(taggedList) > 0 && selectValue >= 0 {
		lastToken = &taggedList[0]
		taggedList = taggedList[1:]

		selectValue -= lastToken.Rarity
		log.Infof("Now selecting up to: %f", selectValue)
	}

	return lastToken
}

func (i *Inventory) Load(path string) error {
	// Read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "Failed to read inventory file.")
	}

	// Parse the YAML tokens
	var tokens []Token
	err = yaml.Unmarshal(data, &tokens)
	if err != nil {
		return errors.Wrap(err, "Failed to parse yaml file")
	}

	// Add all the tokens
	for _, t := range tokens {
		t.Normalize()

		if t.IsValid() {
			i.Add(t)
		}
	}

	return nil
}
