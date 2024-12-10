package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	FirstName   	string    	`json:"firstName"`
	Password 		string 		`json:"password"`
	PasswordHash  	string    	`json:"password_hash"`
}

func main() {
	filePath := "data.json"
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var people []Person
	err = json.Unmarshal(fileData, &people)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	fmt.Println("Parsed Data:")
	for _, person := range people {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", person.ID, person.Name, person.Age)
	}

	fmt.Println("\nReading file line-by-line:")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading lines: %v\n", err)
	}
}
