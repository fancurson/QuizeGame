package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
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
	GameTime(data)

	fmt.Println("Thanks for game!")
}

type quiz struct {
	q string
	a string
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

func GameTime(data []quiz) {
	count := 0
	for i, d := range data {
		fmt.Printf("Question #%d: %s = ", i+1, d.q)
		var input string
		fmt.Scanf("%s \n", &input)
		if input == d.a {
			count++
		}
	}
	fmt.Printf("You score %d out of %d \n", count, len(data))
}
