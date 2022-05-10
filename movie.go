package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

type Movie map[string]string

type User struct {
	Movies []int `json:"movies"`
	UserID int   `json:"user_id"`
}

type Movies struct {
	Movie       Movie  `json:"movies"`
	Users       []User `json:"users"`
	Threshold   int
	CountMovies int
}

func NewMovies(threshold int, count int) *Movies {
	return &Movies{
		Threshold:   threshold,
		CountMovies: count + 1,
	}
}

func (m *Movies) LoadJSON(filename string) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Movies) ComputeRecommend(like []int) map[int]string {
	recommended := make(map[int]string)

	a := Sparse(like, m.CountMovies)

	for _, u := range m.Users {
		b := Sparse(u.Movies, m.CountMovies)
		corr := consine_similarity(a, b)
		if corr >= m.Threshold {
			m.AddRecommendation(recommended, u.Movies)
		}
	}

	return recommended
}

func (m *Movies) AddRecommendation(recommended map[int]string, likes []int) {
	for _, like := range likes {
		recommended[like] = m.Movie[strconv.Itoa(like)]
	}
	return
}
