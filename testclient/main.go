package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3232")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// send a test string
	n, err := conn.Write([]byte("ECHO\r\nHELLO"))
	if err != nil {
		panic(err)
	}
	log.Printf("Sent %d bytes\n", n)

	// read the response
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
		return
	}

	log.Printf("Received %d bytes\n", n)
	fmt.Println(string(buf[:n]))
}
