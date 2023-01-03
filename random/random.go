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
