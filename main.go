package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
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

	// Grouping the records into two groups as stated in the assignment
	// Splitting the groups into states using a map
	var peopleOver30 = make(map[string][]Person)
	var peopleUnder30 = make(map[string][]Person)

	headerline := true
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

		// Skip headerline
		if headerline {
			headerline = false
			continue
		}

		// Convert age string
		personAge, err3 := strconv.Atoi(line[2])
		if err3 != nil {
			fmt.Printf("Error converting \"%s\" to integer, setting to 0\n", line[2])
		}

		// Store line in Person struct
		person := Person{
			firstName: line[0],
			lastName:  line[1],
			age:       personAge,
			state:     line[3],
		}

		currState := line[3]

		// Append the person to the correct map and slice
		if personAge >= 30 {
			// Check if the map has set a value for the current state
			_, stExist := peopleOver30[currState]

			// If the map has already set a value for the current state, append the map
			// Otherwise make a new slice with the current person
			if stExist {
				peopleOver30[currState] = append(peopleOver30[currState], person)
			} else {
				peopleOver30[currState] = []Person{person}
			}
		} else {
			_, stExist := peopleUnder30[currState]
			// If the map has already set a value for the current state, append the map
			// Otherwise make a new slice with the current person
			if stExist {
				peopleUnder30[currState] = append(peopleUnder30[currState], person)
			} else {
				peopleUnder30[currState] = []Person{person}
			}
		}
	}

	// Output

	fmt.Println("30+ Years Old")
	fmt.Println("--------------------")

	// Over 30 Map
	for state, people := range peopleOver30 {
		fmt.Printf("State: %s\n", state)
		for _, person := range people {
			fmt.Printf("Document: %s %s, %s, %v\n", person.firstName, person.lastName, person.state, person.age)
		}
		fmt.Println("")
	}

	fmt.Println("<30 Years Old")
	fmt.Println("--------------------")

	// Under 30 Map
	for state, people := range peopleUnder30 {
		fmt.Printf("State: %s\n", state)
		for _, person := range people {
			fmt.Printf("Document: %s %s, %s, %v\n", person.firstName, person.lastName, person.state, person.age)
		}
		fmt.Println("")
	}
}
