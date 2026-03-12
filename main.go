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
	clientMsg := NewclientMsg("SET key1 Hello\nSET key2 BYE\nSET key3 newkeyval\nGET key1\nGET key2\nDEL key2")

	wg.Add(1)
	//Server goroutine
	go func() {
		kvstore.Start(ready)
	}()
	<-ready
	//Client goroutine
	go func() {
		defer wg.Done()
		clientMsg.ClientStart()
	}()
	wg.Wait()
	fmt.Println(kvstore.entry)
	fmt.Println("Main exited")

}
