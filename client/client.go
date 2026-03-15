package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type clientMsg struct {
	opval string
	key   string
	value string
	ttl   string
}

func NewclientMsg() *clientMsg {
	return &clientMsg{}
}

func (c *clientMsg) ClientStart() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error in client connection")
	}

	reader := bufio.NewReader(conn)
	msg := ""
	fmt.Println("Type exit for exiting the code..")
	for msg != "exit" {
		readerval := bufio.NewReader(os.Stdin)
		msg, _ := readerval.ReadString('\n')
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Fatal("Error while reading the client msg")
		}
		strval, err := reader.ReadString('\n')
		fmt.Println("Server Response To client-->", strval)
		if strval == "BYE" {
			conn.Close()
		}

		if err != nil {
			log.Fatal("close")
		}
	}

}
