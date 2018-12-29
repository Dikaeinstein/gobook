// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

import "sync"

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New returns a new instance of a Memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get returns the response and error if any from the HTTP GET requests to the key
// Concurrency safe
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
