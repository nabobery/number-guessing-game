package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Difficulty levels
type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

func (d Difficulty) String() string {
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

func (d Difficulty) Attempts() int {
	return [...]int{10, 5, 3}[d]
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	difficulty := chooseDifficulty()
	fmt.Printf("Great! You have selected the %s difficulty level.\n", difficulty)

	playGame(difficulty)
}

func chooseDifficulty() Difficulty {
	fmt.Println("Please select the difficulty level:")
	fmt.Println("Easy (10 chances)")
	fmt.Println("Medium (5 chances)")
	fmt.Println("Hard (3 chances)")

	var choice int
	for {
		fmt.Print("Enter your choice (0 for Easy, 1 for Medium, 2 for Hard): ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 0 || choice > 2 {
			fmt.Println("Invalid input. Please enter 0, 1, or 2.")
			continue
		}
		break
	}
	return Difficulty(choice)
}

func playGame(difficulty Difficulty) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	secretNumber := random.Intn(100) + 1
	attempts := difficulty.Attempts()
	fmt.Println("Let's start the game!")

	for attempt := 1; attempt <= attempts; attempt++ {
		fmt.Printf("Attempt %d/%d. Enter your guess: ", attempt, attempts)
		var guess int
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			attempt-- // Decrement attempt to allow retry
			continue
		}

		if guess == secretNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempt)
			return
		} else if guess < secretNumber {
			fmt.Printf("Incorrect! The number is greater than your guess: %d.\n", guess)
		} else {
			fmt.Printf("Incorrect! The number is less than your guess: %d.\n", guess)
		}
	}

	fmt.Printf("You ran out of attempts. The correct number was %d.\n", secretNumber)
}
