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
	"github.com/zpxio/octogen/rng"
	"testing"
)

type RenderSuite struct {
	suite.Suite
}

func TestRenderSuite(t *testing.T) {
	suite.Run(t, new(RenderSuite))
}

func (s *RenderSuite) buildSampleInventory() *Inventory {
	i := CreateInventory()

	i.AddToken("Animal", "Aardvark", 1.0, Tags{"type": "mammal", "env": "ground", "family": "orycteropod"})
	i.AddToken("Animal", "Boomalope", 2.0, Tags{"type": "cryptid", "env": "ground", "family": "deer"})
	i.AddToken("Animal", "Capybara", 1.0, Tags{"type": "mammal", "env": "ground", "family": "rodent"})
	i.AddToken("Animal", "Cladoselache", 1.0, Tags{"type": "fish", "env": "water", "family": "shark"})

	i.AddToken("Description", "Angry", 1.0, Tags{"tone": "negative"})
	i.AddToken("Description", "Confused", 1.5, Tags{"tone": "negative"})
	i.AddToken("Description", "Reluctant", 1.0, Tags{"tone": "neutral"})
	i.AddToken("Description", "Happy", 2.5, Tags{"tone": "positive"})

	i.AddToken("AnimalType", "mammal", 3.0, Tags{})
	i.AddToken("AnimalType", "fish", 1.2, Tags{})
	i.AddToken("AnimalType", "cryptid", 3.0, Tags{})

	return i
}

func (s *RenderSuite) TestReplaceNextToken() {
	t := "Example: [Animal]"
	i := s.buildSampleInventory()
	x := CreateState()

	working, replaced := replaceNextToken(t, i, x, rng.UseStatic(0))
	s.True(replaced)

	s.Equal("Example: Aardvark", working)
}

func (s *RenderSuite) TestReplaceNextToken_Multiple() {
	t := "Example: [Description] [Animal]"
	i := s.buildSampleInventory()
	x := CreateState()

	working, replaced := replaceNextToken(t, i, x, rng.UseStatic(0))

	s.Equal("Example: Angry [Animal]", working)

	working, replaced = replaceNextToken(working, i, x, rng.UseStatic(0))

	s.True(replaced)
	s.Equal("Example: Angry Aardvark", working)
}

func (s *RenderSuite) TestReplaceNextToken_Tagged() {
	t := "Example: [Animal:family=rodent]"
	i := s.buildSampleInventory()
	x := CreateState()

	working, replaced := replaceNextToken(t, i, x, rng.UseStatic(0))

	s.True(replaced)
	s.Equal("Example: Capybara", working)
}

func (s *RenderSuite) TestReplaceNextToken_Nested() {
	t := "Example: [Animal:type=[AnimalType]]"
	i := s.buildSampleInventory()
	x := CreateState()

	working, replaced := replaceNextToken(t, i, x, rng.UseStatic(1))

	s.True(replaced)
	s.Equal("Example: [Animal:type=cryptid]", working)
}

func (s *RenderSuite) TestReplaceNextToken_NoToken() {
	t := "Example: Done"
	i := s.buildSampleInventory()
	x := CreateState()

	working, replaced := replaceNextToken(t, i, x, rng.UseStatic(1))

	s.False(replaced)
	s.Equal("Example: Done", working)
}

func (s *RenderSuite) TestReplaceNextVar() {
	t := "Example: [Animal:type=[$type]]"

	state := CreateState()
	state.Vars["type"] = "mammal"

	working, replaced := replaceNextVar(t, state)

	s.True(replaced)
	s.Equal("Example: [Animal:type=mammal]", working)
}

func (s *RenderSuite) TestReplaceNextVar_NoVar() {
	t := "Example: [Animal:type=amphibian]"

	state := CreateState()
	state.Vars["type"] = "mammal"

	working, replaced := replaceNextVar(t, state)

	s.False(replaced)
	s.Equal("Example: [Animal:type=amphibian]", working)
}

func (s *RenderSuite) TestReplaceNextVar_UndefinedVar() {
	t := "Example: [Animal:type=[$selectType]]"

	state := CreateState()
	state.Vars["type"] = "mammal"

	working, replaced := replaceNextVar(t, state)

	s.False(replaced)
	s.Equal("Example: [Animal:type=[$selectType]]", working)
}

func (s *RenderSuite) TestRender() {
	t := "Example: [Animal:type=[AnimalType]]"
	i := s.buildSampleInventory()

	result := Render(t, i, CreateState(), rng.UseStatic(0))

	s.Equal("Example: Aardvark", result)
}

func (s *RenderSuite) TestRender_WithVar() {
	t := "Example: [Animal:type=[$type]]"
	i := s.buildSampleInventory()
	x := CreateState()
	x.Vars["type"] = "mammal"

	result := Render(t, i, x, rng.UseStatic(0))

	s.Equal("Example: Aardvark", result)
}

func (s *RenderSuite) TestRender_MultipleRender() {
	t := "Example: [Animal:type=[$type]]"
	i := s.buildSampleInventory()
	x := CreateState()
	x.Vars["type"] = "mammal"

	result := Render(t, i, x, rng.UseStatic(0))
	s.Equal("Example: Aardvark", result)

	result2 := Render(t, i, x, rng.UseStatic(1))
	s.Equal("Example: Capybara", result2)
}

func (s *RenderSuite) TestRender_WithVarSet() {
	t := "Example: [Animal:family=ape] Sentience: [$sentience]"
	i := s.buildSampleInventory()

	t1 := i.AddToken("Animal", "Human", 1.0, Tags{"type": "mammal", "env": "land", "family": "ape"})
	t1.OnRenderSet("human", "normal")
	t1.OnRenderSet("sentience", "full")

	t2 := i.AddToken("Animal", "Chimpanzee", 1.0, Tags{"type": "mammal", "env": "land", "family": "ape"})
	t2.OnRenderSet("sentience", "high")

	t3 := i.AddToken("AnimalFamily", "ape", 1.0, Tags{"special": "true"})
	t3.OnRenderSet("sentience", "moderate")

	result1 := Render(t, i, CreateState(), rng.UseStatic(0))
	s.Equal("Example: Human Sentience: full", result1)

	result2 := Render(t, i, CreateState(), rng.UseStatic(1.2))
	s.Equal("Example: Chimpanzee Sentience: high", result2)
}
