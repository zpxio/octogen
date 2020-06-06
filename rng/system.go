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

import "math/rand"

// SystemRand defines a RandomSource that uses the native operating system's random number
// generation.
type SystemRand struct {
}

// UseSystem creates a new SystemRand.
func UseSystem() *SystemRand {
	return &SystemRand{}
}

// Next returns a new pseudo-random number from the operating systems random number source.
func (s *SystemRand) Next() float64 {
	return rand.Float64()
}
