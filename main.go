package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Guess the number between 1 and 100!")

	secretNumber := rand.Intn(100) + 1
	numGuesses := 0

	for {
		var guess int
		fmt.Print("Your guess: ")
		fmt.Scanln(&guess)
		numGuesses++

		if guess == secretNumber {
			fmt.Printf("Congratulations, you guessed the number in %d guesses!\n", numGuesses)
			break
		} else if guess < secretNumber {
			fmt.Println("Too low, try again.")
		} else {
			fmt.Println("Too high, try again.")
		}
	}
}
