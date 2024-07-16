package main

import (
	"flag"
	"fmt"
	"mime"
	"net"
	"os"
	"path/filepath"
	"strings"
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

	requestLine := strings.SplitN(string(req), "\r\n", 2)[0]
	fmt.Println("Request line:", requestLine)
	if !strings.HasPrefix(requestLine, "GET ") {
		fmt.Println("Invalid request method")
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\n"))
		return
	}

	requestedFile := strings.Fields(requestLine)[1]
	if requestedFile == "/" {
		requestedFile = "/index.html"
	}

	requestedFile = "." + filepath.Clean(requestedFile)

	if _, err := os.Stat(requestedFile); os.IsNotExist(err) {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}

	fileContent, err := os.ReadFile(requestedFile)
	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
		conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(requestedFile))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\nContent-Type: %s\r\n\r\n", contentType)
	conn.Write(fileContent)
}
