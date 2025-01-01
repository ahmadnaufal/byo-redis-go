package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	byoredisgo "github.com/ahmadnaufal/byo-redis-go"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3232")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// init loop for input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")

		// CLI will always be parsed in array of bulk strings
		tokens := strings.Split(input, " ")
		req := byoredisgo.Array{}
		for _, token := range tokens {
			req.Values = append(req.Values, &byoredisgo.BulkString{Value: token})
		}

		// send input to server
		n, err := conn.Write(req.Serialize())
		if err != nil {
			log.Println("Error sending message to server:", err.Error())
			continue
		}

		log.Printf("Sent %d bytes: %s\n", n, input)

		// read response from server
		buf := make([]byte, 1024)
		n, err = conn.Read(buf)
		if err != nil {
			log.Println("Error reading from server:", err.Error())
			continue
		}

		log.Printf("Received %d bytes\n", n)
		res, err := byoredisgo.Construct(buf[:n])
		if err != nil {
			log.Println("Error constructing response:", err.Error())
			continue
		}

		log.Println(res.String())
	}
}
