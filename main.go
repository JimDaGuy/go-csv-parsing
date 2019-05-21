package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	firstName string
	lastName  string
	age       string
	state     string
}

func main() {
	// Open CSV file
	csvFile, err := os.Open("assignment-data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Parse csv file
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// Slice to hold all the people in it
	var people []Person

	// Read through lines of the file
	for {
		// Read next line
		line, err2 := reader.Read()

		// End of file
		if err2 == io.EOF {
			break
		} else if err2 != nil {
			// Fatal error
			log.Fatal(err2)
		}

		// Store line in Person struct
		person := Person{
			firstName: line[0],
			lastName:  line[1],
			age:       line[2],
			state:     line[3],
		}
		// Print out the person struct
		fmt.Println(person)
		// Append the person to a slice of people depending on their age
		people = append(people, person)
	}
	fmt.Println("Hello World!")
}
