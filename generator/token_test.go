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
	testId := "ID"
	testContent := "Content"
	testRarity := 3.2
	tags := Tags{"A": "1", "B": "2"}

	i := BuildToken(testId, testContent, testRarity, tags)

	s.Equal(testId, i.Id)
	s.Equal(testContent, i.Content)
	s.InDelta(testRarity, i.Rarity, 0.001)
	s.Len(i.Tags, len(tags))
}

func (s *TokenSuite) TestIsValid_Happy() {
	testId := "ID"
	testContent := "Content"
	testRarity := 3.2
	tags := Tags{"A": "1", "B": "2"}

	i := BuildToken(testId, testContent, testRarity, tags)

	s.True(i.IsValid())
}

func (s *TokenSuite) TestIsValid_NoId() {
	testId := ""
	testContent := "Content"
	testRarity := 3.2
	tags := Tags{"A": "1", "B": "2"}

	i := BuildToken(testId, testContent, testRarity, tags)

	s.False(i.IsValid())
}

func (s *TokenSuite) TestIsValid_NoContent() {
	testId := "ID"
	testContent := ""
	testRarity := 3.2
	tags := Tags{"A": "1", "B": "2"}

	i := BuildToken(testId, testContent, testRarity, tags)

	s.False(i.IsValid())
}

func (s *TokenSuite) TestIsValid_ZeroRarity() {
	testId := "ID"
	testContent := ""
	testRarity := 0.0
	tags := Tags{"A": "1", "B": "2"}

	i := BuildToken(testId, testContent, testRarity, tags)

	s.False(i.IsValid())
}

func (s *TokenSuite) TestIsValid_NegativeRarity() {
	testId := "ID"
	testContent := "Content"
	testRarity := -1.0
	tags := Tags{"A": "1", "B": "2"}

	i := BuildToken(testId, testContent, testRarity, tags)

	s.False(i.IsValid())
}

func (s *TokenSuite) TestNormalize_Noop() {
	testId := "ID"
	testContent := "Content"
	testRarity := 3.2
	tags := Tags{"A": "1", "B": "2"}

	t := Token{
		Id:      testId,
		Content: testContent,
		Rarity:  testRarity,
		Tags:    tags,
		SetVars: make(map[string]string),
	}

	t.Normalize()

	s.Equal(testId, t.Id)
	s.Equal(testContent, t.Content)
	s.Equal(testRarity, t.Rarity)
}

func (s *TokenSuite) TestNormalize_BadId() {
	testId := "ID  "
	testContent := "Content"
	testRarity := 3.2
	tags := Tags{"A": "1", "B": "2"}

	t := Token{
		Id:      testId,
		Content: testContent,
		Rarity:  testRarity,
		Tags:    tags,
		SetVars: make(map[string]string),
	}

	t.Normalize()

	s.Equal("ID", t.Id)
	s.Equal(testContent, t.Content)
	s.Equal(testRarity, t.Rarity)
}

func (s *TokenSuite) TestNormalize_BadContent() {
	testId := "ID"
	testContent := "    Content    "
	testRarity := 3.2
	tags := Tags{"A": "1", "B": "2"}

	t := Token{
		Id:      testId,
		Content: testContent,
		Rarity:  testRarity,
		Tags:    tags,
		SetVars: make(map[string]string),
	}

	t.Normalize()

	s.Equal(testId, t.Id)
	s.Equal(testContent, t.Content)
	s.Equal(testRarity, t.Rarity)
}

func (s *TokenSuite) TestNormalize_ZeroRarity() {
	testId := "ID"
	testContent := "Content"
	testRarity := 0.0
	tags := Tags{"A": "1", "B": "2"}

	t := Token{
		Id:      testId,
		Content: testContent,
		Rarity:  testRarity,
		Tags:    tags,
		SetVars: make(map[string]string),
	}

	t.Normalize()

	s.Equal(testId, t.Id)
	s.Equal(testContent, t.Content)
	s.Equal(1.0, t.Rarity)
}

func (s *TokenSuite) TestNormalize_NegativeRarity() {
	testId := "ID"
	testContent := "Content"
	testRarity := -1.2
	tags := Tags{"A": "1", "B": "2"}

	t := Token{
		Id:      testId,
		Content: testContent,
		Rarity:  testRarity,
		Tags:    tags,
		SetVars: make(map[string]string),
	}

	t.Normalize()

	s.Equal(testId, t.Id)
	s.Equal(testContent, t.Content)
	s.Equal(1.0, t.Rarity)
}
