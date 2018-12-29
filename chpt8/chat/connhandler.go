package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

func handleConn(conn net.Conn) {
	ch := make(chan string, 4) // Buffered outgoing client messages channel, so most messages are not dropped when writer is not ready to accept a message
	done := make(chan struct{})
	uCli := userClient{c: ch}
	go clientWriter(conn, uCli.c)
	input := bufio.NewScanner(conn)

	who := ""
	fmt.Fprintln(conn, "Please enter username: ") // Block for user to enter name
	if ok := input.Scan(); ok {
		who = input.Text()
	}
	if who == "" {
		who = conn.RemoteAddr().String() // Fall back to user addr
	}

	uCli.name = who
	uCli.c <- "You are " + who
	messages <- uCli.name + " has arrived"
	entering <- uCli

	var heartbeat int
	var mu sync.Mutex
	go func(conn net.Conn) {
		tick := time.NewTicker(3 * time.Minute)
	loop:
		for {
			select {
			case <-tick.C:
				if heartbeat <= 0 {
					tick.Stop()
					conn.Close()
					return
				}
				mu.Lock()
				heartbeat = 0
				mu.Unlock()
			case <-done:
				tick.Stop() // Stop ticker and release resource before stoping goroutine
				break loop
			}
		}
	}(conn)
	for input.Scan() {
		messages <- uCli.name + ": " + input.Text()
		mu.Lock()
		heartbeat++
		mu.Unlock()
	}
	// NOTE: ignoring potential errors from input.Err()
	leaving <- uCli
	messages <- who + " has left"
	close(done)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
