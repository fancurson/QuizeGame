package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type quiz struct {
	q string
	a string
}

func main() {

	timeLimit := flag.Int("limit", 3, "time limit for quiz hame")

	csvFileName, err := os.Open("problems.csv")
	if err != nil {
		exit(fmt.Sprintf("Open file error: %v", err))
	}
	defer csvFileName.Close()

	reader := csv.NewReader(csvFileName)
	records, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("read error: %v", err))
	}

	data := FillData(records)
	GameTime(data, timeLimit)
	fmt.Println("Thanks for the game!")
}

func GameTime(data []quiz, limit *int) {
	count := 0
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	asw := make(chan string)

	for i, d := range data {
		go func() {
			fmt.Printf("Question #%d: %s = ", i+1, d.q)
			var input string
			fmt.Scanf("%s \n", &input)
			asw <- input
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYour score %d out of %d \n", count, len(data))
			return
		case answer := <-asw:
			if answer == d.a {
				count++
			}
		}
	}
	fmt.Printf("You score %d out of %d \n", count, len(data))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func FillData(records [][]string) []quiz {
	data := make([]quiz, len(records))
	for i, line := range records {
		data[i] = quiz{
			a: strings.TrimSpace(line[1]),
			q: line[0],
		}
	}
	return data
}
