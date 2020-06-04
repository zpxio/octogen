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

func (i *Inventory) Add(id string, content string, rarity float64, tags map[string]string) {
	x := BuildToken(id, content, rarity, tags)

	i.dictionary[id] = append(i.dictionary[id], x)
	i.selectRange[id] = rarity + i.selectRange[id]
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

func (i *Inventory) PickValue(selector *Selector, offset float64) string {
	taggedList, selectRange := i.getTokens(selector)

	// Pick the first value that exceeds the offset value
	selectValue := offset * selectRange
	var lastContent = ""

	for len(taggedList) > 0 && selectValue >= 0 {
		top := taggedList[0]
		taggedList = taggedList[1:]

		lastContent = top.Content
		selectValue -= top.Rarity
	}

	return lastContent
}
