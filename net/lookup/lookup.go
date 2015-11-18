/* lookup www.google.com
 */
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	checkArgs(os.Args)

	hostname := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", hostname)

	checkError(err)

	mask := addr.IP.DefaultMask()
	ones, bits := mask.Size()
	network := addr.IP.Mask(mask)

	fmt.Println("Address is", addr)
	fmt.Println("Default mask length is", bits)
	fmt.Println("Leading ones count is", ones)
	fmt.Println("Mask is (hex)", mask)
	fmt.Println("Network is", network)

	addrs, err := net.LookupHost(hostname)

	checkError(err)

	fmt.Println()

	for _, s := range addrs {
		fmt.Println(s)
	}
}

func checkArgs(args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", args[0])
		os.Exit(1)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
