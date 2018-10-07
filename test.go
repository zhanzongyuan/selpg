package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	stdin := new(bytes.Buffer)
	q := make(chan int)
	go func() {
		cmd := exec.Command("cat")
		cmd.Stdin = stdin
		// <-ctx.Done()

		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
		q <- 0
	}()

	io.WriteString(stdin, "hihihihi")
	<-q
	/*
		cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
		// stdout, err := cmd.StdoutPipe()
		stdout := new(bytes.Buffer)
		cmd.Stdout = stdout
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		var person struct {
			Name string
			Age  int
		}

		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
		if err := json.NewDecoder(stdout).Decode(&person); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	*/
	/*
		cmd := exec.Command("sleep", "5")
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Now wait ...")
		ctx, cannel := context.WithCancel(context.Background())
		defer cannel()

		q := make(chan int)
		go func(ctx context.Context) {
			dots := []string{"\r.   ", "\r..  ", "\r... ", "\r...."}
			for {
				select {
				case <-ctx.Done():
					fmt.Println("")
					q <- 0
					return
				default:
					for i := 0; i < len(dots); i++ {
						fmt.Print(dots[i])
						time.Sleep(time.Millisecond * 100)
					}
				}
			}
		}(ctx)

		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
		cannel()
		<-q
		fmt.Println("To the end.")
	*/
	/*
		c := make(chan int)
		ctx, cannel := context.WithTimeout(context.Background(), time.Second)
		defer cannel()
		n := 0
		go func(ctx context.Context) {
			for {
				select {
				case c <- 1:
					time.Sleep(100 * time.Millisecond)
					// n++
				case c <- 0:
					time.Sleep(100 * time.Millisecond)
					// n++
				case <-ctx.Done():
					fmt.Println(ctx.Err())
					close(c)
					return
				}
			}
		}(ctx)
		for i := range c {
			n++
			fmt.Println(n, i)
		}
	*/
	/*
		cmd := exec.Command("cat")
		// cmd.Stdin = bytes.NewBufferString("test.go")
		// out := new(bytes.Buffer)
		// cmd.Stdout = out
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	*/

	/*
		if path, err := exec.LookPath("python"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("path of python:", path)
		}
	*/

}
