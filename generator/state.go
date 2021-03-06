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

import "github.com/apex/log"

// State includes configurable state that can be used to configure individual executions of
// a Generator.
type State struct {
	Vars map[string]string
}

// CreateState builds a new, empty State.
func CreateState() *State {
	return &State{
		Vars: make(map[string]string),
	}
}

// SetVars sets the values of one or more variables based on the supplied map.
func (s *State) SetVars(v map[string]string) {
	for varName, val := range v {
		s.Vars[varName] = val
		log.Infof("Set variable: %s=%s", varName, val)
	}
}
