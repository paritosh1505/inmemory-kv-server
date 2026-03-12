package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type clientMsg struct {
	msg string
}

func NewclientMsg(msgC string) *clientMsg {
	return &clientMsg{
		msg: msgC,
	}
}

func (c *clientMsg) ClientStart() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error in client connection")
	}
	defer conn.Close()
	conn.Write([]byte(c.msg))
	reader := bufio.NewReader(conn)
	strval, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("close")
	}
	fmt.Println("**********", strval)
	fmt.Println("Client connection started")
}
