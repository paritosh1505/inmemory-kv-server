package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type kvstore struct {
	entry       map[string]string
	port        string
	mux         sync.Mutex
	clientCount int
}

func Newkvstore(port string) *kvstore {
	return &kvstore{
		entry: make(map[string]string),
		port:  port,
	}
}

func (k *kvstore) Start() {
	listener, err := net.Listen("tcp", k.port)
	if err != nil {
		log.Fatal("error while opening the port")
	}
	fmt.Println("Server connection started")
	for {
		conn, err := listener.Accept()
		k.mux.Lock()
		k.clientCount++
		fmt.Println("Client connected to server..", k.clientCount)
		k.mux.Unlock()
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
		go k.CleanExpiredKey(msg, conn)
	}

}
func (kvstore *kvstore) DataStorage(msg string, conn net.Conn) {
	trimMsg := strings.TrimSpace(msg)
	result := strings.Fields(trimMsg)
	operation := result[0]
	operation = strings.ToUpper(operation)
	switch operation {
	case "SET":
		fmt.Println("Set Called in Server..")
		kvstore.mux.Lock()
		kvstore.entry[result[1]] = result[2]
		kvstore.mux.Unlock()
		fmt.Fprint(conn, "SET_OK\n")

	case "GET":
		fmt.Println("GET called in server..")
		kvstore.mux.Lock()
		value, ok := kvstore.entry[result[1]]
		kvstore.mux.Unlock()
		if !ok {
			fmt.Fprint(conn, "NOT_FOUND\n")
		} else {
			fmt.Fprintf(conn, "GET_OK %s\n", value)
		}

	case "DEL":
		fmt.Println("DEL called in server..")
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
	case "EXIT":
		fmt.Fprint(conn, "BYE\n")
		kvstore.clientCount--
		conn.Close()
	}

}

func (k *kvstore) CleanExpiredKey(msg string, conn net.Conn) {
	trimmsg := strings.TrimSpace(msg)
	msgArr := strings.Fields(trimmsg)
	fmt.Println("len of msg is", len(msgArr))
	if len(msgArr) == 4 {
		chanval := make(chan struct{})

		fmt.Println("************", msgArr)
		if msgArr[3] != "" {
			timer, err := strconv.Atoi(msgArr[3])
			if err != nil {
				log.Fatal("Error while converting string timer to integer timer")
			}
			go func() {
				time.Sleep(time.Duration(timer) * time.Microsecond)
				chanval <- struct{}{}
			}()
			<-chanval
			delete(k.entry, msgArr[1])
			fmt.Fprintf(conn, "Key Deleted after %d microsecond\n", timer)
		}
	}

}
