package anagram

import "testing"

func TestAnagram(t *testing.T) {
	cases := []struct {
		in   [2]string
		want bool
	}{{[2]string{"aba", "aab"}, true}}

	for _, c := range cases {
		got := anagram(c.in[0], c.in[1])
		if got != c.want {
			t.Errorf("anagram(%q, %q) == %t, want %t", c.in[0], c.in[1], got, c.want)
		}
	}
}
