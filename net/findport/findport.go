/* findport tcp https
 */
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	checkArgs(os.Args)

	networkType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(networkType, service)

	checkError(err)

	fmt.Println("Service port", port)
}

func checkArgs(args []string) {
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type service\n", args[0])
		os.Exit(1)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
