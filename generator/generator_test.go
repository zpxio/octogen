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

func (s *GeneratorSuite) TestCreateGenerator() {
	i := BuildSampleInventory()
	t := "Test [Animal]"
	g := CreateGenerator(t, i)

	s.NotNil(g)
	s.Equal(t, g.instructions)
	s.Same(i, g.inventory)
}
func (s *GeneratorSuite) TestUseRandomSource() {
	i := BuildSampleInventory()
	t := "Test [Animal]"
	g := CreateGenerator(t, i)

	s.NotNil(g)

	customRng := rng.UseStatic(1)
	g.UseRandomSource(customRng)

	s.Same(customRng, g.rng)
}

func (s *GeneratorSuite) TestRun() {
	i := BuildSampleInventory()
	t := "Test [Animal]"
	g := CreateGenerator(t, i)
	g.UseRandomSource(rng.UseStatic(0))

	s.NotNil(g)

	result := g.Run()

	s.Equal("Test Aardvark", result)
}

func (s *GeneratorSuite) TestRun_Multiple() {
	i := BuildSampleInventory()
	t := "Test [Animal]"
	g := CreateGenerator(t, i)

	r := rng.UseManual()
	r.Add(0, 1)
	g.UseRandomSource(r)

	s.NotNil(g)

	result1 := g.Run()
	result2 := g.Run()

	s.Equal("Test Aardvark", result1)
	s.Equal("Test Cladoselache", result2)
}

func (s *GeneratorSuite) TestRunWithState() {
	i := BuildSampleInventory()
	t := "Test [Animal:type=[$animalType]]"
	g := CreateGenerator(t, i)
	g.UseRandomSource(rng.UseStatic(0))

	s.NotNil(g)

	x := CreateState()
	x.Vars["animalType"] = "mammal"
	result := g.RunWithState(x)

	s.Equal("Test Aardvark", result)
}
