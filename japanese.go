// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// have a hash map of the correct answers
// have a list of hiragana
// have a list of letters
// have a generate random function
// have a get correct answers function
// slice should be in format of string slice with each as individual letters

func main() {
	hiragana := []string{"あ", "い", "う", "え", "お", "か", "き", "く", "け", "こ", "さ", "し", "す", "せ", "そ"}

	words := hiragana
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
	fmt.Print("insert y value here: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
	fmt.Println(correctAnswer(words))
}

func correctAnswer(input []string) []string {
	m := map[string]string{
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
	}

	output := []string{}
	for _, s := range input {
		output = append(output, m[s])
	}
	return output
}
