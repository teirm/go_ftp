// bare bones client implementation
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

const (
	defaultAddress string = "127.0.0.1"
	defaultPort    string = "0"
)

func main() {
	ip := flag.String("address", defaultAddress, "address to connect to")
	port := flag.String("port", defaultPort, "port to connect to")
	flag.Parse()

	log.Printf("ip: %s port: %s\n", *ip, *port)

	address := *ip + ":" + *port

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Unable to connect to %s: %v\n", address, err)
	}

	fmt.Fprintf(conn, "TEST\n")

	if err := conn.Close(); err != nil {
		log.Fatalf("Failed to close connection: %v\n", err)
	}
}
