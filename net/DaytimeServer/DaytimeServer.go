package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"

	listener, err := net.Listen("tcp", service)

	checkError(err)

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go func() {
			defer conn.Close()

			handleClient(conn)
		}()
	}
}

func handleClient(conn net.Conn) {
	daytime := time.Now()

	msg := fmt.Sprint(daytime)

	conn.Write([]byte(msg))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
