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
	for {
		fmt.Println("I'm thinking of a number between 1 and 100.")
		difficulty := chooseDifficulty()
		fmt.Printf("Great! You have selected the %s difficulty level.\n", difficulty)

		playGame(difficulty)

		if !playAgain() {
			break
		}
	}
	fmt.Println("Thanks for playing!")
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
	startTime := time.Now()
	var guess int
	var attempt int
	for attempt = 1; attempt <= attempts; {
		fmt.Printf("Attempt %d/%d. Enter your guess: ", attempt, attempts)
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if guess == secretNumber {
			duration := time.Since(startTime)
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts and it took you %v.\n", attempt, duration)
			updateHighScore(difficulty, attempt, duration)
			return
		} else if guess < secretNumber {
			fmt.Printf("Incorrect! The number is greater than your guess: %d.\n", guess)
		} else {
			fmt.Printf("Incorrect! The number is less than your guess: %d.\n", guess)
		}
		attempt++
		if attempt <= attempts && attempt%2 == 0 {
			hint := provideHint(secretNumber, guess)
			fmt.Printf("Hint: %s\n", hint)
		}
	}

	fmt.Printf("You ran out of attempts. The correct number was %d.\n", secretNumber)
}

func playAgain() bool {
	var choice string
	for {
		fmt.Print("Do you want to play again? (yes/no): ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter yes or no.")
			continue
		}
		if choice == "yes" {
			return true
		} else if choice == "no" {
			return false
		} else {
			fmt.Println("Invalid input. Please enter yes or no.")
		}
	}
}

func provideHint(secretNumber, guess int) string {
	if secretNumber > guess {
		return "The number is higher than your guess."
	} else {
		return "The number is lower than your guess."
	}
}

var highScores = make(map[Difficulty]struct {
	attempts int
	duration time.Duration
})

func updateHighScore(difficulty Difficulty, attempts int, duration time.Duration) {
	currentScore, ok := highScores[difficulty]
	if !ok || attempts < currentScore.attempts || (attempts == currentScore.attempts && duration < currentScore.duration) {
		highScores[difficulty] = struct {
			attempts int
			duration time.Duration
		}{attempts, duration}
		fmt.Println("New high score!")
		printHighScores()
	}
}

func printHighScores() {
	fmt.Println("High Scores:")
	for difficulty, score := range highScores {
		fmt.Printf("%s: %d attempts in %v\n", difficulty, score.attempts, score.duration)
	}
}
