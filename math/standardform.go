package math

import (
	"errors"
	"fmt"
	"math"
)

func StandardForm(n float64) string {
	if n == 0 {
		return "0"
	}
	exp := int(math.Floor(math.Log10(math.Abs(n))))
	coeff := n / math.Pow(10, float64(exp))

	return fmt.Sprintf("%.f * 10^%d", coeff, exp)
}

func NeedInt(n interface{}) (int, error) {
	value, ok := n.(int)
	if !ok {
		return 0, errors.New("not an int")

	}
	return value + 1, nil
}
