package calc

import (
	"testing"
)

func TestPrecedence(t *testing.T) {
	tests := []struct {
		op       rune
		expected int
	}{
		{'+', 1},
		{'-', 1},
		{'*', 2},
		{'/', 2},
		{'(', 0},
		{')', 0},
		{'^', -1},
	}

	for _, test := range tests {
		got := precedence(test.op)
		if got != test.expected {
			t.Errorf("precedence(%c) = %d; want %d", test.op, got, test.expected)
		}
	}
}

func TestApplyOperator(t *testing.T) {
	tests := []struct {
		left     float64
		right    float64
		op       rune
		expected float64
		hasError bool
	}{
		{3, 4, '+', 7, false},
		{10, 5, '-', 5, false},
		{6, 7, '*', 42, false},
		{8, 2, '/', 4, false},
		{8, 0, '/', 0, true},
		{5, 3, '^', 0, true},
	}

	for _, test := range tests {
		got, err := applyOperator(test.left, test.right, test.op)
		if test.hasError && err == nil {
			t.Errorf("applyOperator(%f, %f, %c) expected error; got nil", test.left, test.right, test.op)
		} else if !test.hasError && err != nil {
			t.Errorf("applyOperator(%f, %f, %c) unexpected error: %v", test.left, test.right, test.op, err)
		} else if !test.hasError && got != test.expected {
			t.Errorf("applyOperator(%f, %f, %c) = %f; want %f", test.left, test.right, test.op, got, test.expected)
		}
	}
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		expression string
		expected   []string
		hasError   bool
	}{
		{"3 + 4", []string{"3", "+", "4"}, false},
		{"10 - 5", []string{"10", "-", "5"}, false},
		{"6 * 7", []string{"6", "*", "7"}, false},
		{"8 / 2", []string{"8", "/", "2"}, false},
		{"(3 + 4) * 5", []string{"(", "3", "+", "4", ")", "*", "5"}, false},
		{"8 / 0", []string{"8", "/", "0"}, false},
		{"5 ^ 3", nil, true},
		{"3 + a", nil, true},
		{"", []string{}, false},
	}

	for _, test := range tests {
		got, err := tokenize(test.expression)
		if test.hasError && err == nil {
			t.Errorf("tokenize(%q) expected error; got nil", test.expression)
		} else if !test.hasError && err != nil {
			t.Errorf("tokenize(%q) unexpected error: %v", test.expression, err)
		} else if !test.hasError && !equal(got, test.expected) {
			t.Errorf("tokenize(%q) = %v; want %v", test.expression, got, test.expected)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		error      bool
		expression string
		result     float64
	}{
		{error: false, expression: "1 +  52 ", result: 53},
		{error: true, expression: "6!", result: 0},
		{error: false, expression: "1+1", result: 2},
		{error: false, expression: "-1+2", result: 1},
		{error: false, expression: "-1*5", result: -5},
		{error: true, expression: "-3^2", result: 0},
		{error: true, expression: "-3^3", result: 0},
		{error: false, expression: "(2+2)*2", result: 8},
		{error: false, expression: "2+2*2", result: 6},
		{error: false, expression: "1/2", result: 0.5},
		{error: true, expression: "2+2^2", result: 0},
		{error: true, expression: "10 % 1000", result: 0},
		{error: true, expression: "10%", result: 0},
		{error: true, expression: "-1!", result: 0},
		{error: true, expression: "2+2**2", result: 0},
		{error: true, expression: "1+1*", result: 0},
		{error: true, expression: "((2+2-*(2", result: 0},
		{error: true, expression: "", result: 0},
		{error: true, expression: "1/0", result: 0},
		{error: true, expression: "das/52", result: 0},
		{error: true, expression: "WAS+d", result: 0},
	}

	for _, test := range tests {
		val, err := Solve(test.expression)
		if err != nil && !test.error {
			t.Fatalf("Solve(%q) returned unexpected error: %s", test.expression, err)
		}
		if err == nil && test.error {
			t.Fatalf("Solve(%q) returned no error", test.expression)
		}
		if val != test.result {
			t.Fatalf("Solve(%q) = %f; want %f", test.expression, val, test.result)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
