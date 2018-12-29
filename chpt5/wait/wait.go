package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
	log.SetPrefix("wait: ")
	log.SetFlags(0)
}

func main() {
	err := WaitForServer("https://bad.gopl.io")
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// os.Exit(1)
		log.Fatalf("Site is down: %v\n", err)
	}
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 30 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
