package engine

import (
	"errors"
	"io"
)

type SelectOptions struct {
	StartPage int
	EndPage   int
	EndFlag   byte
	FlagLimit int
}

// utils
func SelectPages(in io.Reader, out io.Writer, opts *SelectOptions) error {
	// process input stream
	pageIter, flagIter, writedFlag := 1, 0, false

	// deal page with flag '\f'
	buffer := make([]byte, 16)
	n, err := in.Read(buffer)

	for err == nil {
		accStart, accEnd := 0, n

		for i := 0; i < n; i++ {
			// count iterator
			if opts.EndFlag == buffer[i] {

				flagIter = (flagIter + 1) % opts.FlagLimit
				// next page
				if flagIter == 0 {
					pageIter++
					// find output interval in byte buffer.
					if pageIter == opts.StartPage {
						accStart = i + 1
					} else if pageIter == opts.EndPage+1 {
						accEnd = i + 1
					}
				}
			}
		}

		if pageIter >= opts.StartPage {
			writedFlag = true
			io.WriteString(out, string(buffer[accStart:accEnd]))
		}
		if pageIter > opts.EndPage {
			break
		}
		n, err = in.Read(buffer)
	}
	if writedFlag {
		return nil
	} else {
		return errors.New("usage: page number out of file range or input stream is empty.")
	}
}
