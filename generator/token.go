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

import "strings"

// Type Token represents a single item which can be placed into the generated output of a Generator.
type Token struct {
	Category   string
	Content    string
	Rarity     float64
	Properties map[string]string
	SetVars    map[string]string
}

// Type Properties defines the structure used to store token properties.
type Properties map[string]string

func BuildToken(category string, content string, rarity float64, tags Properties) Token {
	t := Token{
		Category:   category,
		Content:    content,
		Rarity:     rarity,
		Properties: tags,
		SetVars:    make(map[string]string),
	}

	return t
}

// OnRenderSet defines a State variable to set when this Token is rendered to Generator output.
func (t *Token) OnRenderSet(variable string, value string) {
	t.SetVars[variable] = value
}

// Normalize updates the Token to ensure that it matches required behaviors. Categories must not start
// or end with whitespace. Rarities must not be zero or negative. If the Rarity is invalid, it is set
// to a default of 1.0
func (t *Token) Normalize() {
	t.Category = strings.TrimSpace(t.Category)
	if t.Rarity <= 0.0 {
		t.Rarity = 1.0
	}
}

// IsValid checks to see if this Token is valid and usable for generation. Invalid Tokens may become valid
// if Normalize is called on them, but this is not guaranteed. If you are reading or creating Tokens with
// non-validated input, you should call Normalize before checking validity.
func (t *Token) IsValid() bool {
	if t.Category == "" {
		return false
	}

	if t.Content == "" {
		return false
	}

	if t.Rarity <= 0.0 {
		return false
	}

	return true
}
