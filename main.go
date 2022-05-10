package main

import (
	"fmt"
	"log"
)

func main() {
	log.Printf("Hello Selfbook!")

	threshold := 40
	count_movies := 70
	movies := NewMovies(threshold, count_movies)

	err := movies.LoadJSON("movies.json")
	if err != nil {
		panic(err)
	}
	likes := []int{5, 9, 13, 26, 34, 37, 53}
	recommended := movies.ComputeRecommend(likes)
	fmt.Println(recommended)
}
