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
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		conn.SetDeadline(time.Now().Add(30 * time.Second))

		fmt.Println(string(buf[0:n]))

		_, err = conn.Write(buf[0:n])

		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
