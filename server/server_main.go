package main

func main() {
	//chanval := make(chan struct{})
	serverinst := Newkvstore(":8080")
	serverinst.Start()
	//<-chanval
}
