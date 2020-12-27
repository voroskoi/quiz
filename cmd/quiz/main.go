package main

import (
	"flag"
	"fmt"

	"github.com/voroskoi/quiz"
)

func main() {
	quizfile := flag.String("quiz", "problems.csv", "CSV file containing the quiz questions")
	flag.Parse()
	err := quiz.Start(quizfile)
	if err != nil {
		fmt.Println(err)
	}
}
