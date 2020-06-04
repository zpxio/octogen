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
	"github.com/zpxio/octogen/rng"
	"regexp"
	"strings"
)

const RoundsMax = uint8(30)

var selectorRegex *regexp.Regexp

func init() {
	selectorRegex = regexp.MustCompile(`\[(\w+)(:((,?(\w+(!?=\w+)?)?)+))?]`)
}

type Template struct {
	working string
	rounds  uint8
}

func InitTemplate(t string) *Template {
	return &Template{
		working: t,
		rounds:  RoundsMax,
	}
}

func (t *Template) replaceNextToken(i *Inventory, offset float64) bool {
	// Find the next token
	matches := selectorRegex.FindStringSubmatch(t.working)

	if matches == nil {
		return false
	}

	log.Infof("Found tokens: %#v", matches)
	fullMatch := matches[0]
	selectorId := matches[1]
	selectorOptions := matches[3]

	selector := ParseSelector(selectorId, selectorOptions)

	tv := i.PickValue(selector, offset)

	t.working = strings.Replace(t.working, fullMatch, tv, 1)
	log.Infof("Working value is now: %s", t.working)

	return true
}

func (t *Template) Render(i *Inventory, source rng.RandomSource) string {
	replaced := true

	// Replace until there's nothing left to replace or you exceed the replacement count
	for replaced && t.rounds > 0 {
		replaced = t.replaceNextToken(i, source.Next())
	}

	return t.working
}
