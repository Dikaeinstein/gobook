package comma

import (
	"bytes"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer

	for i, v := range s {
		if i > 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
