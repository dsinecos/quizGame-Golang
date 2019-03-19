package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type quizQuestion struct {
	question string
	answer   string
}

type quizScore struct {
	correct   int
	incorrect int
}

var csvFilename *string

func init() {
	csvFilename = flag.String("filename", "", "Name of csv file storing quiz data")
	flag.Parse()
}

func main() {
	fmt.Println("CSV Filename :", *csvFilename)

	if *csvFilename == "" {
		panic("No CSV filename for quiz data specified in flag")
	}

	fmt.Println("\nQuiz Game")
	csvFile, error := os.Open(*csvFilename)

	if error != nil {
		panic(error)
	}

	reader := csv.NewReader(csvFile)

	var quizPaper []quizQuestion
	var userScore quizScore

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		quizPaper = append(quizPaper, quizQuestion{
			question: line[0],
			answer:   line[1],
		})
	}

	for _, quizItem := range quizPaper {
		fmt.Println("Question : ", quizItem.question)
		var userAnswer string
		fmt.Scan(&userAnswer)

		if userAnswer == quizItem.answer {
			fmt.Println("Correct")
			userScore.correct = userScore.correct + 1
		} else {
			fmt.Println("Incorrect")
			userScore.incorrect = userScore.incorrect + 1
		}
	}

	fmt.Println("Final score")
	fmt.Println("Total questions : " + strconv.Itoa(len(quizPaper)))
	fmt.Println("You got " + strconv.Itoa(userScore.correct) + " correct.")
}
