package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Result struct {
	// Dir         string
	// ImportPath  string
	// Name        string
	// Target      string
	// Root        string
	// Match       []string
	// Stale       bool
	// StaleReason string
	// GoFiles     []string
	// Imports     []string
	Deps []string
}

func main() {
	args := os.Args[1:]
	cmd := exec.Command("go", "list", "-json", ".")
	// var out bytes.Buffer
	// cmd.Stdout = &out
	var data []byte
	var err error
	if data, err = cmd.Output(); err != nil {
		log.Fatal(err)
	}
	var result *Result
	r := bytes.NewReader(data)
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		fmt.Println(err)
	}
	for _, d := range result.Deps {
		a := strings.Join(args, ",")
		if strings.Contains(a, d) {
			m, _ := json.MarshalIndent(result.Deps, "", "    ")
			fmt.Println(string(m))
			return
		}
	}
	fmt.Println("Package does not match any dependency in the workspace")
}
