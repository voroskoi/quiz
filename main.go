package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	quizfile := flag.String("quiz", "problems.csv", "CSV file containing the quiz questions")
	flag.Parse()
	file, err := os.Open(*quizfile)
	if err != nil {
		log.Printf("Unable to open %s (%s)", *quizfile, err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Printf("Unable to read CVS: %s", err)
	}
	var questions, correct int
	for i, line := range lines {
		fmt.Printf("Question (%d): %s\n", i, line[0])
		var ans string
		_, err = fmt.Scanln(&ans)
		if err != nil {
			log.Printf("Error reading the answer: %s", err)
		}
		if ans == line[1] {
			correct++
		}
		questions++
	}
	fmt.Printf("Questions asked: %d, right answer: %d\n", questions, correct)
}
