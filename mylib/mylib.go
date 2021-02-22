package mylib

import (
	"math"
)

// min is return maximum value
func min(nums ...int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

// max is return maximum value
func max(nums ...int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}

// gcd is Euclidean Algorithm
func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}
