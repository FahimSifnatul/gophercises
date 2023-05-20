package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	partNo := *flag.Int("part", 1, "specifies which part of this lesson should run")
	fileName := *flag.String("csv", "problems.csv", "csv file name to read")

	var total, score int
	var err error
	switch partNo {
	case 1:
		total, score, err = partOneHandler(fileName)
	default:
		log.Fatalln("invalid part no for lesson 1")
	}

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("You scored %d out of %d\n", score, total)
}
