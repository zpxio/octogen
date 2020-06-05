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
	"github.com/stretchr/testify/suite"
	"testing"
)

type StateSuite struct {
	suite.Suite
}

func TestStateSuite(t *testing.T) {
	suite.Run(t, new(StateSuite))
}

func (s *StateSuite) TestCreateState() {
	x := CreateState()

	s.Empty(x.Vars)
}

func (s *StateSuite) TestSetVars() {
	x := CreateState()

	vars := map[string]string{"A": "6", "B": "2"}
	x.SetVars(vars)

	s.NotEmpty(x.Vars)
	s.Equal("6", x.Vars["A"])
	s.Equal("2", x.Vars["B"])
	s.Len(x.Vars, 2)
}
