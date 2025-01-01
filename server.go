package byoredisgo

import (
	"io"
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
			continue
		}

		// this line handles the requests
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Println("Accepted new connection from", conn.RemoteAddr())
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection from %s closed.\n", conn.RemoteAddr())
				return
			}

			log.Println("Error reading:", err.Error())
			continue
		}

		log.Printf("Received %d bytes\n", n)

		var bResponse []byte
		res, err := handlePayload(buf[:n])
		if err != nil {
			log.Println("Error handling command:", err.Error())
			bResponse = []byte(err.Error())
		} else {
			bResponse = res.Serialize()
		}

		nr, _ := conn.Write(bResponse)
		log.Printf("Written %d bytes back as reply\n", nr)
	}
}
