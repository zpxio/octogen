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

type SelectorSuite struct {
	suite.Suite
}

func TestSelectorSuite(t *testing.T) {
	suite.Run(t, new(SelectorSuite))
}

func (s *SelectorSuite) TestParseSelector_Simple() {

	testId := "animal"

	x := ParseSelector(testId, "type=mammal,env!=water,enabled")

	s.NotNil(x)
	s.Equal(testId, x.Id)

	s.Contains(x.Require, "type")
	s.Equal("mammal", x.Require["type"])
	s.Contains(x.Exclude, "env")
	s.Equal("water", x.Exclude["env"])
	s.Contains(x.Exists, "enabled")
	s.True(x.Exists["enabled"])
}

func (s *SelectorSuite) TestIsSimple_Sure() {
	x := ParseSelector("test", "")

	s.True(x.IsSimple())
}

func (s *SelectorSuite) TestIsSimple_HasRequire() {
	x := ParseSelector("test", "x=y")

	s.False(x.IsSimple())
}

func (s *SelectorSuite) TestIsSimple_HasExclude() {
	x := ParseSelector("test", "x!=y")

	s.False(x.IsSimple())
}

func (s *SelectorSuite) TestIsSimple_HasExists() {
	x := ParseSelector("test", "x")

	s.False(x.IsSimple())
}

func (s *SelectorSuite) TestMatchesToken_IdOnly() {
	x := ParseSelector("animal", "")

	t := BuildToken("animal", "Gemsbok", 1.0, Tags{})
	s.True(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_SingleRequire() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "type=mammal")

	s.True(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_MultiRequire() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "type=mammal,family=antelope")

	s.True(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_SingleRequire_Mismatch() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "type=fish")

	s.False(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_MultiRequire_MixedMismatch() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "type=mammal,family=prosimian")

	s.False(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_SingleRestrict() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "type!=fish")

	s.True(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_MultiRestrict() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "type!=fish,family!=mustelid")

	s.True(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_SingleRestrict_Mismatch() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "type!=mammal")

	s.False(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_MultiRestrict_MixedMismatch() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "type!=fish,family!=antelope")

	s.False(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_SingleExists() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "extant")

	s.True(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_MultiExists() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "family,extant")

	s.True(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_SingleExists_Mismatch() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "prey")

	s.False(x.MatchesToken(&t))
}

func (s *SelectorSuite) TestMatchesToken_MultiExists_MixedMismatch() {
	t := BuildToken("animal", "Gemsbok", 1.0, Tags{"type": "mammal", "continent": "africa", "family": "antelope", "extant": ""})
	x := ParseSelector("animal", "extant,robotic")

	s.False(x.MatchesToken(&t))
}
