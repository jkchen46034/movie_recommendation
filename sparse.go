package main

func Sparse(v []int, capacity int) []int {
	vector := make([]int, capacity)
	for _, e := range v {
		vector[e] = 1
	}
	return vector
}
