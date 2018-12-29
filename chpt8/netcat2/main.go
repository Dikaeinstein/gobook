package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var port = flag.String("port", "8000", "port to connect to")

func main() {
	flag.Parse()
	address := fmt.Sprintf("localhost:%s", *port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
