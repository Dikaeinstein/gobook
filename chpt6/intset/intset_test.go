package intset

import (
	"sort"
	"testing"
)

var hasTest = []struct {
	in   int
	want bool
}{
	{in: 1, want: true},
	{in: 2, want: true},
	{in: 3, want: false},
}

var AddTest = []struct {
	in   int
	want bool
}{
	{in: 1, want: true},
	{in: 2, want: true},
	{in: 3, want: true},
}

// var stringTest = []struct{
// 	in []int
// 	want string
// }{
// 	{in: []int{1, 2, 3, 4}, want: "{1, 2, 3, 4}"},
// 	{in: []int{1, 2, 3, 4}, want: "{1, 2, 3, 4}"},
// }

func TestHas(t *testing.T) {
	intSet := IntSet{words: []uint64{}}
	intSet.Add(1)
	intSet.Add(2)
	for _, c := range hasTest {
		if got := intSet.Has(c.in); got != c.want {
			t.Errorf("intSet.Has(%v) = %v", c.in, got)
		}
	}
}

func TestAddAll(t *testing.T) {
	intSet := IntSet{words: []uint64{}}
	intSet.AddAll(1, 2, 3, 2)
	for _, c := range AddTest {
		if got := intSet.Has(c.in); got != c.want {
			t.Errorf("intSet.AddAll(%v, %v, %v) failed!", 1, 2, 3)
		}
	}
}

func TestLen(t *testing.T) {
	intSet := IntSet{words: []uint64{}}
	intSet.AddAll(1, 2, 3)
	if got := intSet.Len(); got != 3 {
		t.Errorf("intSet.Len() = %v", got)
	}
}

func TestString(t *testing.T) {
	intSet := IntSet{words: []uint64{}}
	intSet.AddAll(1, 2, 3)
	intSet.AddAll(64, 84)
	if got := intSet.String(); got != "{1 2 3 64 84}" {
		t.Errorf("intSet.String() = %v", got)
	}
}

func TestElem(t *testing.T) {
	intSet := IntSet{words: []uint64{}}
	intSet.AddAll(1, 2, 3)
	want := []int{1, 2, 3}
	got := intSet.Elem()
	for _, e := range got {
		if sort.SearchInts(want, e) >= len(got) {
			t.Errorf("intSet.Elem() = %v; want : %v", got, want)
		}
	}
}

func TestIntSet(t *testing.T) {
	m := make(map[int]bool)
	intSet := IntSet{words: []uint64{}}
	m[1] = true
	intSet.Add(1)
	if m[1] != intSet.Has(1) {
		t.Errorf("intSet.Has(%v) != m[%v]", 1, 1)
	}
}
