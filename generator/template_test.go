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

type GeneratorSuite struct {
	suite.Suite
}

func TestGeneratorSuite(t *testing.T) {
	suite.Run(t, new(GeneratorSuite))
}

func (s *GeneratorSuite) buildSampleInventory() *Inventory {
	i := CreateInventory()

	i.Add("Animal", "Aardvark", 1.0, Tags{"type": "mammal", "env": "ground", "family": "orycteropod"})
	i.Add("Animal", "Boomalope", 2.0, Tags{"type": "cryptid", "env": "ground", "family": "deer"})
	i.Add("Animal", "Capybara", 1.0, Tags{"type": "mammal", "env": "ground", "family": "rodent"})
	i.Add("Animal", "Cladoselache", 1.0, Tags{"type": "fish", "env": "water", "family": "shark"})

	i.Add("Description", "Angry", 1.0, Tags{"tone": "negative"})
	i.Add("Description", "Confused", 1.5, Tags{"tone": "negative"})
	i.Add("Description", "Reluctant", 1.0, Tags{"tone": "neutral"})
	i.Add("Description", "Happy", 2.5, Tags{"tone": "positive"})

	i.Add("AnimalType", "mammal", 3.0, Tags{})
	i.Add("AnimalType", "fish", 1.2, Tags{})
	i.Add("AnimalType", "cryptid", 3.0, Tags{})

	return i
}

func (s *GeneratorSuite) TestInitTemplate() {
	i := "Example: [Animal]"
	t := InitTemplate(i)

	s.Equal(i, t.working)
	s.Equal(RoundsMax, t.rounds)
}

func (s *GeneratorSuite) TestReplaceNextToken() {
	t := InitTemplate("Example: [Animal]")
	i := s.buildSampleInventory()

	replaced := t.replaceNextToken(i, 0.0)
	s.True(replaced)

	s.Equal("Example: Aardvark", t.working)
}

func (s *GeneratorSuite) TestReplaceNextToken_Multiple() {
	t := InitTemplate("Example: [Description] [Animal]")
	i := s.buildSampleInventory()

	t.replaceNextToken(i, 0.0)

	s.Equal("Example: Angry [Animal]", t.working)

	replaced := t.replaceNextToken(i, 0.0)
	s.True(replaced)

	s.Equal("Example: Angry Aardvark", t.working)
}

func (s *GeneratorSuite) TestReplaceNextToken_Tagged() {
	t := InitTemplate("Example: [Animal:family=rodent]")
	i := s.buildSampleInventory()

	replaced := t.replaceNextToken(i, 0.0)
	s.True(replaced)

	s.Equal("Example: Capybara", t.working)
}

func (s *GeneratorSuite) TestReplaceNextToken_Nested() {
	t := InitTemplate("Example: [Animal:type=[AnimalType]]")
	i := s.buildSampleInventory()

	replaced := t.replaceNextToken(i, 1.0)
	s.True(replaced)

	s.Equal("Example: [Animal:type=cryptid]", t.working)
}

func (s *GeneratorSuite) TestReplaceNextToken_NoToken() {
	t := InitTemplate("Example: Done")
	i := s.buildSampleInventory()

	replaced := t.replaceNextToken(i, 1.0)
	s.False(replaced)

	s.Equal("Example: Done", t.working)
}

func (s *GeneratorSuite) TestRender() {
	t := InitTemplate("Example: [Animal:type=[AnimalType]]")
	i := s.buildSampleInventory()

	result := t.Render(i, rng.UseStatic(0))

	s.Equal("Example: Aardvark", result)
}
