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

package rng

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RngSuite struct {
	suite.Suite
}

func TestRngSuite(t *testing.T) {
	suite.Run(t, new(RngSuite))
}

func (s *RngSuite) TestManualInit() {
	r := UseManual()

	s.NotNil(r)
	s.Empty(r.values)
}

func (s *RngSuite) TestManualInit_WithValues() {
	r := UseManual(0.1, 0.2, 0.3)

	s.NotNil(r)
	s.NotEmpty(r.values)
	s.InDelta(0.1, r.Next(), 0.001)
	s.InDelta(0.2, r.Next(), 0.001)
	s.InDelta(0.3, r.Next(), 0.001)
}

func (s *RngSuite) TestPanicOnEmpty() {
	r := UseManual()

	s.NotNil(r)

	s.Panics(func() {
		r.Next()
	})
}

func (s *RngSuite) TestUsage() {
	r := UseManual()

	s.Require().NotNil(r)

	r.Add(1.0)
	s.Equal(1.0, r.Next())

	r.Add(2.0, 3.0, 4.0)
	s.Equal(2.0, r.Next())

	r.Clear()
	r.Add(5.0, 6.0)
	s.Equal(5.0, r.Next())
	s.Equal(6.0, r.Next())
}
