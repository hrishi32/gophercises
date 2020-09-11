package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("file", "problems.csv", "CSV file name to a problem")
	timeLimit := flag.Int("limit", 30, "Time in seconds for quiz timeout")
	flag.Parse()

	fmt.Println("Press any key to start the quiz..")
	var starter int
	fmt.Scanf("%d", &starter)

	filePtr, err := os.Open(*fileName)

	if err != nil {
		fmt.Println("Couldn't open ", fileName, " file", err)
	}

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
			ansChannel := make(chan bool)

			go p.askQuestion(ansChannel)

			timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

			select {
			case r := <-ansChannel:
				timer.Stop()
				if r {
					result++
					fmt.Println("Correct Ans")
				} else {
					fmt.Println("Wrong Ans")
				}

			case <-timer.C:
				fmt.Println("You got", result, "correct out of", nProblems, ".")
				return
			}

		} else if err == io.EOF {
			fmt.Println("You got", result, "correct out of", nProblems, ".")
			break
		} else {
			fmt.Println("An error occured while reading the file.", err)
			os.Exit(1)
		}

	}

}

func (p problem) askQuestion(c chan bool) {

	fmt.Printf("Question: %s\n", p.question)
	var userAnswer string

	fmt.Scanln(&userAnswer)

	c <- p.checkAnswer(userAnswer)
}

func (p problem) checkAnswer(answer string) bool {
	return p.answer == answer
}
