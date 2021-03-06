/* ASN1DaytimeClient localhost:1200
 */
package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {
	checkArgs(os.Args)

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)

	checkError(err)

	defer conn.Close()

	result, err := ioutil.ReadAll(conn)

	checkError(err)

	var daytime time.Time

	_, err = asn1.Unmarshal(result, &daytime)

	checkError(err)

	fmt.Println("After marshal / unmarshal:", daytime)
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
