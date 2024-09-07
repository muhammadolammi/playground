package main

import (
	"fmt"
	"math"
	"strings"
)

var HEXAVALUE = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}

func func1() string {

	return "Hello Word"
}

func func2() func() string {
	var n int

	return func() string {
		n++
		return fmt.Sprintf("Hello word, the %dth time", n)
	}
}

func func3(name string) func() string {

	return func() string {
		return fmt.Sprintf("Hello word, called by %s", name)
	}
}

func func4(name string) func(state string) string {

	return func(state string) string {
		return fmt.Sprintf("Hello word, called by %s from state %s", name, state)
	}
}
func hexaDecimalToDecimal(p *int) int64 {
	if p == nil {
		panic("working with a nil address")
	}

	total := int64(0)

	addressString := fmt.Sprintf("%X", p)                   // Convert pointer to uppercase hex string
	addressString = strings.TrimPrefix(addressString, "0X") // Remove "0x" if present

	for i, char := range reverseString(addressString) { // Process from right to left
		if value, exists := HEXAVALUE[string(char)]; exists {
			total += int64(value * int(math.Pow(16, float64(i))))
		}
	}

	return total
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
