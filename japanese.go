package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Clear screen.
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// Builds list of hiragana to practice.
	hr1 := []string{"あ", "い", "う", "え", "お"}
	hr2 := []string{"か", "き", "く", "け", "こ"}
	hr3 := []string{"さ", "し", "す", "せ", "そ"}
	hr4 := []string{"た", "ち", "つ", "て", "と"}
	hr5 := []string{"な", "に", "ぬ", "ね", "の"}
	h := [][]string{hr1, hr2, hr3, hr4, hr5}
	words := make([]string, 0)
	for _, v := range h {
		words = append(words, v...)
	}

	// Shuffles list.
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	// Quiz the user for each word.
	for _, word := range words {
		fmt.Println(word)
		fmt.Print("Type in your answer:")
		complete := false
		for !complete {
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			ans, err := correctAnswer(word)
			if err != nil {
				fmt.Println(err)
			}
			if input.Text() == ans {
				fmt.Println("Correct!")
				complete = true
			} else {
				fmt.Println("Incorrect. Try again.")
			}
		}
	}

	fmt.Println("Practice complete.")
}

func correctAnswer(input string) (string, error) {
	hiraganaToLetters := map[string]string{
		"あ": "a",
		"い": "i",
		"う": "u",
		"え": "e",
		"お": "o",
		"か": "ka",
		"き": "ki",
		"く": "ku",
		"け": "ke",
		"こ": "ko",
		"さ": "sa",
		"し": "shi",
		"す": "su",
		"せ": "se",
		"そ": "so",
		"た": "ta",
		"ち": "chi",
		"つ": "tsu",
		"て": "te",
		"と": "to",
		"な": "na",
		"に": "ni",
		"ぬ": "nu",
		"ね": "ne",
		"の": "no",
	}

	val, ok := hiraganaToLetters[input]
	if !ok {
		return "", fmt.Errorf("Input not found")
	}
	return val, nil
}
