// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Selpg select file pages Go programs.

Usage: selpg [OPTION...] [FILE]...
	-s, --start int            Page number of the file where you want to print start from. (default 1)
    -e, --end int              Page number of the file where you want to print end to. (default 1)
	-l, --limit int            Line number for one page. (default 72)
	-f, --pbflag               Flag to find page break or not.
	-d, --destination string   Printer destination to print choesn page.

Examples

select page 1 to print.
	selpg -s1 -e1 *

*/

package main
