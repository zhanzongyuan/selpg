# `selpg` -- a simple golang CLI program demo

> This is my first golang CLI program. So it will be a simple demo for me to learn how to construct CLI program under CLI.

## Install

[Go environment](https://golang.org/doc/install) is required!

```shell
go get github.com/zhanzongyuan/selpg
```



## Usage

**Command format:**

```
selpg -s [start_page] -e [end_page] [option...] -- [path]
```

**Options**:

`-s, --start int `:

​	Page number of the file where you want to print start from (must be positive).
`-e, --end int`:

​	Page number of the file where you want to print end to (must be positive).
`-l, --limit int`:

​	Line number for one page. (default 72)
`-f, --pbflag`:

​	Flag to find page break or not.
`-d, --destination string`:

​	Printer destination to print choesn page.

**Example**:



## Design

`processStream(io.Reader, io.Writer) error` : This function processes the input stream with interface `io.Reader`, and will write them to output interface `io.Writer`.

`runPrinter(io.Reader, chan error)` : This function will execute `lp -d` command in a goroutine, so we can print output to a printer in the system. We design a reader-writer pipe to connect its input with `processStream` output synchronously. Also `chan error` can help printer goroutine quit synchronously and report error to main thread.

![](http://pg2vkewkk.bkt.clouddn.com/18-10-11/53247950.jpg)

