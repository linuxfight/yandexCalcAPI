package calc

import "errors"

var (
	divisionByZero       = errors.New("division by zero")
	unsupportedOperation = errors.New("unsupported operation")
	invalidCharacter     = errors.New("invalid character in expression")
	invalidToken         = errors.New("invalid token")
	invalidExpression    = errors.New("invalid expression")
)
