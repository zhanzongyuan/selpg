package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

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
func processStream(in io.Reader, out io.Writer, q chan int) error {
	// process input stream
	pageIter, flagIter := 1, 0
	buffer := make([]byte, 16)
	n, err := in.Read(buffer)
	for err == nil {
		accStart, accEnd := 0, n

		for i := 0; i < n; i++ {
			// count iterator
			if pageendFlag == buffer[i] {

				flagIter = (flagIter + 1) % limitFlag
				// next page
				if flagIter == 0 {
					pageIter++
					if pageIter == *startPage {
						accStart = i + 1
					} else if pageIter == *endPage+1 {
						accEnd = i + 1
					}
				}
			}
		}

		if pageIter >= *startPage {
			io.WriteString(out, string(buffer[accStart:accEnd]))
		}
		if pageIter > *endPage {
			break
		}
		n, err = in.Read(buffer)
	}
	if *destination != "" {
		q <- exitCode
	}
	return nil
}

/*
func printer() {

}*/

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

	// make output io
	var out io.ReadWriter
	q := make(chan int)
	if *destination == "" {
		out = os.Stdout
	} else {
		// create lp printer to the destination
		out = new(bytes.Buffer)
		go func() {
			cmd := exec.Command("lp", "-d", *destination)
			cmd.Stdin = out
			quitCode := <-q
			if quitCode != 0 {
				return
			}
			stdoutStderr, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(string(stdoutStderr))
				exitCode = 2
				q <- exitCode
				log.Fatal(err)
			}
			fmt.Println(string(stdoutStderr))
			q <- 0
		}()

	}

	// TODO: process input from stdin
	if flag.NArg() == 0 {
		if err := processStream(os.Stdin, out, q); err != nil {
			reportErr(err)
			exitCode = 2
		}
		if *destination != "" {
			q <- exitCode
		}
		return
	}

	// TODO: process input file from file name
	path := flag.Arg(0)

	// check that file is valid
	f, err := os.Open(path)
	defer f.Close()
	if _, err2 := f.Stat(); err2 != nil || err != nil {
		reportErr(err)
		exitCode = 2
		if *destination != "" {
			q <- exitCode
		}
		return
	}

	if err := processStream(f, out, q); err != nil {
		reportErr(err)
		exitCode = 2
		if *destination != "" {
			q <- exitCode
		}
		return
	}

	if *destination != "" {
		<-q
	}
}
