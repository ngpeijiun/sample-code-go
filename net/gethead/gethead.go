/* gethead www.google.com:80
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	checkArgs(os.Args)

	service := os.Args[1]

	addr, err := net.ResolveTCPAddr("tcp", service)

	checkError(err)

	conn, err := net.DialTCP("tcp", nil, addr)

	checkError(err)

	defer conn.Close()

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))

	checkError(err)

	result, err := ioutil.ReadAll(conn)

	checkError(err)

	fmt.Println(string(result))
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
