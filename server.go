package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	port := flag.String("p", "8080", "Port to listen on")
	flag.Parse()

	address := fmt.Sprintf("0.0.0.0:%s", *port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Failed to bind to port %s\n", *port)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	req := make([]byte, 1024)
	_, err := conn.Read(req)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		return
	}

}
