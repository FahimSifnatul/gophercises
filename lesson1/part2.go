package main

import (
	"fmt"
	"time"
)

func partTwoHandler(fileName string, t int, shuffle bool) error {
	q, err := getQuizFromCsv(fileName)
	if err != nil {
		return err
	}

	if shuffle {
		shuffleQuiz(q)
	}

	done := make(chan bool)
	defer close(done)

	var zero int
	var score *int
	score = &zero
	go startQuiz2(q, score, done)

	select {
	case <-done:
		fmt.Printf("You scored %d out of %d\n", *score, len(q))
		return nil
	case <-time.After(time.Duration(t) * time.Second):
		fmt.Printf("\nYou scored %d out of %d\n", *score, len(q))
		return nil
	}
}
