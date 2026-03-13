package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ready := make(chan struct{})
	portno := ":8080"
	//clientRequest := "SET key1 Hello SET key2 Hello SET key3 Hello\nGET key3\nGET key2\nDELETE key1\n"
	kvstore := Newkvstore(portno)

	wg.Add(3)
	//Server goroutine
	go func() {
		kvstore.Start(ready)
	}()
	<-ready
	//Client1 goroutine
	go func() {
		defer wg.Done()
		client1 := NewclientMsg("SET key1 Hello\nSET key2 BYE\nSET key3 newkeyval\nGET key9\nGET key2")
		client1.ClientStart()
	}()
	//Client2 goroutine
	go func() {
		defer wg.Done()
		client2 := NewclientMsg("SET key4 Saionara\nSET key5 HastalaVista\nSET key6 Salom\nDEL key1")
		client2.ClientStart()
	}()
	//Client3 goroutine
	go func() {
		defer wg.Done()
		client3 := NewclientMsg("SET key7 Dhoom\nDEL key10\nDEL key1\nDEL key2")
		client3.ClientStart()
	}()
	wg.Wait()
	fmt.Println(kvstore.entry)
	fmt.Println("Main exited")

}
