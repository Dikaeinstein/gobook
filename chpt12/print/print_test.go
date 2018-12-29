package print

import (
	"strings"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	cases := []struct{ in interface{} }{
		{time.Hour},
		{new(strings.Replacer)},
	}
	for _, c := range cases {
		Print(c.in)
	}
}
