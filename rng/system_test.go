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

func (s *RngSuite) TestSystemInit() {
	r := UseSystem()

	s.NotNil(r)
}

func (s *RngSuite) TestSystemUsage() {
	r := UseSystem()

	s.Require().NotNil(r)

	for i := 0; i < 1000; i++ {
		s.InDelta(0.5, r.Next(), 0.5)
	}
}
