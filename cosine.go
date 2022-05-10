package main

import (
	"math"
)

func consine_similarity(a []int, b []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum = sum + (a[i] * b[i])
	}

	a_dist := 0
	for i := 0; i < len(a); i++ {
		a_dist = a_dist + a[i]*a[i]
	}

	b_dist := 0
	for i := 0; i < len(a); i++ {
		b_dist = b_dist + b[i]*b[i]
	}

	if a_dist == 0 || b_dist == 0 {
		return 0
	}

	return int(float64(100*sum)/(math.Sqrt(float64(a_dist))*math.Sqrt(float64(b_dist))) + 0.5)
}
