// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// FileLine type
type FileLine struct {
	count int
	name  string
}

func countLines(f *os.File, counts map[string]FileLine) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		fileLine := FileLine{
			counts[key].count + 1,
			f.Name(),
		}
		counts[key] = fileLine
	}
	// NOTE: ignoring potential errors from input.Err()
}

func dup2() {
	counts := make(map[string]FileLine)
	args := os.Args[1:]
	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, fileLine := range counts {
		if fileLine.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", fileLine.count, line, fileLine.name)
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	args := os.Args
	for _, fileName := range args {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	// dup()
	dup2()
	// dup3()
}
