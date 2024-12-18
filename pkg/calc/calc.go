package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Helper function to check operator precedence
func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	case '!':
		return 4
	case '(', ')':
		return 0
	default:
		return -1
	}
}

// Helper function to calculate factorial
func factorial(n int) float64 {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return float64(result)
}

// Evaluate a simple mathematical expression
func applyOperator(left, right float64, op rune) (float64, error) {
	switch op {
	case '+':
		return left + right, nil
	case '-':
		return left - right, nil
	case '*':
		return left * right, nil
	case '/':
		if right == 0 {
			return 0, errors.New("division by zero")
		}
		return left / right, nil
	case '^':
		return math.Pow(left, right), nil // Exponentiation
	case '!':
		if left != float64(int(left)) || left < 0 {
			return 0, errors.New("factorial is only defined for non-negative integers")
		}
		return factorial(int(left)), nil // Factorial (only for left)
	default:
		return 0, fmt.Errorf("unsupported operation: %c", op)
	}
}

// Tokenizes the input expression
func tokenize(expression string) ([]string, error) {
	var tokens []string
	var current string
	for _, r := range expression {
		if r == ' ' {
			continue
		}
		if r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')' || r == '^' || r == '!' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, string(r))
		} else if (r >= '0' && r <= '9') || r == '.' { // Handle numbers (including float)
			current += string(r)
		} else {
			return nil, errors.New("invalid character in expression")
		}
	}
	if current != "" {
		tokens = append(tokens, current)
	}
	return tokens, nil
}

func Solve(expression string) (float64, error) {
	var output []float64
	var operatorStack []rune

	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if num, err := strconv.ParseFloat(token, 64); err == nil {
			output = append(output, num)
		} else if len(token) == 1 && strings.ContainsRune("+-*/()^!", rune(token[0])) {
			op := rune(token[0])
			if op == '(' {
				operatorStack = append(operatorStack, op)
			} else if op == ')' {
				// Process until we find a '('
				for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != '(' {
					right := output[len(output)-1]
					output = output[:len(output)-1]
					left := output[len(output)-1]
					output = output[:len(output)-1]
					result, err := applyOperator(left, right, operatorStack[len(operatorStack)-1])
					if err != nil {
						return 0, err
					}
					output = append(output, result)
					operatorStack = operatorStack[:len(operatorStack)-1]
				}
				operatorStack = operatorStack[:len(operatorStack)-1] // pop '('
			} else if op == '!' {
				// Handle factorial, which is a unary operator
				if len(output) > 0 {
					left := output[len(output)-1]
					output = output[:len(output)-1]
					result, err := applyOperator(left, 0, op) // Factorial is applied to left only
					if err != nil {
						return 0, err
					}
					output = append(output, result)
				}
			} else { // Operator
				// Handle the precedence
				for len(operatorStack) > 0 && precedence(operatorStack[len(operatorStack)-1]) >= precedence(op) {
					right := output[len(output)-1]
					output = output[:len(output)-1]
					left := output[len(output)-1]
					output = output[:len(output)-1]
					result, err := applyOperator(left, right, operatorStack[len(operatorStack)-1])
					if err != nil {
						return 0, err
					}
					output = append(output, result)
					operatorStack = operatorStack[:len(operatorStack)-1]
				}
				operatorStack = append(operatorStack, op)
			}
		} else {
			return 0, fmt.Errorf("invalid token: %s", token)
		}
	}

	// Process remaining operators
	for len(operatorStack) > 0 && len(output) > 1 {
		right := output[len(output)-1]
		output = output[:len(output)-1]
		left := output[len(output)-1]
		output = output[:len(output)-1]
		result, err := applyOperator(left, right, operatorStack[len(operatorStack)-1])
		if err != nil {
			return 0, err
		}
		output = append(output, result)
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	if len(operatorStack) > 0 {
		return 0, fmt.Errorf("invalid token")
	}

	if len(output) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return output[0], nil
}
