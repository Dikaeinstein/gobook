package main

import (
	"io"
	"log"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	io.Copy(c, c)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		handleConn(conn)
	}
}
