# gir
A tools to list exposed items in a Go package (somehow similar to Python `dir()`).

# Source
`go get github.com/m4ns0ur/gir`

# Usage
```
$ gir
Usage: gir [options] package[.item]

Gir shows all items in the package scope.

Options:
  -l list
    	Show as a list
  -v verbose
    	Show more verbose details (use with -list)
```
```
$ gir fmt
[Errorf, Formatter, Fprint, Fprintf, Fprintln, Fscan, Fscanf, Fscanln, GoStringer, Print, Printf, Println, Scan, ScanState, Scanf, Scanln, Scanner, Sprint, Sprintf, Sprintln, Sscan, Sscanf, Sscanln, State, Stringer, init]
```
```
$ gir fmt.Printf
func Printf(format string, a ...interface{}) (n int, err error)
```
