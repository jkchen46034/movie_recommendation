package main

import (
	"testing"
)

func TestCosineSimilarity_0(t *testing.T) {
	a := []int{0, 0, 0}
	b := []int{0, 0, 0}
	corr := consine_similarity(a, b)
	if corr != 0 {
		t.Fatalf("should be 0")
	}
}

func TestCosineSimilarity_1(t *testing.T) {
	a := []int{1, 0, 1, 0, 1}
	b := []int{0, 1, 0, 1, 0}
	corr := consine_similarity(a, b)
	if corr != 0 {
		t.Fatalf("should be 0")
	}
}

func TestCosineSimilarity_2(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	b := []int{1, 1, 1, 1, 1}
	corr := consine_similarity(a, b)
	if corr != 100 {
		t.Fatalf("should be 100")
	}
}

func TestCosineSimilarity_3(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	b := []int{1, 1, 1, 1, 0}
	corr := consine_similarity(a, b)
	if corr != 89 {
		t.Fatalf("should be 89")
	}
}

func TestCosineSimilarity_4(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	b := []int{1, 1, 1, 0, 0}
	corr := consine_similarity(a, b)
	if corr != 77 {
		t.Fatalf("should be 77")
	}
}

func TestCosineSimilarity_5(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	b := []int{1, 1, 0, 0, 0}
	corr := consine_similarity(a, b)
	if corr != 63 {
		t.Fatalf("should be 63")
	}
}

func TestCosineSimilarity_6(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	b := []int{1, 0, 0, 0, 0}
	corr := consine_similarity(a, b)
	if corr != 45 {
		t.Fatalf("should be 45")
	}
}

func TestCosineSimilarity_7(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	b := []int{0, 0, 0, 0, 0}
	corr := consine_similarity(a, b)
	if corr != 0 {
		t.Fatalf("should be 0")
	}
}
