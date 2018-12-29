package comma

import (
	"testing"
)

func TestComma(t *testing.T) {
	cases := []struct {
		in, want string
	}{{"12345", "12,345"}, {"123456789", "123,456,789"}}

	for _, c := range cases {
		got := comma(c.in)
		if got != c.want {
			t.Errorf("comma(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestComma2(t *testing.T) {
	cases := []struct {
		in, want string
	}{{"12345", "12,345"}, {"123456789", "123,456,789"}}

	for _, c := range cases {
		got := comma2(c.in)
		if got != c.want {
			t.Errorf("comma(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
