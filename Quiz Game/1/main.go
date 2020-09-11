package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("file", "problems.csv", "CSV file name to a problem")
	flag.Parse()

	filePtr, err := os.Open(*fileName)

	if err != nil {
		fmt.Println("Couldn't open ", fileName, " file", err)
	}

	fmt.Println("Press any key to start the quiz..")
	var starter int
	fmt.Scanf("%d", &starter)

	problems := csv.NewReader(filePtr)

	result, nProblems := 0, 0

	for {
		line, err := problems.Read()

		if err == nil {
			nProblems++
			p := problem{
				question: line[0],
				answer:   line[1],
			}

			userAnswer := p.askQuestion()
			if p.checkAnswer(userAnswer) {
				result++
			}

		} else if err == io.EOF {
			break
		} else {
			fmt.Println("An error occured while reading the file.", err)
			os.Exit(1)
		}

	}

	fmt.Println("You got", result, "correct out of", nProblems, ".")
}

func (p problem) askQuestion() string {
	fmt.Printf("Question: %s\n", p.question)
	var userAnswer string
	fmt.Scanln(&userAnswer)

	return userAnswer
}

func (p problem) checkAnswer(answer string) bool {
	return p.answer == answer
}
