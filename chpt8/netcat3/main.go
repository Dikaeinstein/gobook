package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var port = flag.Int("port", 4000, "port to connect to")

func main() {
	flag.Parse()
	var rAddr net.TCPAddr
	rAddr.IP = net.ParseIP("127.0.0.1")
	rAddr.Port = *port
	conn, err := net.DialTCP("tcp4", nil, &rAddr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Print("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
