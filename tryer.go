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

// New - construct new Tryer with Outer.
func New(o Outer) Tryer {
	return func(v ...interface{}) []interface{} {
		j := 0

		for i := 0; i < len(v); i++ {
			switch v[i].(type) {
			case error:
				o(v[i].(error))
			case nil:
			default:
				v[j] = v[i]
				j++
			}
		}

		return v[:j]
	}
}
