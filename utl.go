package main

import (
	"io"
	"log"
	"net"
)

func ResendForever(cn1 net.Conn, cn2 net.Conn) {
	log.Printf("Create copy %s <-> %s", cn1.RemoteAddr(), cn2.RemoteAddr())
	defer log.Printf("Close copy %s <-> %s", cn1.RemoteAddr(), cn2.RemoteAddr())

	end := make(chan int)

	copier := func(from net.Conn, to net.Conn) {
		io.Copy(from, to)
		end <- 1
	}

	go copier(cn1, cn2)
	go copier(cn2, cn1)

	<-end
}
