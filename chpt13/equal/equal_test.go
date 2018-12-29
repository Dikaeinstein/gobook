package equal

import (
	"testing"
)

func TestEqual(t *testing.T) {
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	cases := []struct {
		x, y interface{}
		want bool
	}{
		{*a, *a, true},
		{*b, *b, true},
		{*c, *c, true},
		{*a, *b, false},
		{*a, *c, false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string(nil), []string{}, true},
		{map[string]int(nil), map[string]int{}, true},
	}
	for _, c := range cases {
		if got := Equal(c.x, c.y); got != c.want {
			t.Errorf("Equal(%v, %v) = %v", c.x, c.y, got)
		}
	}
}

func TestNumEqual(t *testing.T) {
	cases := []struct {
		x, y interface{}
		want bool
	}{
		{1, 1, true},
		{1.8, 1.9, true},
	}
	for _, c := range cases {
		if got := NumEqual(c.x, c.y); got != c.want {
			t.Errorf("NumEqual(%v, %v) = %v", c.x, c.y, got)
		}
	}
}
