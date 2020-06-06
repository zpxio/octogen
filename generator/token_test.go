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

type TokenSuite struct {
	suite.Suite
}

func TestTokenSuite(t *testing.T) {
	suite.Run(t, new(TokenSuite))
}

func (s *TokenSuite) TestBuildInstruction_Simple() {
	testCategory := "ID"
	testContent := "Content"
	testRarity := 3.2
	tags := Properties{"A": "1", "B": "2"}

	i := BuildToken(testCategory, testContent, testRarity, tags)

	s.Equal(testCategory, i.Category)
	s.Equal(testContent, i.Content)
	s.InDelta(testRarity, i.Rarity, 0.001)
	s.Len(i.Properties, len(tags))
}

func (s *TokenSuite) TestIsValid_Happy() {
	testCategory := "ID"
	testContent := "Content"
	testRarity := 3.2
	tags := Properties{"A": "1", "B": "2"}

	i := BuildToken(testCategory, testContent, testRarity, tags)

	s.True(i.IsValid())
}

func (s *TokenSuite) TestIsValid_NoCategory() {
	testCategory := ""
	testContent := "Content"
	testRarity := 3.2
	tags := Properties{"A": "1", "B": "2"}

	i := BuildToken(testCategory, testContent, testRarity, tags)

	s.False(i.IsValid())
}

func (s *TokenSuite) TestIsValid_NoContent() {
	testCategory := "ID"
	testContent := ""
	testRarity := 3.2
	tags := Properties{"A": "1", "B": "2"}

	i := BuildToken(testCategory, testContent, testRarity, tags)

	s.False(i.IsValid())
}

func (s *TokenSuite) TestIsValid_ZeroRarity() {
	testCategory := "ID"
	testContent := ""
	testRarity := 0.0
	tags := Properties{"A": "1", "B": "2"}

	i := BuildToken(testCategory, testContent, testRarity, tags)

	s.False(i.IsValid())
}

func (s *TokenSuite) TestIsValid_NegativeRarity() {
	testCategory := "ID"
	testContent := "Content"
	testRarity := -1.0
	tags := Properties{"A": "1", "B": "2"}

	i := BuildToken(testCategory, testContent, testRarity, tags)

	s.False(i.IsValid())
}

func (s *TokenSuite) TestNormalize_Noop() {
	testCategory := "ID"
	testContent := "Content"
	testRarity := 3.2
	tags := Properties{"A": "1", "B": "2"}

	t := Token{
		Category:   testCategory,
		Content:    testContent,
		Rarity:     testRarity,
		Properties: tags,
		SetVars:    make(map[string]string),
	}

	t.Normalize()

	s.Equal(testCategory, t.Category)
	s.Equal(testContent, t.Content)
	s.Equal(testRarity, t.Rarity)
}

func (s *TokenSuite) TestNormalize_BadCategory() {
	testCategory := "ID  "
	testContent := "Content"
	testRarity := 3.2
	tags := Properties{"A": "1", "B": "2"}

	t := Token{
		Category:   testCategory,
		Content:    testContent,
		Rarity:     testRarity,
		Properties: tags,
		SetVars:    make(map[string]string),
	}

	t.Normalize()

	s.Equal("ID", t.Category)
	s.Equal(testContent, t.Content)
	s.Equal(testRarity, t.Rarity)
}

func (s *TokenSuite) TestNormalize_BadContent() {
	testCategory := "ID"
	testContent := "    Content    "
	testRarity := 3.2
	tags := Properties{"A": "1", "B": "2"}

	t := Token{
		Category:   testCategory,
		Content:    testContent,
		Rarity:     testRarity,
		Properties: tags,
		SetVars:    make(map[string]string),
	}

	t.Normalize()

	s.Equal(testCategory, t.Category)
	s.Equal(testContent, t.Content)
	s.Equal(testRarity, t.Rarity)
}

func (s *TokenSuite) TestNormalize_ZeroRarity() {
	testCategory := "ID"
	testContent := "Content"
	testRarity := 0.0
	tags := Properties{"A": "1", "B": "2"}

	t := Token{
		Category:   testCategory,
		Content:    testContent,
		Rarity:     testRarity,
		Properties: tags,
		SetVars:    make(map[string]string),
	}

	t.Normalize()

	s.Equal(testCategory, t.Category)
	s.Equal(testContent, t.Content)
	s.Equal(1.0, t.Rarity)
}

func (s *TokenSuite) TestNormalize_NegativeRarity() {
	testCategory := "ID"
	testContent := "Content"
	testRarity := -1.2
	tags := Properties{"A": "1", "B": "2"}

	t := Token{
		Category:   testCategory,
		Content:    testContent,
		Rarity:     testRarity,
		Properties: tags,
		SetVars:    make(map[string]string),
	}

	t.Normalize()

	s.Equal(testCategory, t.Category)
	s.Equal(testContent, t.Content)
	s.Equal(1.0, t.Rarity)
}
