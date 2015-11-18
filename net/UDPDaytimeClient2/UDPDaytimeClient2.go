/* UDPDaytimeClient2 localhost:1200
 */
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	checkArgs(os.Args)

	service := os.Args[1]

	conn, err := net.Dial("udp", service)

	checkError(err)

	_, err = conn.Write([]byte("anything"))

	checkError(err)

	var buf [512]byte

	n, err := conn.Read(buf[0:])

	checkError(err)

	fmt.Println(string(buf[0:n]))
}

func checkArgs(args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", args[0])
		os.Exit(1)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
