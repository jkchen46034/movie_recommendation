package main

import (
	"testing"
)

func TestCosineSparse_0(t *testing.T) {

	a := []int{2, 3, 8}
	s := Sparse(a, 80)

	if s[2] != 1 {
		t.Fatalf("should be 1")
	}
	if s[3] != 1 {
		t.Fatalf("should be 1")
	}
	if s[8] != 1 {
		t.Fatalf("should be 1")
	}
	if s[0] != 0 {
		t.Fatalf("should be 0")
	}
	if s[79] != 0 {
		t.Fatalf("should be 0")
	}

}
