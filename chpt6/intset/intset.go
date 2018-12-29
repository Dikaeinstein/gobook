package intset

import (
	"bytes"
	"fmt"
)

// IntSet is a set of small non-negative integers.
// Its zero value represents an empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len return the number of elements
func (s *IntSet) Len() int {
	length := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			// Check if bit is set
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return length
}

// Remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	// clear bit: invert bit ^(bit) then AND bit
	s.words[word] &^= (1 << uint(bit))
}

// Clear all elements from the set
func (s *IntSet) Clear() {
	for i := 0; i < len(s.words); i++ {
		for j := 0; j < 64; j++ {
			s.words[i] &^= (1 << uint(j))
		}
	}
}

// Copy of the set
func (s *IntSet) Copy() *IntSet {
	t := IntSet{words: []uint64{}}
	for _, word := range s.words {
		t.words = append(t.words, word)
	}
	return &t
}

// AddAll allows a list of values to be added, such as s.AddAll(1, 2, 3).
func (s *IntSet) AddAll(args ...int) {
	for _, arg := range args {
		s.Add(arg)
	}
}

// IntersectWith sets s to the intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// Elem returns a slice containing the elements of the set,
// suitable for iterating over with a range loop.
func (s *IntSet) Elem() []int {
	elements := []int{}
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elements = append(elements, 64*i+j)
			}
		}
	}
	return elements
}
