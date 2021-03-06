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
	"os"
	"path/filepath"
	"testing"
)

type InventorySuite struct {
	suite.Suite
}

func TestInventorySuite(t *testing.T) {
	suite.Run(t, new(InventorySuite))
}

func DataDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Failed to look up current directory")
	}
	basedir := filepath.Dir(cwd)
	return filepath.Join(basedir, "testdata")
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
	tags := Properties{"A": "1", "B": "2"}

	i.AddToken(testId, testContent, testRarity, tags)

	s.Len(i.dictionary, 1)
	s.InDelta(testRarity, i.selectRange[testId], 0.001)
}

func BuildSampleInventory() *Inventory {
	i := CreateInventory()

	i.AddToken("Animal", "Aardvark", 1.0, Properties{"type": "mammal", "env": "ground", "family": "orycteropod"})
	i.AddToken("Animal", "Boomalope", 2.0, Properties{"type": "cryptid", "env": "ground", "family": "deer"})
	i.AddToken("Animal", "Capybara", 1.0, Properties{"type": "mammal", "env": "ground", "family": "rodent"})
	i.AddToken("Animal", "Cladoselache", 2.5, Properties{"type": "fish", "env": "water", "family": "shark"})

	i.AddToken("Description", "Angry", 1.0, Properties{"tone": "negative"})
	i.AddToken("Description", "Confused", 1.5, Properties{"tone": "negative"})
	i.AddToken("Description", "Reluctant", 1.0, Properties{"tone": "neutral"})
	i.AddToken("Description", "Happy", 2.5, Properties{"tone": "positive"})

	i.AddToken("AnimalType", "mammal", 3.0, Properties{})
	i.AddToken("AnimalType", "fish", 1.2, Properties{})
	i.AddToken("AnimalType", "cryptid", 3.0, Properties{})

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

func (s *InventorySuite) TestLoad_Simple() {
	testFile := filepath.Join(DataDir(), "inv_animals.yml")

	i := CreateInventory()

	s.Empty(i.dictionary)

	err := i.Load(testFile)

	s.NoError(err)
	s.Len(i.dictionary, 3)
	s.Len(i.dictionary["Animal"], 4)
	s.Len(i.dictionary["Description"], 4)
	s.Len(i.dictionary["AnimalType"], 3)
}

func (s *InventorySuite) TestLoad_MissingFile() {
	testFile := filepath.Join(DataDir(), "file_not_found.yml")

	i := CreateInventory()
	s.Empty(i.dictionary)

	err := i.Load(testFile)

	s.Error(err)
}

func (s *InventorySuite) TestLoad_BadFile() {
	testFile := filepath.Join(DataDir(), "not_tokens.yml")

	i := CreateInventory()
	s.Empty(i.dictionary)

	err := i.Load(testFile)

	s.NoError(err)
	s.Len(i.dictionary, 0)
}

func (s *InventorySuite) TestLoad_NotActuallyYaml() {
	testFile := filepath.Join(DataDir(), "not_yaml.json")

	i := CreateInventory()
	s.Empty(i.dictionary)

	err := i.Load(testFile)

	s.Error(err)
}

func (s *InventorySuite) TestLoad_PartiallyBadFile() {
	testFile := filepath.Join(DataDir(), "partially_bad.yml")

	i := CreateInventory()
	s.Empty(i.dictionary)

	err := i.Load(testFile)

	s.NoError(err)
	s.Len(i.dictionary, 2)
	s.Len(i.dictionary["Animal"], 4)
	s.Len(i.dictionary["AnimalType"], 3)

}
