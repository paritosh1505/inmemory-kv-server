package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type kvstore struct {
	entry map[string]string
	port  string
}

func Newkvstore(port string) *kvstore {
	return &kvstore{
		entry: make(map[string]string),
		port:  port,
	}
}

func (k *kvstore) Start(chanval chan struct{}) {
	listener, err := net.Listen("tcp", k.port)
	if err != nil {
		log.Fatal("error while opening the port")
	}
	fmt.Println("Server connection started")
	close(chanval)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error while connecting the client")
		}
		go k.handleConn(conn)
	}

}
func (k *kvstore) handleConn(conn net.Conn) {
	//defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}
		k.DataStorage(msg, conn)
	}

}
func (kvstore *kvstore) DataStorage(msg string, conn net.Conn) {
	trimMsg := strings.TrimSpace(msg)
	result := strings.Fields(trimMsg)

	operation := result[0]
	switch operation {
	case "SET":
		kvstore.entry[result[1]] = result[2]

		fmt.Println("Server Set")
		fmt.Fprint(conn, "OK\n")
	case "GET":
		fmt.Println("result val,=>", result[1])
		result[1] = strings.ReplaceAll(result[1], "\n", "")
		entry := kvstore.entry[result[1]]
		fmt.Println("key val is", kvstore.entry[result[1]])
		fmt.Fprint(conn, "OK\n")

		fmt.Println("Entry at index ", result[1], " is", entry)
	case "DEL":
		fmt.Println("Delete Operation started")
		delete(kvstore.entry, result[1])
	}

}
