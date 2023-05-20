package main

import (
	"fmt"
)

func partOneHandler(fileName string, shuffle bool) error {
	q, err := getQuizFromCsv(fileName)
	if err != nil {
		return err
	}

	if shuffle {
		shuffleQuiz(q)
	}

	var zero int
	var score *int
	score = &zero
	startQuiz(q, score)
	fmt.Printf("You scored %d out of %d\n", *score, len(q))
	return nil
}
