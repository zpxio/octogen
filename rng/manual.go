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

// ManualRand defines a RandomSource which is manually fed random values to be
// retrieved. This is very useful for taking complete control of random number generation
// for unit testing.
type ManualRand struct {
	values []float64
}

// UseManual creates a new ManualRand containing the supplied values.
func UseManual(v ...float64) *ManualRand {
	r := &ManualRand{
		values: []float64{},
	}

	r.Add(v...)

	return r
}

// Next retrieves the next value in the queue of random numbers. If no values have been stored in
// buffer, then the function panics.
func (r *ManualRand) Next() float64 {
	if len(r.values) < 1 {
		panic("attempt to read from empty random source")
	}

	next := r.values[0]
	r.values = r.values[1:]

	return next
}

// Clear removes all stored values in the random queue.
func (r *ManualRand) Clear() {
	r.values = []float64{}
}

// Add adds new values to the queue of values to supply via the Next function.
func (r *ManualRand) Add(v ...float64) {
	r.values = append(r.values, v...)
}
