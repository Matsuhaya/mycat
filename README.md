# mycat

This is my own cat command.

## Usage

Get a list of file paths and output them in the order given in the standard output.

```
$ go run main.go a.txt b.txt

a
aa
aaa
b
bb
bbb
```

When the `-n` option is specified, a line number is displayed for each line.
Line numbers are serial numbers in all files.

```
$ go run main.go -n a.txt b.txt

1: a
2: aa
3: aaa
4: b
5: bb
6: bbb
```