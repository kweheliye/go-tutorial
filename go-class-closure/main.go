package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}
}
