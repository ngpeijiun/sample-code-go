package main

import (
	"encoding/asn1"
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

	mdata, err := asn1.Marshal(daytime)

	if err != nil {
		return
	}

	conn.Write(mdata)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
