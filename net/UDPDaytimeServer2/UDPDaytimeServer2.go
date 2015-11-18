package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"

	conn, err := net.ListenPacket("udp", service)

	checkError(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn net.PacketConn) {
	var buf [512]byte

	_, addr, err := conn.ReadFrom(buf[0:])

	if err != nil {
		return
	}

	daytime := time.Now()

	msg := fmt.Sprint(daytime)

	conn.WriteTo([]byte(msg), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
