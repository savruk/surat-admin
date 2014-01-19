package main

import (
	"flag"

	"github.com/savruk/surat"
)

var flush bool

func init() {
	flag.BoolVar(&flush, "flush", false, "Delete all the keys!?")
}

func main() {
	flag.Parse()
	if flush {
		surat.Flush()
	} else {
		surat.Run()
	}

}
