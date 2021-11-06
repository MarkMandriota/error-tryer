// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

// Package tryer - provides new simply error handling method.
package tryer

import (
	"log"
)

var (
	WithNothing = New(Nothing())

	WithPrint = New(Printer(log.Default()))

	WithFatal = New(Fataller(log.Default()))

	WithPanic = New(Panicer())
)

func Nothing() Outer {
	return func(err error) {}
}

func Printer(l *log.Logger) Outer {
	return func(err error) {
		l.Println(err)
	}
}

func Fataller(l *log.Logger) Outer {
	return func(err error) {
		l.Fatalln(err)
	}
}

func Panicer() Outer {
	return func(err error) {
		panic(err)
	}
}

func Sender(s chan<- error) Outer {
	return func(err error) {
		s <- err
	}
}

func Compose(o ...Outer) Outer {
	return func(err error) {
		for _, outer := range o {
			outer(err)
		}
	}
}
