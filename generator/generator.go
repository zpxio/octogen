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

// Package generator contains functionality to generate randomized, structured text based on a
// mixture of randomized selection and configured selections.
package generator

import "github.com/zpxio/octogen/rng"

type Generator struct {
	instructions string
	inventory    *Inventory
	rng          rng.RandomSource
}

// CreateGenerator creates a reusable text generator based on the instructions provided and the
// given inventory. Each time the generator is used, it can potentially produce different output.
func CreateGenerator(instructions string, inventory *Inventory) *Generator {
	g := &Generator{
		instructions: instructions,
		inventory:    inventory,
		rng:          rng.UseSystem(),
	}

	return g
}

// Run executes the generator with a new empty State.
func (g *Generator) Run() string {
	state := CreateState()

	return g.RunWithState(state)
}

// RunWithState executes the generator with the supplied State. This function is used to execute Generators
// with pre-defined state for instruction sets that require variable substitution.
func (g *Generator) RunWithState(state *State) string {
	result := Render(g.instructions, g.inventory, state, g.rng)

	return result
}

// UseRandomSource assigns a RandomSource to use when picking tokens.
func (g *Generator) UseRandomSource(rng rng.RandomSource) {
	g.rng = rng
}
