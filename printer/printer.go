package printer

import (
	"io"
	"os"
	"os/exec"
)

// printer goroutine
func RunPrinter(destination *string, reader io.Reader, quit chan error) {
	cmd := exec.Command("lp", "-d", *destination)
	cmd.Stdin = reader

	// create command standard output and input output reader
	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		quit <- err
		return
	}

	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		quit <- err
		return
	}

	// start command and wait
	if err := cmd.Start(); err != nil {
		quit <- err
		return
	}
	if _, err := io.Copy(os.Stdout, stdoutReader); err != nil {
		quit <- err
		return
	}
	if _, err := io.Copy(os.Stderr, stderrReader); err != nil {
		quit <- err
		return
	}
	if err := cmd.Wait(); err != nil {
		quit <- err
		return
	}
	quit <- nil
}
