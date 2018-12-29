package format

import (
	"os"
	"strconv"
	"testing"
)

var intCases = []struct {
	in   interface{}
	want string
}{
	{3, "3"},
	{-3, "-3"},
}

var stringCases = []struct {
	in   interface{}
	want string
}{
	{"string", strconv.Quote("string")},
}

func TestAny(t *testing.T) {
	for _, c := range intCases {
		if got := Any(c.in); got != c.want {
			t.Errorf("format.Any(%v) = %v; want %v", c.in, got, c.want)
		}
	}
	for _, c := range stringCases {
		if got := Any(c.in); got != c.want {
			t.Errorf("format.Any(%v) = %v; want %v", c.in, got, c.want)
		}
	}
}

func TestDisplay(t *testing.T) {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	Display("stranglove", strangelove)
	Display("os.Stderr", os.Stderr)
}
