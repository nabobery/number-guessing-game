# number-guessing-game

https://roadmap.sh/projects/number-guessing-game

A simple command-line number guessing game where the computer randomly selects a number, and the user has to guess it within a limited number of attempts.

## Installation

1. Clone the repository:

   ```bash
   https://github.com/nabobery/number-guessing-game.git
   cd number-guessing-game
   ```

2. Build the CLI:

   ```bash
   # For Linux/Mac
   go build -o number-guessing-game

   # For Windows
   go build -o number-guessing-game.exe main.go
   ```

3. Run the executable:

   ```bash
   ./number-guessing-game
   ```

---

## How to Play

1. The game starts with a welcome message and the rules.
2. The computer selects a random number between 1 and 100.
3. You choose a difficulty level (easy, medium, or hard), which determines the number of chances you get.
   - Easy: 10 chances
   - Medium: 5 chances
   - Hard: 3 chances
4. You enter your guess.
5. The game tells you if your guess is too high or too low.
6. The game ends when you guess correctly or run out of chances.

## Sample Output

```bash
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
Please select the difficulty level:
Easy (10 chances)
Medium (5 chances)
Hard (3 chances)
Enter your choice: 2
Great! You have selected the Medium difficulty level.
Let's start the game!
Enter your guess: 50
Incorrect! The number is less than 50.
Enter your guess: 25
Incorrect! The number is greater than 25.
Enter your guess: 35
Incorrect! The number is less than 35.
Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.

```

## Features

- Multiple difficulty levels (easy, medium, hard)
- Feedback on guesses (higher or lower)
- Tracks the number of attempts
- Allow multiple rounds of the game.
- Add a timer to track how long it takes to guess the number.
- Implemented a hint system.
- Tracks high scores for each difficulty level.


## Contributing

Contributions to the Expense Tracker project are welcome! If you find a bug or want to suggest an improvement, please open an issue or submit a pull request.

## License

This project is open-source and free to use under the [MIT License](LICENSE). Contributions are welcome!

```

```
