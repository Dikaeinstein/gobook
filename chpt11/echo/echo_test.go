package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	cases := []struct {
		newLine bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, " ", []string{}, "\n"},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
		{true, "", []string{"My name is Dika"}, "My name is Dika\n"},
	}
	for _, c := range cases {
		descr := fmt.Sprintf("echo(%v, %q, %q)",
			c.newLine, c.sep, c.args)
		out = new(bytes.Buffer) // captured output
		if err := echo(c.newLine, c.sep, c.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != c.want {
			t.Errorf("%s = %q, want %q", descr, got, c.want)
		}
	}
}
