package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 98; i++ {
		result := Fibonacci(i)
		fmt.Printf("%d: %d\r\n", i, result)
	}
}

func Fibonacci(n int) uint64 {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
