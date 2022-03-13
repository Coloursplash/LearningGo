package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func str_in_slice(str string, s []string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	words := []string{"apple", "banana", "orange", "lemon"}

	rand.Seed(time.Now().UTC().UnixNano())
	word := words[rand.Intn(len(words))]

	guessed_letters := make([]string, 0)

	guess := 6
	exit := false

	fmt.Println(word)
	for guess > 0 && !exit {
		fmt.Println(guess, "guesses left")

		for _, c := range word {
			if str_in_slice(string(c), guessed_letters) {
				fmt.Print(c)
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()

		fmt.Println("Enter a letter:")
		var char rune
		fmt.Scanln(&char)

		if str_in_slice(string(char), guessed_letters) || unicode.IsLetter(char) {
			continue
		}

		if !strings.ContainsRune(word, char) {
			guess--
		}

		exit = true
		for _, c := range word {
			if !str_in_slice(string(c), guessed_letters) {
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
