package calc

import (
	"strconv"
	"strings"
)

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '(', ')':
		return 0
	default:
		return -1
	}
}

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
			return 0, divisionByZero
		}
		return left / right, nil
	default:
		return 0, unsupportedOperation
	}
}

func tokenize(expression string) ([]string, error) {
	var tokens []string
	var current string
	for i, r := range expression {
		if r == ' ' {
			continue
		}
		if r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			if r == '-' && (i == 0 || strings.ContainsRune("+-*/(", rune(expression[i-1]))) {
				current += string(r)
			} else {
				tokens = append(tokens, string(r))
			}
		} else if r >= '0' && r <= '9' || r == '.' {
			current += string(r)
		} else {
			return nil, invalidCharacter
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

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			output = append(output, num)
		} else if len(token) == 1 && strings.ContainsRune("+-*/()", rune(token[0])) {
			op := rune(token[0])
			if op == '(' {
				operatorStack = append(operatorStack, op)
			} else if op == ')' {
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
				operatorStack = operatorStack[:len(operatorStack)-1]
			} else {
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
			return 0, invalidToken
		}
	}

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
		return 0, invalidToken
	}

	if len(output) != 1 {
		return 0, invalidExpression
	}

	return output[0], nil
}
