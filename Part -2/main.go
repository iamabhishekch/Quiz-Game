package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFile := flag.String("csv", "problems.csv", "your problem csv file for Quize Game")
	flag.Parse()

	file, err := os.Open(*csvFile)
	exitOnError("Failed to Open csv file", err)

	r := csv.NewReader(file)

	line, err := r.ReadAll()
	exitOnError("Failed to read the csv file", err)

	problems := parseLines(line)

	var totalCorrect int
	for i, v := range problems {
		fmt.Printf("Problem %d: %s\n", i+1, v.q)

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == v.a {
			totalCorrect++
			fmt.Println("Correct Answer!")
		} else {
			fmt.Println("Wrong Answer :)")
		}

	}

	fmt.Printf("You have scored %d Out of %d.", totalCorrect, len(problems))

}

type problem struct {
	q string
	a string
}

func exitOnError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
		os.Exit(1)
	}

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}
