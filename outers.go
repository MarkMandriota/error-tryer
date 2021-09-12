// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

// Package tryer - provides new simply error handling method.
package tryer

import (
	"io"
	"log"
	"os"
)

var (
	TryWithNothing = New(Nothing())

	TryWithPrint = New(DefaultPrinter())

	TryWithFatal = New(DefaultFataller())

	TryWithPanic = New(Panicer())
)

func Nothing() Outer {
	return func(err error) {}
}

func Printer(w io.Writer, prefix string, flag int) Outer {
	l := log.New(w, prefix, flag)

	return func(err error) {
		l.Println(err)
	}
}

func DefaultPrinter() Outer {
	return Printer(os.Stderr, "", log.LstdFlags)
}

func Fataller(w io.Writer, prefix string, flag int) Outer {
	l := log.New(w, prefix, flag)

	return func(err error) {
		l.Fatalln(err)
	}
}

func DefaultFataller() Outer {
	return Fataller(os.Stderr, "", log.LstdFlags)
}

func Panicer() Outer {
	return func(err error) {
		panic(err)
	}
}
