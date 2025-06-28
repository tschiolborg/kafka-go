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
	c, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Successfully bound to port 9092 and accepted a connection")

	// Make sure to use buf[:n] slice to avoid reading stale data
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
			break
		}
		if n > 0 {
			fmt.Printf("Received data: %s\n", string(buf[:n]))
		}
	}

}
