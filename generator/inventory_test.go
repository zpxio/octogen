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

type InventorySuite struct {
	suite.Suite
}

func TestInventorySuite(t *testing.T) {
	suite.Run(t, new(InventorySuite))
}

func (s *InventorySuite) TestCreateInventory() {
	i := CreateInventory()

	s.Empty(i.dictionary)
}

func (s *InventorySuite) TestAdd_Simple() {
	i := CreateInventory()

	testId := "ID"
	testContent := "Content"
	testRarity := 3.2
	tags := Tags{"A": "1", "B": "2"}

	i.AddToken(testId, testContent, testRarity, tags)

	s.Len(i.dictionary, 1)
	s.InDelta(testRarity, i.selectRange[testId], 0.001)
}

func BuildSampleInventory() *Inventory {
	i := CreateInventory()

	i.AddToken("Animal", "Aardvark", 1.0, Tags{"type": "mammal", "env": "ground", "family": "orycteropod"})
	i.AddToken("Animal", "Boomalope", 2.0, Tags{"type": "cryptid", "env": "ground", "family": "deer"})
	i.AddToken("Animal", "Capybara", 1.0, Tags{"type": "mammal", "env": "ground", "family": "rodent"})
	i.AddToken("Animal", "Cladoselache", 2.5, Tags{"type": "fish", "env": "water", "family": "shark"})

	i.AddToken("Description", "Angry", 1.0, Tags{"tone": "negative"})
	i.AddToken("Description", "Confused", 1.5, Tags{"tone": "negative"})
	i.AddToken("Description", "Reluctant", 1.0, Tags{"tone": "neutral"})
	i.AddToken("Description", "Happy", 2.5, Tags{"tone": "positive"})

	i.AddToken("AnimalType", "mammal", 3.0, Tags{})
	i.AddToken("AnimalType", "fish", 1.2, Tags{})
	i.AddToken("AnimalType", "cryptid", 3.0, Tags{})

	return i
}

func (s *InventorySuite) TestGetInstructions_Simple() {
	i := BuildSampleInventory()

	x, r := i.getTokens(ParseSelector("Animal", ""))

	s.InDelta(6.5, r, 0.001)
	s.Len(x, 4)
}

func (s *InventorySuite) TestGetInstructions_SingleTag() {
	i := BuildSampleInventory()

	x, r := i.getTokens(ParseSelector("Animal", "type=mammal"))

	s.InDelta(2, r, 0.001)
	s.Len(x, 2)
}

func (s *InventorySuite) TestGetInstructions_MultiTag() {
	i := BuildSampleInventory()

	x, r := i.getTokens(ParseSelector("Description", "tone=negative"))

	s.InDelta(2.5, r, 0.001)
	s.Len(x, 2)
}

func (s *InventorySuite) TestGetInstructions_NotFound() {
	i := BuildSampleInventory()

	x, r := i.getTokens(ParseSelector("ZipCode", ""))

	s.InDelta(0, r, 0.001)
	s.Len(x, 0)
}

func (s *InventorySuite) TestPick_Simple() {
	i := BuildSampleInventory()
	sel := ParseSelector("Animal", "")
	c := i.Pick(sel, 0)

	s.NotNil(c)
	s.Equal("Aardvark", c.Content)

	c2 := i.Pick(sel, 1.0)
	s.NotNil(c2)
	s.Equal("Cladoselache", c2.Content)
}

func (s *InventorySuite) TestPick_NoId() {
	i := BuildSampleInventory()
	c := i.Pick(ParseSelector("ZipCode", ""), 0)

	s.Nil(c)
}

func (s *InventorySuite) TestPick_NoMatchingTags() {
	i := BuildSampleInventory()
	c := i.Pick(ParseSelector("Animal", "type=bird"), 0)

	s.Nil(c)
}
