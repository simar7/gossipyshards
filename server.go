package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Listener int

func (l *Listener) GetLine(line []byte, ack *bool) error {
	fmt.Println(string(line))
	return nil
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8001")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
