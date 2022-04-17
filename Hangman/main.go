package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func rune_in_slice(c rune, arr []rune) bool {
	for _, v := range arr {
		if v == c {
			return true
		}
	}

	return false
}

func char_in_string(c rune, s string) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}

	return false
}

func main() {
	words := []string{"apple", "banana", "orange", "lemon"}

	rand.Seed(time.Now().UTC().UnixNano())
	word := words[rand.Intn(len(words))]

	guessed_letters := make([]rune, 0)

	guess := 6
	exit := false

	fmt.Println(word)
	for guess > 0 && !exit {
		fmt.Println(guess, "guesses left")

		for _, c := range word {
			if rune_in_slice(c, guessed_letters) {
				fmt.Print(string(c))
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()

		fmt.Println("Enter a letter:")
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// check for empty inputs
		if input == "" {
			continue
		}
		char := []rune(input)[0]

		if rune_in_slice(char, guessed_letters) || !unicode.IsLetter(char) {
			fmt.Println()
			if !unicode.IsLetter(char) {
				fmt.Println("Only enter lowercase letters!")
			} else {
				fmt.Println("You've already guessed that")
			}
			continue
		}

		guessed_letters = append(guessed_letters, char)

		if !char_in_string(char, word) {
			guess--
		}

		exit = true
		for _, c := range word {
			if !rune_in_slice(c, guessed_letters) {
				exit = false
			}
		}

		fmt.Println()
	}

	if guess == 0 {
		fmt.Println("You lost!")
		fmt.Println("The word was", word)
	} else {
		fmt.Println("You won!")
		fmt.Println("You had", guess, "guesses left to get", word)
	}
}
