# gir
A tools to list all exported items in scope of a Go package (somehow similar to Python `dir()`).

# Install
`GO111MODULE=on go get github.com/m4ns0ur/gir`

Make sure `$GOPATH/bin` is in the path.

# Usage
```
$ gir -h
Usage: gir [options] package[:item]

Gir shows all exported items in the package scope.

Options:
  -h help
    	Show this help
  -l list
    	Show as a list
  -u unexported
    	Show unexported items as well
  -v verbose
    	Show more verbose details (use with -list)
```
```
$ gir fmt
[Errorf, Formatter, Fprint, Fprintf, Fprintln, Fscan, Fscanf, Fscanln, GoStringer, Print, Printf, Println, Scan, ScanState, Scanf, Scanln, Scanner, Sprint, Sprintf, Sprintln, Sscan, Sscanf, Sscanln, State, Stringer, init]
```
```
$ gir fmt:Printf
func Printf(format string, a ...interface{}) (n int, err error)
```
```
$ gir github.com/fatih/color:Bold
const Bold Attribute
```
