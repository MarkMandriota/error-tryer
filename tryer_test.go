// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

// Package tryer - provides new simply error handling method.
package tryer

import (
	"errors"
	"fmt"
	"testing"
)

func TestTryer(t *testing.T) {
	defer WithPrint(recover())

	t.Log(WithPrint(foo())[0].(int))

	t.Log(WithPrint(bar())[0].(string))

	t.Log(WithPrint(foobar())...)
}

func BenchmarkTryer(b *testing.B) {
	v, err := foo()

	for i := 0; i < b.N; i++ {
		_ = WithNothing(v, err)
	}
}

func foo() (int, error) {
	return 123, errors.New("error foo: something went wrong")
}

func bar() (string, error) {
	return "hello", errors.New("error bar: something went wrong")
}

func foobar() (int, string, error, error) {
	a, errfoo := foo()

	b, errbar := bar()

	return a, b, fmt.Errorf("error foobar: %w", errfoo), fmt.Errorf("error foobar: %w", errbar)
}
