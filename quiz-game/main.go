package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "A csv file which has the questions and answers in the format of question,answer")
	timeLimit := flag.Int("limit", 10, "Time limit for the question to answer")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Unable to open the file: %s\n", *csvFile))
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		exit("Unable to read the contents of the file")
	}
	// Reading all the lines in the CSV file
	problems := parseCsvFile(records)

	// Adding a timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var score int // Score
	for idx, prblm := range problems {
		fmt.Printf("Problem #%d: %v = ", idx+1, prblm.question)

		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Your Score: ", score)
			return
		case ans := <-answerChan:
			if ans == prblm.answer {
				fmt.Println("Correct!!")
				score++
			} else {
				fmt.Println("\nWrong!!")
			}
		}
	}
}

func parseCsvFile(records [][]string) []problem {
	problems := make([]problem, len(records))
	for idx, value := range records {
		problems[idx] = problem{
			question: value[0],
			answer:   strings.TrimSpace(value[1]),
		}
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
