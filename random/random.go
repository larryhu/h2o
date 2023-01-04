package random

import (
	"math/rand"
	"time"
)

// Obtain data samples according to the original data and the specified length.
// src original data.
// n the specified length.
// return new data
func Sample[T any](src []T, n int) []T {
	length := len(src)
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	var result []T
	for i := 0; i < n; i++ {
		result = append(result, src[r.Intn(length)])
	}
	return result
}

// Obtain data samples according to the original data.
// src original data.
// return new data
func Choice[T any](src []T) T {
	length := len(src)
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	return src[r.Intn(length)]
}

// Divide a number into multiple numbers, and the sum is equal to the number.
// like  count = s1 + s2 + s3 return [s1,s2,s3] .
// count: target number.
// n: number of splits.
func NumberSplit(count, n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	newNum := count
	var result []int
	for i := 0; i < n-1; i++ {
		a := r.Intn(newNum)
		result = append(result, a)
		newNum -= a
	}
	return append(result, newNum)
}
