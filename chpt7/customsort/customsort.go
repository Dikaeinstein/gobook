package customsort

import (
	"github.com/dikaeinstein/gobook/chpt7/sorting"
)

type customSort struct {
	t    []*sorting.Track
	less func(x, y *sorting.Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
