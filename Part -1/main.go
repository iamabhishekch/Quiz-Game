package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	// taking file from the user and parsing that
	csvFile := flag.String("csv", "problems.csv", "your problem csv file for Quize Game")
	flag.Parse()

	// opening the file from the user - passing pointer because flag.String returns *String
	file, err := os.Open(*csvFile)
	exitOnError("Failed to Open csv file",err )

	// executing csv file reader which says like - hey I want to read it as a CSV — so treat commas as field separators, handle quotes properly, and give me rows and columns as Go strings."
	r := csv.NewReader(file)

	// now it returns *csv.Reader which has these methods
	// Read() — reads one row at a time ([]string)

	// ReadAll() — reads all rows at once ([][]string)
	line, err := r.ReadAll()
	exitOnError("Failed to read the csv file", err)
	
	// calling parseLine to parse the line that we get from RealAll() and save Problem struct to problems 
	problems:= parseLines(line)

	var totalCorrect int
	//ranging over a []struct to get que only 
	for i, v := range problems{
		fmt.Printf("Problem %d: %s\n", i +1, v.q)

		// taking user input and matching that input to answer 
		var answer string 
		fmt.Scanf("%s\n", &answer)
		if answer == v.a{
			totalCorrect ++
			fmt.Println("Correct Answer!")
		}else{
			fmt.Println("Wrong Answer :)")
		}

	}

	fmt.Printf("You have scored %d Out of %d.", totalCorrect, len(problems))

	
	

}

// struct to parse the que. and ans.

type problem struct {
	q string
	a string
}

// handling error

func exitOnError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
		os.Exit(1)
	}

}

// ParseLines func to parser the line so we can get the que and ans OR []problem
func parseLines(lines [][]string,)[]problem{
	ret := make([]problem, len(lines))
	for i, line:= range lines{
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}
