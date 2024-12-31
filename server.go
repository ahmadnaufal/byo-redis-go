package main

import (
	"log"
	"net"
)

func StartServer() error {
	ln, err := net.Listen("tcp", ":3232")
	if err != nil {
		return err
	}

	log.Println("Listening for connections at :3232...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err.Error())
		}
		// this line handles the requests
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
		return
	}

	payload := string(buf[:n])
	log.Printf("Received %d bytes: %s\n", n, payload)

	var bResponse []byte
	res, err := handlePayload(payload)
	if err != nil {
		log.Println("Error handling command:", err.Error())
		bResponse = []byte(err.Error())
	} else {
		bResponse = []byte(res)
	}

	nr, _ := conn.Write(bResponse)
	log.Printf("Written %d bytes back as reply\n", nr)

}
