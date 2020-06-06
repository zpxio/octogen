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

// The maximum number of replacement rounds allowed during Rendering.
const RoundsMax = uint8(30)

var selectorRegex *regexp.Regexp
var varRegex *regexp.Regexp

func init() {
	selectorRegex = regexp.MustCompile(`\[(\w+)(:((,?(\w+(!?=\w+)?)?)+))?]`)
	varRegex = regexp.MustCompile(`\[\$(\w+)]`)
}

// replaceNextToken replaces the first complete token found
func replaceNextToken(working string, i *Inventory, s *State, source rng.RandomSource) (string, bool) {
	// Find the next token
	matches := selectorRegex.FindStringSubmatch(working)

	if matches == nil {
		return working, false
	}

	log.Infof("Found tokens: %#v", matches)
	fullMatch := matches[0]
	selectorId := matches[1]
	selectorOptions := matches[3]

	selector := ParseSelector(selectorId, selectorOptions)

	tv := i.Pick(selector, source.Next())
	working = strings.Replace(working, fullMatch, tv.Content, 1)
	s.SetVars(tv.SetVars)

	log.Infof("Working value is now: %s", working)

	return working, true
}

// replaceNextVar replaces the next variable found
func replaceNextVar(working string, s *State) (string, bool) {
	matches := varRegex.FindAllStringSubmatch(working, 20)

	if matches == nil {
		log.Infof("Found no variable matches")
		return working, false
	}

	for _, m := range matches {
		tag := m[0]
		varName := m[1]

		val := s.Vars[varName]
		log.Infof("Found Var reference: %s=%s", varName, val)
		if val != "" {
			working = strings.Replace(working, tag, val, 1)
			return working, true
		}
	}

	return working, false
}

// Render generates output from the supplied instruction string using the Inventory, State and RandomSource.
// The instructions are rendered by replacing one element at a time, selecting the first complete Token or
// first complete Variable found (in that order). Tokens or Variables which include other Token or Variable
// references are considered invalid/incomplete and will be skipped.
func Render(instruction string, i *Inventory, state *State, source rng.RandomSource) string {
	var replaced bool
	rounds := RoundsMax
	working := instruction

	// Keep trying until there aren't changes or all the rounds are expended
	for rounds > 0 {
		// Try to replace tokens
		working, replaced = replaceNextToken(working, i, state, source)

		// Try to replace variables if no tokens were replaced
		if !replaced {
			working, replaced = replaceNextVar(working, state)
		}

		if !replaced {
			rounds = 0
		} else {
			rounds -= 1
		}
	}

	return working
}
