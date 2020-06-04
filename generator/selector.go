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
	"regexp"
)

var optRegex *regexp.Regexp

const (
	OptId    = 1
	OptValue = 4
	OptType  = 3

	OptTypeRequire = "="
	OptTypeExclude = "!="
	OptTypeExists  = ""
)

func init() {
	optRegex = regexp.MustCompile(`(\w+)((!?=)(\w+))?`)
}

type Selector struct {
	Id      string
	Require map[string]string
	Exclude map[string]string
	Exists  map[string]bool
}

func ParseSelector(id string, options string) *Selector {
	s := Selector{
		Id:      id,
		Require: make(map[string]string),
		Exclude: make(map[string]string),
		Exists:  make(map[string]bool),
	}

	// Parse the options
	//optParts := strings.Split(options, ",")

	parseGroups := optRegex.FindAllStringSubmatch(options, -1)

	log.Infof("Parsed Selector: %#v", parseGroups)

	for _, group := range parseGroups {
		switch group[OptType] {
		case OptTypeExists:
			s.Exists[group[OptId]] = true
			break
		case OptTypeRequire:
			s.Require[group[OptId]] = group[OptValue]
			break
		case OptTypeExclude:
			s.Exclude[group[OptId]] = group[OptValue]
			break
		}
	}

	return &s
}

func (s *Selector) IsSimple() bool {
	if len(s.Require) > 0 {
		return false
	} else if len(s.Exclude) > 0 {
		return false
	} else if len(s.Exists) > 0 {
		return false
	}

	return true
}

func (s *Selector) MatchesToken(t *Token) bool {
	// Check Require
	for k, v := range s.Require {
		if v != t.Tags[k] {
			return false
		}
	}

	// Check Exclude
	for k, v := range s.Exclude {
		tv, exists := t.Tags[k]
		if exists && tv == v {
			return false
		}
	}

	// Check Exists
	for k, _ := range s.Exists {
		_, exists := t.Tags[k]
		if !exists {
			return false
		}
	}

	return true
}
