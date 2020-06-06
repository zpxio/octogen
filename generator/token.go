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

type Token struct {
	Id      string
	Content string
	Rarity  float64
	Tags    map[string]string
	SetVars map[string]string
}

type Tags map[string]string

func BuildToken(id string, content string, rarity float64, tags Tags) Token {
	t := Token{
		Id:      id,
		Content: content,
		Rarity:  rarity,
		Tags:    tags,
		SetVars: make(map[string]string),
	}

	return t
}

func (t *Token) OnRenderSet(variable string, value string) {
	t.SetVars[variable] = value
}
