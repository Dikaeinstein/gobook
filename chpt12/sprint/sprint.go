package sprint

import (
	"reflect"
	"strconv"
)

// Sprint returns a string
func Sprint(x interface{}) string {
	reflect.ValueOf
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, map, chan, slice, pointer, struct, func
		return "???"
	}
}
