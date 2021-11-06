// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

// Package tryer - provides new simply error handling method.
package tryer

type (
	// Outer - outing errors.
	Outer func(err error)

	// Tryer - trying all values and for each non-nil error, calls Outer.
	// Returns all not error/nil values.
	Tryer func(v ...interface{}) []interface{}
)

func New(o Outer) Tryer {
	return func(v ...interface{}) []interface{} {
		j := len(v)

		for i := len(v) - 1; i > 0; i-- {
			if x, ok := v[i].(error); ok && x != nil {
				o(x)
			} else {
				j--
				v[j] = v[i]
			}
		}

		return v[:j]
	}
}
