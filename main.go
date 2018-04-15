package main

import (
	"flag"
	"fmt"
	"go/importer"
	"go/types"
	"os"
	"strings"
)

var list bool
var verbose bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gir [options] package[.item]\n\nGir shows all items in the package scope.\n\nOptions:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&list, "l", false, "Show as a `list`")
	flag.BoolVar(&verbose, "v", false, "Show more `verbose` details (use with -list)")
}

func main() {
	flag.Parse()
	if !flag.Parsed() || flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	pn := flag.Arg(0)
	var it string
	if strings.Contains(pn, ".") {
		pp := strings.Split(pn, ".")
		if len(pp) > 2 || pp[0] == "" || pp[1] == "" {
			flag.Usage()
			os.Exit(1)
		}
		pn = pp[0]
		it = pp[1]
	}

	p, err := importer.Default().Import(pn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	s := p.Scope()
	if it != "" {
		o := s.Lookup(it)
		if o == nil {
			fmt.Fprintf(os.Stderr, "can't find item: %q in the scope of import: %q\n", it, pn)
			os.Exit(1)
		}
		fmt.Println(stripInternalRef(o, pn))
		os.Exit(0)
	}

	nn := s.Names()
	if !list {
		fmt.Printf("[%s]\n", strings.Join(nn, ", "))
		os.Exit(0)
	}

	for _, n := range nn {
		o := s.Lookup(n)
		if o == nil {
			continue
		}
		if verbose {
			fmt.Println(stripInternalRef(o, pn))
		} else {
			fmt.Printf("%s\t%s\n", strings.Split(o.String(), " ")[0], o.Name())
		}
	}
}

func stripInternalRef(o types.Object, pn string) string {
	return strings.Replace(o.String(), pn+".", "", -1)
}
