package main

import (
	"testing"
)

func TestMovie_LoadJSON_0(t *testing.T) {

	var movies Movies
	err := movies.LoadJSON("adsa1334dZQmovies.json")
	if err == nil {
		t.Fatalf("should not be nil")
	}
}

func TestMovie_LoadJSON_1(t *testing.T) {

	var movies Movies
	err := movies.LoadJSON("movies.json")
	if err != nil {
		t.Fatalf("should be nil")
	}
}
