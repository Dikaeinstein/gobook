package basename

import (
	"strings"
)

func basename(s string) string {
	// Discard everything before last '/'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s += s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s += s[:i]
			break
		}
	}

	return s
}

// Simpler version using strings.lastIndex
func basename2(s string) string {
	i := strings.LastIndex(s, "/")
	s = s[i+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}
