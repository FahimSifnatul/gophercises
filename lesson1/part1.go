package main

import "fmt"

func partOneHandler(fileName string) error {
	q, err := getQuizFromCsv(fileName)
	if err != nil {
		return err
	}

	//score := make(map[string]int)
	var zero int
	var score *int
	score = &zero
	startQuiz(q, score)
	fmt.Printf("You scored %d out of %d\n", *score, len(q))
	return nil
}
