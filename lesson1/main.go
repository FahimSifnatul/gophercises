package main

import (
	"flag"
	"log"
)

func main() {
	part := flag.Int("part", 1, "specifies which part of this lesson should run")
	fileName := flag.String("csv", "problems.csv", "csv file name to read")
	t := flag.Int("time", 30, "set time limit in seconds for quiz")
	shuffle := flag.Bool("shuffle", false, "will the question and answer in csv be shuffled or not")
	flag.Parse()

	var err error
	switch *part {
	case 1:
		err = partOneHandler(*fileName, *shuffle)
	case 2:
		// max time limit
		if *t > 3600 {
			*t = 3600
		}
		// for negative value, reset to default
		if *t < 1 {
			*t = 30
		}
		err = partTwoHandler(*fileName, *t, *shuffle)
	default:
		log.Fatalln("invalid part no for lesson 1")
	}

	if err != nil {
		log.Fatalln(err)
	}
}
