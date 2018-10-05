package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

var (
	// mandatory options
	startPage = flag.IntP("start", "s", -1, "Page number of the file where you want to print start from.")
	endPage   = flag.IntP("end", "e", -1, "Page number of the file where you want to print end to.")

	// optional options
	limitLine     = flag.IntP("limit", "l", 72, "Line number for one page.")
	pagebreakFlag = flag.BoolP("pbflag", "f", false, "Flag to find page break or not.")
	destination   = flag.StringP("destination", "d", "", "Printer destination to print choesn page.")
)

func main() {
	flag.Parse()
	fmt.Println(flag.SortFlags())
	// flag.SortFlags = false
	fmt.Println(*startPage, *endPage, *limitLine, *pagebreakFlag, *destination)
}
