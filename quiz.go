package quiz

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func readquiz(quizfile *string) ([][]string, error) {
	file, err := os.Open(*quizfile)
	defer file.Close()
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

	timer := time.NewTimer(3 * time.Second)

	var questions, correct int
	for i, line := range lines {
		if len(line) != 2 {
			log.Printf("Every line should consist of a quiestion and an aswer, this one does not: %v", line)
			continue
		}
		fmt.Printf("Question (%d): %s\n", i, line[0])
		questions++

		ansCh := make(chan string)
		go func() {
			var ans string
			_, err = fmt.Scanf("%s\n", &ans)
			if err != nil {
				log.Printf("Error reading the answer: %s", err)
			}
			ansCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time is up! Questions asked: %d, right answer: %d\n", questions, correct)
			return nil
		case ans := <-ansCh:
			if ans == line[1] {
				correct++
			}
		}
	}
	fmt.Printf("Questions asked: %d, right answer: %d\n", questions, correct)
	return nil
}
