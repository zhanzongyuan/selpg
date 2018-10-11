# `selpg` -- a simple golang CLI program demo

> This is my first golang CLI program. So it will be a simple demo for me to learn how to construct CLI program under CLI.

## Install

[Go installation](https://golang.org/doc/install) required!

After installing go and setting workspace, try:

```shell
go get -u github.com/zhanzongyuan/selpg
```

Then, check if install the command:

```shell
selpg -h
```

If you see the help information, then you can use the `selpg` to print you file as you want!

<br>

## Usage

**Command format:**

```
selpg -s [start_page] -e [end_page] [option...] -- [path]
```

**Options**:

`-s, --start int `: Page number of the file where you want to print start from (must be positive).

`-e, --end int`: Page number of the file where you want to print end to (must be positive).

`-l, --limit int`: Line number for one page. (default 72)

`-f, --pbflag`: Flag to find page break or not.

`-d, --destination string`: Printer destination to print choesn page.

**Examples**:

1. Print the first page on screen

```shell
selpg -s1 -e1 input_file
# or input from redirect
selpg -s1 -e1 < input_file
# or input from pipe
other_command | selpg -s1 -e1
```

2. Print the first page with 66 lines limit

```shell
selpg -s1 -e1 -l66 input_file # every page of file will be limited to 66 lines
```

3. Print the first page by page break`'\f'`

```shell
# notice: '-f' flag and '-l' flag shouldn't be set at the same time
# set '-f' flag mean every page will end by '\f'.
selpg -s1 -e1 -f input_file
```

4. Print the first page to printer with destination. You can use [cups-pdf](http://terokarvinen.com/2011/print-pdf-from-command-line-cups-pdf-lpr-p-pdf) to install command line pdf printer to try this flag.

```shell
# Print the first page to PDF printer (install cups-pdf first)
selpg -s1 -e1 -dPDF input_file 
```

<br>

## Design

`processStream(io.Reader, io.Writer) error` : This function processes the input stream with interface `io.Reader`, and will write them to output interface `io.Writer`.

`runPrinter(io.Reader, chan error)` : This function will execute `lp -d` command in a goroutine, so we can print output to a printer in the system. We design a reader-writer pipe to connect its input with `processStream` output synchronously. Also `chan error` can help printer goroutine quit synchronously and report error to main thread.

![](http://pg2vkewkk.bkt.clouddn.com/18-10-11/53247950.jpg)

