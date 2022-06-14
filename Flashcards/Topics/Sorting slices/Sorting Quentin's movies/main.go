package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Movie struct {
	Name   string
	Rating int
}

func main() {
	movies := readMovies()

	// Implement the sorting logic within the sort.Slice function below!
	sort.Slice(movies, func(i, j int) bool {
		if movies[i].Rating != movies[j].Rating {
			return movies[i].Rating > movies[j].Rating
		}
		return movies[i].Name > movies[j].Name
	})

	// The print statement below outputs the sorted slice, do not delete it!
	fmt.Println(movies)
}

// readMovies is a helper to read the movies for this task
// DO NOT EDIT
func readMovies() []Movie {
	var (
		name   string
		rating int
		movies []Movie
	)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		name = scanner.Text()

		scanner.Scan()
		fmt.Sscan(scanner.Text(), &rating)

		movies = append(movies, Movie{name, rating})
	}
	return movies
}
