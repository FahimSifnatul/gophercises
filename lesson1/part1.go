package main

func partOneHandler(fileName string) (int, int, error) {
	q, err := getQuizFromCsv(fileName)
	if err != nil {
		return 0, 0, err
	}

	total := len(q)
	score := getQuizResult(q)
	return total, score, nil
}
