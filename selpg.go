package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/zhanzongyuan/selpg/engine"
	"github.com/zhanzongyuan/selpg/printer"
)

// some args
var (
	// mandatory options
	startPage = flag.IntP("start", "s", 1, "Page number of the file where you want to print start from. (must be positive)")
	endPage   = flag.IntP("end", "e", 1, "Page number of the file where you want to print end to. (must be positive)")

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

func reportErr(err error) {
	exitCode = 2
	fmt.Fprintln(os.Stderr, err)
	usage()
}

// main process
func main() {
	// warp selpgMain() function, so defer won't be execute after os.Exit(exitCode)
	selpgMain()
	os.Exit(exitCode)
}

func selpgMain() {
	// check flag correction
	flag.Parse()
	shortFlag := make(map[string]int)
	flag.Visit(func(f *flag.Flag) {
		shortFlag[f.Shorthand] = 1
	})

	if shortFlag["l"] == 1 && shortFlag["f"] == 1 {
		reportErr(errors.New("usage: arguments -l and -f can not be set at the same time!"))
		return
	}
	if shortFlag["e"] == 0 || shortFlag["s"] == 0 {
		reportErr(errors.New("usage: rguments -s and -e is needed!"))
		return
	} else if *startPage <= 0 || *endPage <= 0 || *startPage > *endPage {
		reportErr(errors.New("usage: rguments -s and -e must be positive, and argument -e must be equal or greater than -s"))
		return
	}

	// create selpg options
	selpgOpts := engine.SelectOptions{
		*startPage,
		*endPage,
		'\n',
		72,
	}
	selpgOpts.FlagLimit = *limitLine
	if *pagebreakFlag {
		selpgOpts.FlagLimit = 1
		selpgOpts.EndFlag = byte('\f')
	}

	// switch output writer to (os.Stdout | lp -d)
	var out io.Writer
	quit := make(chan error)
	if *destination == "" {
		out = os.Stdout
	} else {
		// create lp printer to the destination
		reader, writer := io.Pipe()
		out = writer
		go printer.RunPrinter(destination, reader, quit)
		defer func() {
			writer.Close()
			if err := <-quit; err != nil {
				exitCode = 2
				log.Fatal(err)
			}
		}()
	}

	// process input from stdin
	if flag.NArg() == 0 {
		// check stdin input mode, do not accept ModeCharDevice mode
		stat, err := os.Stdin.Stat()
		if err != nil {
			exitCode = 2
			log.Fatal(err)
		}
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			reportErr(errors.New("usage: invalid standard input!"))
			return
		}
		// process stdin stream and select page
		if err := engine.SelectPage(os.Stdin, out, &selpgOpts); err != nil {
			reportErr(err)
		}
		return
	}

	// process input file from file name
	path := flag.Arg(0)
	f, err := os.Open(path)
	defer f.Close()
	if _, err2 := f.Stat(); err2 != nil || err != nil {
		reportErr(err)
		return
	}
	if err := engine.SelectPage(f, out, &selpgOpts); err != nil {
		reportErr(err)
		return
	}
}
