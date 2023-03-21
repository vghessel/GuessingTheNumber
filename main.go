package main

import (
	"fmt"
	"math/rand"
)

var attempts int = 0

func main() {
	randomNumber := randomGenerate()

	displayIntro()

	Dificulty()

	command := readCommand()

	attempts := commandCases(command)

	startGame(attempts, randomNumber)

}

func displayIntro() {
	fmt.Println("          Welcome to    ")
	fmt.Println("   ---- GUESSING GAME ----")
	fmt.Println("Find the number from 0 to 100")
	fmt.Println("")
}

func Dificulty() {
	fmt.Println("          Dificulty      ")
	fmt.Println("(1) Easy  (2) Medium  (3) Hard")

}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("Selected Dificulty:", command)
	fmt.Print("")

	return command
}

func startGame(attempts int, randomNumber int) {

	for i := 1; i <= attempts; i++ {
		fmt.Println("Attempt", i, "of", attempts)
		numberGuessed := guessANumber()
		if numberGuessed == randomNumber {
			fmt.Println("CONGRATULATIONS, YOU ROCKS!")
			break
		} else {
			if i == attempts {
				fmt.Println("You lose!")
				fmt.Println("The number was:", randomNumber)
				break
			}
			if numberGuessed > randomNumber {
				fmt.Println("a little lower...")
				fmt.Println("")
			}
			if numberGuessed < randomNumber {
				fmt.Println("a little higher...")
				fmt.Println("")
			}
		}

	}
}

func guessANumber() int {
	var numberGuessed int
	fmt.Println("Guess The Number:")
	fmt.Scan(&numberGuessed)
	return numberGuessed
}

func randomGenerate() int {
	randomNumber := rand.Intn(100)
	return randomNumber
}

func commandCases(command int) int {
	if command == 1 {
		attempts = 12
		fmt.Println("Easy Peasy :)")
	}
	if command == 2 {
		attempts = 8
		fmt.Println("Respectable :)")
	}
	if command == 3 {
		attempts = 5
		fmt.Println("So full of himself °-°")
	}

	return attempts
}
