package eval

import (
	"fmt"
	"math"
	"testing"
)

// TestEval tests the eval package
// func TestEval(t *testing.T) {
// 	cases := []struct {
// 		expr string
// 		env  Env
// 		want string
// 	}{
// 		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
// 		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
// 		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
// 		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
// 		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
// 		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
// 	}
// 	var prevExpr string
// 	for _, c := range cases {
// 		// Print expr only when it changes.
// 		if c.expr != prevExpr {
// 			fmt.Printf("\n%s\n", c.expr)
// 			prevExpr = c.expr
// 		}
// 		expr, err := Parse(c.expr)
// 		if err != nil {
// 			t.Error(err) // parse error
// 			continue
// 		}
// 		got := fmt.Sprintf("%.6g", expr.Eval(c.env))
// 		fmt.Printf("\t%v => %s\n", c.env, got)
// 		if got != c.want {
// 			t.Errorf("%s.Eval() in %v = %q, want %q\n",
// 				c.expr, c.env, got, c.want)
// 		}
// 	}
// }

func TestCoverage(t *testing.T) {
	var tests = []struct {
		input string
		env   Env
		want  string // expected error from Parse/Check or result from Eval
	}{
		{"x % 2", nil, "unexpected '%'"},
		{"!true", nil, "unexpected '!'"},
		{"log(10)", nil, `unknown function "log"`},
		{"sqrt(1, 2)", nil, "call to sqrt has 2 args, want 1"},
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"-x * -x", Env{"x": 2}, "4"},
		{"-x * +x", Env{"x": 2}, "-4"},
		{"~x * x", Env{"x": 2}, "unexpected operator '~'"},
	}
	for _, test := range tests {
		expr, err := Parse(test.input)
		if err == nil {
			err = expr.Check(map[Var]bool{})
		}
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("%s: got %q, want %q", test.input, err, test.want)
			}
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%s: %v => %s, want %s",
				test.input, test.env, got, test.want)
		}
	}
}
