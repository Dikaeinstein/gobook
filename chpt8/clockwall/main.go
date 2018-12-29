package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	locAddr := parseArgs()
	for l, a := range locAddr {
		clockClient(l, a)
	}
}

type Location string
type Addr string

func clockClient(l Location, a Addr) {
	conn, err := net.Dial("tcp", string(a))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func parseArgs() map[Location]Addr {
	args := os.Args[1:]
	location := make(map[Location]Addr)
	for _, a := range args {
		locAddr := strings.Split(a, "=")
		location[Location(locAddr[0])] = Addr(locAddr[1])
	}
	return location
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func printClockWall() {
	const format = "%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 3, ' ', 0)
	fmt.Fprintf(tw, format, "Location", "Clock")
	fmt.Fprintf(tw, format, "--------", "-----")

}
