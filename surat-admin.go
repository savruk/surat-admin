package main

// import (
// 	"flag"

// 	"github.com/savruk/surat"
// )

// var flush bool

// func init() {
// 	flag.BoolVar(&flush, "flush", false, "Delete all the keys!?")
// }

// func main() {
// 	flag.Parse()
// 	if flush {
// 		surat.Flush()
// 	} else {
// 		surat.Run()
// 	}

// }

import (
	"flag"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
