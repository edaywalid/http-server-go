package main

import "flag"

func main() {
	port := flag.String("p", "8080", "Port to listen on")
	flag.Parse()
}
