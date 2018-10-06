package main

import (
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

// some args
var (
	// mandatory options
	startPage = flag.IntP("start", "s", 1, "Page number of the file where you want to print start from.")
	endPage   = flag.IntP("end", "e", 1, "Page number of the file where you want to print end to.")

	// optional options
	limitLine     = flag.IntP("limit", "l", 72, "Line number for one page.")
	pagebreakFlag = flag.BoolP("pbflag", "f", false, "Flag to find page break or not.")
	destination   = flag.StringP("destination", "d", "", "Printer destination to print choesn page.")
)

var (
	pageendFlag = byte('\n')
	limitFlag   = 72
)

// system variable
var (
	exitCode = 0
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: selpg [OPTION...] [FILE]...")
	flag.PrintDefaults()
}

// initial flag here
func init() {
	flag.CommandLine.SortFlags = false
	flag.Usage = usage
}

// utils
func processStream(in io.Reader) error {
	pageIter, flagIter := 1, 0

	buffer := make([]byte, 1)
	_, err := in.Read(buffer)
	for err == nil {
		if pageIter >= *startPage && pageIter <= *endPage {
			fmt.Fprintf(os.Stdout, "%s", string(buffer))
		}

		if pageendFlag == buffer[0] {
			flagIter = (flagIter + 1) % limitFlag

			if flagIter == 0 {
				pageIter++
			}
		}

		_, err = in.Read(buffer)
	}
	return nil
}
func reportErr(err error) {
	fmt.Fprintln(os.Stderr, "[Error]:", err)
}

// main process
func main() {
	// warp selpgMain() function, so defer won't be execute after os.Exit(exitCode)
	selpgMain()
	os.Exit(exitCode)
}

func selpgMain() {
	// TODO: check flag correction
	flag.Parse()
	limitFlag = *limitLine
	if *pagebreakFlag {
		limitFlag = 1
		pageendFlag = byte('\f')
	}
	// TODO: process input from stdin
	if flag.NArg() == 0 {
		if err := processStream(os.Stdin); err != nil {
			reportErr(err)
			exitCode = 0
		}
		return
	}

	// TODO: process input from several files name
	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)

		// check that file is valid
		f, err := os.Open(path)
		defer f.Close()
		if _, err2 := f.Stat(); err2 != nil || err != nil {
			reportErr(err)
			exitCode = 2
			return
		}

		if err := processStream(f); err != nil {
			reportErr(err)
			exitCode = 2
			return
		}
	}
}
