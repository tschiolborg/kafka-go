package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Successfully bound to port 9092 and accepted a connection")

		go handle(c)

	}
}

func handle(c net.Conn) {
	defer c.Close()
	fmt.Println("Handling connection")

	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
			break
		}
		if n <= 0 {
			continue
		}
		received := buf[:n]
		fmt.Printf("Received data: %s\n", string(received))

		messageSize := []byte{0, 0, 0, 0}
		correlationId := []byte{0, 0, 0, 7}

		response := append(messageSize, correlationId...)
		_, err = c.Write(response)
		if err != nil {
			fmt.Println("Error writing to connection: ", err.Error())
			break
		}
	}
}
