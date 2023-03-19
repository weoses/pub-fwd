package main

import (
	"encoding/gob"
	"log"
	"net"
)

func RunPubServer() {
	log.Printf("Pub server mode. In %s  Out %s", AddrIn, AddrOut)

	syncChan := make(chan int)
	go func() {
		runPubServer()
		syncChan <- 1
	}()

	<-syncChan

}

func runPubServer() {

	outListener, err := net.Listen("tcp", AddrOut)
	if err != nil {
		log.Fatalf("Error bind out address - %s", err.Error())
	}

	inListener, err := net.Listen("tcp", AddrIn)
	if err != nil {
		log.Fatalf("Error bind in address - %s", err.Error())
	}
	// wait 'OUT' connection from client

	syncChan := make(chan int)
	for {
		go func() { processListeners(inListener, outListener, syncChan) }()
		<-syncChan
	}

}

func processListeners(inListener net.Listener, outListener net.Listener, syncChan chan int) {
	completed := false

	defer func() {
		if !completed {
			syncChan <- 1
		}
	}()

	outConn, err := outListener.Accept()
	if err != nil {
		log.Printf("Error accept out connection - %s", err.Error())
		return
	}
	defer outConn.Close()
	log.Printf("Out connected from %s", outConn.RemoteAddr())

	inConn, err := inListener.Accept()
	if err != nil {
		log.Printf("Error accept in connection - %s", err.Error())
		return
	}
	defer inConn.Close()
	log.Printf("In connected from %s", inConn.RemoteAddr())

	decoder := gob.NewDecoder(outConn)
	encoder := gob.NewEncoder(outConn)

	hs := &cmd{}
	hs.Cmdtype = CONNECT
	err = encoder.Encode(hs)
	if err != nil {
		log.Printf("Error send CONNECT message - %s", err.Error())
		return
	}

	hsResp := &cmd{}
	err = decoder.Decode(hsResp)
	if err != nil {
		log.Printf("Error recv CONNECT response - %s", err.Error())
		return
	}

	if hsResp.Cmdtype == DISCONNECT {
		return
	}

	completed = true
	syncChan <- 1

	ResendForever(inConn, outConn)
}
