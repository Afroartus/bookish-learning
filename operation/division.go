package operation

import "errors"

// Div // Division de dos enteros
func Div(dividend, divisor uint64) (int64, error) {
	if divisor != 0 {
		return int64(dividend / divisor), nil
	}
	return 0, errors.New("invalid operation")
}
