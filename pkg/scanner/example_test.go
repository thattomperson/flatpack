// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner_test

import (
	"fmt"

	"ttp.sh/flatpack/pkg/scanner"
	"ttp.sh/flatpack/pkg/token"
)

func ExampleScanner_Scan() {
	// src is the input that we want to tokenize.
	src := []byte("Math.rand(1) * 500 // Euler")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
	// output
	// 1:1	IDENT	"Math"
	// 1:5	.	""
	// 1:6	IDENT	"rand"
	// 1:10	(	""
	// 1:11	NUMBER	"1"
	// 1:12	)	""
	// 1:14	*	""
	// 1:16	NUMBER	"500"
	// 1:20	;	"\n"
	// 1:20	COMMENT	"// Euler"
}
