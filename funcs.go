package main

import "fmt"

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
