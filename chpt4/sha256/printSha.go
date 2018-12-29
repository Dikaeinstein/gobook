package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var shaPtr = flag.String("s", "sha256", "sha digest")

func printSha() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("message is required to create digest")
		return
	}
	message := os.Args[1]

	switch *shaPtr {
	case "sha256":
		c := sha256.Sum256([]byte(message))
		fmt.Printf("%x\n", c)
	case "sha384":
		c := sha512.Sum384([]byte(message))
		fmt.Printf("%x\n", c)
	case "sha512":
		c := sha512.Sum512([]byte(message))
		fmt.Printf("%x\n", c)
	default:
		fmt.Println("invalid sha!")
	}
}
