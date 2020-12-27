package quiz

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readquiz(quizfile *string) ([][]string, error) {
	file, err := os.Open(*quizfile)
	if err != nil {
		return nil, fmt.Errorf("error opening the quiz file: %s", err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read CSV: %s", err)
	}
	return lines, nil
}

// Start starts a quiz form quizfile.
func Start(quizfile *string) error {
	lines, err := readquiz(quizfile)
	if err != nil {
		return err
	}
	var questions, correct int
	for i, line := range lines {
		if len(line) != 2 {
			log.Printf("Every line should consist of a quiestion and an aswer, this one does not: %v", line)
			continue
		}
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
	return nil
}
