package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"time"
)

type Record struct {
	Time time.Duration
	Name string
}

func loadRecords() []Record {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	var records []Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		fmt.Println(err)
	}

	return records
}

func saveRecords(records []Record) {
	data, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("NEW RECORD!")
}

func setRecords(time time.Duration) {
	var records []Record = loadRecords()

	for i := 0; i < len(records) || (i < 3); i++ {
		if time < records[i].Time {
			var record Record
			fmt.Println("Enter your name: ")
			fmt.Scan(&record.Name)
			record.Time = time
			records = append(records[:i], append([]Record{record}, records[i:]...)...)
			records = records[:3]
			saveRecords(records)
			break
		}
	}
	fmt.Println("\n---Top 3 Records---")
	for i, record := range records {
		fmt.Println(i+1, ". ", record.Name, " - ", record.Time)
	}
}

func giveHint(number int, guess int) {
	if guess > number {
		fmt.Println("Try a lower number")
	} else {
		fmt.Println("Try a higher number")
	}
}

func guessNumber(number int, trys float64) bool {
	var guess int
	fmt.Println("\n------You have ", trys, " trys left------")
	fmt.Println("Enter your guess: ")
	fmt.Scan(&guess)

	if guess == number {
		fmt.Println("You guessed it!")
		return true
	} else {
		fmt.Println("\n----- Better luck next time! -----")
		giveHint(number, guess)
		trys -= 1
	}

	if trys == 0 {
		fmt.Print("\noh oh! You have no more trys left")
		return false
	}
	return guessNumber(number, trys)
}

func main() {
	fmt.Println("\nWelcome to the Number Guessing Game! \nI'm thinking of a number between 1 and 100.")
	number := rand.IntN(100)
	fmt.Println("---Please Select Difficulty---")
	fmt.Println("1. Easy (10 trys)")
	fmt.Println("2. Medium (5 trys)")
	fmt.Println("3. Hard (3 trys)")
	fmt.Println("4. Infinit (Infinite trys)")
	var difficulty int
	fmt.Scan(&difficulty)
	startTime := time.Now()

	var guess bool = false

	switch difficulty {
	case 1:
		guess = guessNumber(number, 10)
	case 2:
		guess = guessNumber(number, 5)
	case 3:
		guess = guessNumber(number, 3)
	case 4:
		guess = guessNumber(number, math.Inf(1))
	default:
		fmt.Println("Invalid input")
	}
	if guess {
		endTime := time.Now()
		finalTime := endTime.Sub(startTime)
		fmt.Println("Time taken: ", finalTime)
		setRecords(finalTime)
	}
}
