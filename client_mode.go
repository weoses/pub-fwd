package main

import (
	"encoding/gob"
	"log"
	"net"
	"time"
)

func RunClientServer() {
	log.Printf("Client server mode. In %s  Out %s", AddrIn, AddrOut)

	syncChan := make(chan int)
	go func() {
		runClientServer()
		syncChan <- 1
	}()

	<-syncChan
}

func runClientServer() {
	syncChan := make(chan int)
	for {
		go func() { processClient(syncChan) }()
		<-syncChan
	}
}
func processClient(syncChan chan int) {

	completed := false

	defer func() {
		if !completed {
			syncChan <- 1
		}
	}()

	inConn, err := net.Dial("tcp", AddrIn)
	if err != nil {
		log.Printf("Cant connect to in server - %s", err.Error())
		time.Sleep(350 * time.Millisecond)
		return
	}
	defer inConn.Close()

	log.Printf("Connected to in server - %s", inConn.RemoteAddr())

	c_in := &cmd{}

	encoder := gob.NewEncoder(inConn)
	decoder := gob.NewDecoder(inConn)

	err = decoder.Decode(c_in)
	if err != nil {
		log.Printf("Cant parse CONNECT from in server - %s", err.Error())
		return
	}

	if c_in.Cmdtype != CONNECT {
		return
	}

	outConn, err := net.Dial("tcp", AddrOut)
	if err != nil {
		log.Printf("Cant connect to out server - %s", err.Error())
		return
	}

	defer outConn.Close()

	c_out := &cmd{}
	c_out.Cmdtype = CONNECT
	err = encoder.Encode(c_out)
	if err != nil {
		log.Printf("Cant send CONNECT to in server - %s", err.Error())
		return
	}

	completed = true
	syncChan <- 1

	ResendForever(inConn, outConn)
}
