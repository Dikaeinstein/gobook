// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New returns a new instance of a Memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get returns the response and error if any from the HTTP GET requests to the key
func (m *Memo) Get(key string) (interface{}, error) {
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	return res.value, res.err
}
