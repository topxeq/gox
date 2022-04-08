package main

import (
	"fmt"
	"time"
)

func fibonacci(c int64) int64 {
	if c < 2 {
		return c
	}

	return fibonacci(c-2) + fibonacci(c-1)
}

func main() {
	startTime := time.Now()

	resultInt := fibonacci(50)

	endTime := time.Now()

	fmt.Printf("Result: %v\n", resultInt)

	fmt.Printf("Duration: %v\n", endTime.Sub(startTime))

}
