package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type kvstore struct {
	entry map[string]string
	port  string
	mux   sync.Mutex
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
	defer conn.Close()
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
		kvstore.mux.Lock()
		kvstore.entry[result[1]] = result[2]
		kvstore.mux.Unlock()
		//fmt.Println("Server Set")
		fmt.Fprint(conn, "SET_OK\n")

	case "GET":
		kvstore.mux.Lock()
		_, ok := kvstore.entry[result[1]]
		kvstore.mux.Unlock()
		if !ok {
			fmt.Fprint(conn, "NOT_FOUND\n")
		} else {
			//fmt.Println("key val is", kvstore.entry[result[1]])
			fmt.Fprint(conn, "GET_OK\n")
			//fmt.Println("Entry at index ", result[1], " is", entry)
		}

	case "DEL":
		//fmt.Println("Delete Operation started for", result[1])
		kvstore.mux.Lock()
		_, ok := kvstore.entry[result[1]]
		kvstore.mux.Unlock()
		if !ok {
			fmt.Fprint(conn, "DEL_ERR\n")
		} else {
			kvstore.mux.Lock()
			delete(kvstore.entry, result[1])
			kvstore.mux.Unlock()
			fmt.Fprint(conn, "DELETE_OK\n")

		}
	}

}
