package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func main() {
	ready := make(chan bool)
	go server(ready)
	<-ready // Wait for the server to signal readiness
	go client()
	var input string
	fmt.Scanln(&input)
}

func server(ready chan<- bool) {
	ln, err := net.Listen("tcp", ":9889")
	if err != nil {
		fmt.Println(err)
		return
	}
	ready <- true // Signal that the server is ready
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(c)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	var msg string
	err := gob.NewDecoder(connection).Decode(&msg)
	if err != nil {
		fmt.Println("Error decoding:", err)
	} else {
		fmt.Println("Received:", msg)
	}
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9889")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	msg := "Madhvesh"
	fmt.Println("Sending:", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("Error encoding:", err)
	}
}
