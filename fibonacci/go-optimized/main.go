package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 50; i++ {
		result := Fibonacci(i)

		epoch := time.Now().Unix()

		fmt.Printf("%d: %d (%d)\r\n", i, result, epoch)
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
