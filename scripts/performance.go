package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/topxeq/tk"
)

func calPi(pointCount int) float64 {
	inCircleCount := 0

	var x, y float64
	var Pi float64

	r := tk.NewRandomGenerator()

	for i := 0; i < pointCount; i++ {
		// x = rand.Float64()
		// y = rand.Float64()
		x = r.Float64()
		y = r.Float64()

		if x*x+y*y < 1 {
			inCircleCount++
		}
	}

	Pi = (4.0 * float64(inCircleCount)) / float64(pointCount)

	return Pi
}

func fibonacci(c int64) int64 {
	if c < 2 {
		return c
	}

	return fibonacci(c-2) + fibonacci(c-1)
}

func fibonacciFlat(c int64) int64 {
	if c < 2 {
		return c
	}

	var fibo int64 = 1
	var fiboPrev int64 = 1
	for i := int64(2); i < c; i++ {
		temp := fibo
		fibo += fiboPrev
		fiboPrev = temp
	}

	return fibo
}

func main() {
	rand.Seed(time.Now().Unix()) // 初始化随机数

	// fmt.Printf("Test 1\n")

	// startTime := time.Now()

	// result := 0.0

	// for i := 0.0; i < 1000000000; i = i + 1 {
	// 	result += i * i
	// }

	// endTime := time.Now()

	// fmt.Printf("Result: %v\n", result)

	// fmt.Printf("Duration: %v\n", endTime.Sub(startTime))

	countT := 10000000

	fmt.Printf("\nGolang Test 2\n")

	startTime := time.Now()

	result := calPi(countT)

	endTime := time.Now()

	fmt.Printf("Result: %v\n", result)

	fmt.Printf("Duration: %v s\n", endTime.Sub(startTime))

	return

	fmt.Printf("Golang Test 2\n")

	startTime = time.Now()

	resultInt := fibonacciFlat(10000000000)

	endTime = time.Now()

	fmt.Printf("Result: %v\n", resultInt)

	fmt.Printf("Duration: %v\n", endTime.Sub(startTime))

	fmt.Printf("Test 2r\n")

	startTime = time.Now()

	resultInt = fibonacci(50)

	endTime = time.Now()

	fmt.Printf("Result: %v\n", resultInt)

	fmt.Printf("Duration: %v\n", endTime.Sub(startTime))

	fmt.Printf("Test 3\n")

	startTime = time.Now()

	result = 0.0

	for i := 0.0; i < 100000000; i = i + 1 {
		result += rand.Float64()
	}

	endTime = time.Now()

	fmt.Printf("Result: %v\n", result)

	fmt.Printf("Duration: %v\n s", endTime.Sub(startTime))

}
