# `selpg` -- a simple golang CLI program

`selpg` is a command line program to print specific page in the file. You can specify the page format to print specific page. Moreover, selpg allows you to print the specific page to system printer.

> This is my first golang CLI program. So it will be a simple demo for me to learn how to construct CLI program under CLI.

## Installation

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

> [My design blog](http://blog.zhanzy.xyz/2018/10/04/Go%E5%BC%80%E5%8F%91CLI%E5%AE%9E%E7%94%A8%E7%A8%8B%E5%BA%8F%E5%88%9D%E4%BD%93%E9%AA%8C/#%E8%AE%BE%E8%AE%A1selpg%E7%A8%8B%E5%BA%8F%E7%BB%93%E6%9E%84)

`engine.SelectPages(io.Reader, io.Writer, SelectOptions) error` : This function processes the input stream with interface `io.Reader`, and will write them to output interface `io.Writer`.

`printer.RunPrinter(*string, io.Reader, chan error)` : This function will execute `lp -d` command in a goroutine, so we can print output to a printer in the system. We design a reader-writer pipe to connect its input with `processStream` output synchronously. Also `chan error` can help printer goroutine quit synchronously and report error to main thread.

![](http://pg2vkewkk.bkt.clouddn.com/selpg%E7%A8%8B%E5%BA%8F%E7%BB%93%E6%9E%84%283%29.png)

<br>
