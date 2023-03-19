package main

import "flag"

var pub bool
var AddrIn string
var AddrOut string

func init() {
	flag.BoolVar(&pub, "public", false, "Public server mode")
	flag.StringVar(&AddrIn, "in", "localhost:8010", "Bind in addr")
	flag.StringVar(&AddrOut, "out", "localhost:8020", "Bind in addr")
}

func main() {
	flag.Parse()
	if pub {
		RunPubServer()
		return
	} else {
		RunClientServer()
	}

}
