package main

const (
	CONNECT = iota
	DATA
	DISCONNECT
)

type cmd struct {
	Cmdtype int
}
