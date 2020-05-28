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

type ManualRand struct {
	values []float64
}

func UseManual() *ManualRand {
	return &ManualRand{
		values: []float64{},
	}
}

func (r *ManualRand) Next() float64 {
	if len(r.values) < 1 {
		panic("attempt to read from empty random source")
	}

	next := r.values[0]
	r.values = r.values[1:]

	return next
}

func (r *ManualRand) Clear() {
	r.values = []float64{}
}

func (r *ManualRand) Add(v ...float64) {
	r.values = append(r.values, v...)
}
