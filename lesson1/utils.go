package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func getQuizFromCsv(name string) ([]Quiz, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("file %s is not found or can't be opened", name)
	}
	defer file.Close()

	results := make([]Quiz, 0)
	r := csv.NewReader(file)
	for line := 1; ; line++ {
		data, err := r.Read()
		if err != nil {
			if err == io.EOF {
				return results, nil
			}
			return nil, err
		}

		if len(data) == 0 {
			continue
		}
		if len(data) != 2 {
			return nil, fmt.Errorf("invalid data in csv file: line no - %d: more than two values", line)
		}

		res := Quiz{Ques: data[0]}
		res.Ans, err = strconv.Atoi(data[1])
		if err != nil {
			return nil, fmt.Errorf("invalid data in csv file: line no - %d: answer can't be convereted to integer", line)
		}
		results = append(results, res)
	}
}

func getQuizResult(qz []Quiz) int {
	var score, ans int

	for i, q := range qz {
		fmt.Printf("Question #%d: %s = ", i+1, q.Ques)

		_, err := fmt.Scanf("%d\n", &ans)
		if err == nil && ans == q.Ans {
			score++
		}
	}
	return score
}
