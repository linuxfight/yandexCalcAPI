package calc

import (
	"fmt"
	"testing"
)

type test struct {
	Expression string
	Result     float64
	Error      bool
}

func TestSolve(t *testing.T) {
	testCases := []test{
		{Error: false, Expression: "1 +  52 ", Result: 53},
		{Error: false, Expression: "6!", Result: 720},
		{Error: false, Expression: "10 % 1000", Result: 100},
		{Error: false, Expression: "1+1", Result: 2},
		{Error: false, Expression: "(2+2)*2", Result: 8},
		{Error: false, Expression: "2+2*2", Result: 6},
		{Error: false, Expression: "1/2", Result: 0.5},
		{Error: false, Expression: "2+2^2", Result: 6},
		{Error: true, Expression: "2+2**2", Result: 0},
		{Error: true, Expression: "1+1*", Result: 0},
		{Error: true, Expression: "((2+2-*(2", Result: 0},
		{Error: true, Expression: "", Result: 0},
		{Error: true, Expression: "1/0", Result: 0},
		{Error: true, Expression: "das/52", Result: 0},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test-%d: %s", i, testCase.Expression), func(t *testing.T) {
			val, err := Solve(testCase.Expression)
			if err != nil && !testCase.Error {
				t.Fatalf("Solve(%q) returned unexpected error: %s", testCase.Expression, err)
			}
			if err == nil && testCase.Error {
				t.Fatalf("Solve(%q) returned no error", testCase.Expression)
			}
			if val != testCase.Result {
				t.Fatalf("Solve(%q) = %f; want %f", testCase.Expression, val, testCase.Result)
			}
		})
	}
}
