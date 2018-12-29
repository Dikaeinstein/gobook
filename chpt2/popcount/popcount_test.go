package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(2048)
	}
}

func BenchmarkPCLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PCLoop(2048)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(2048)
	}
}

func TestPopCount(t *testing.T) {
	cases := []struct {
		in   uint64
		want int
	}{{2, 1}, {3, 2}}

	for _, c := range cases {
		got := PopCount(c.in)
		if got != c.want {
			t.Errorf("PopCount(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPCLoop(t *testing.T) {
	cases := []struct {
		in   uint64
		want int
	}{{2, 1}, {3, 2}}

	for _, c := range cases {
		got := PCLoop(c.in)
		if got != c.want {
			t.Errorf("PCLoop(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPopCount2(t *testing.T) {
	cases := []struct {
		in   uint64
		want int
	}{{2, 1}, {3, 2}}

	for _, c := range cases {
		got := PopCount2(c.in)
		if got != c.want {
			t.Errorf("PopCount2(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
