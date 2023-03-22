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
	hr6 := []string{"は", "ひ", "ふ", "へ", "ほ"}
	hr7 := []string{"ま", "み", "む", "め", "も"}
	hr8 := []string{"や", "ゆ", "よ"}
	hr9 := []string{"ら", "り", "る", "れ", "ろ"}
	hr10 := []string{"わ", "を"}
	hr11 := []string{"ん"}
	dakuOn1 := []string{"が", "ぎ", "ぐ", "げ", "ご"}
	dakuOn2 := []string{"ざ", "じ", "ず", "ぜ", "ぞ"}
	dakuOn3 := []string{"だ", "ぢ", "づ", "で", "ど"}
	dakuOn4 := []string{"ば", "び", "ぶ", "べ", "ぼ"}
	dakuOn5 := []string{"ぱ", "ぴ", "ぷ", "ぺ", "ぽ"}
	h := [][]string{hr1, hr2, hr3, hr4, hr5, hr6, hr7, hr8, hr9, hr10, hr11}
	allDakuOn := [][]string{dakuOn1, dakuOn2, dakuOn3, dakuOn4, dakuOn5}
	var chosen [][]string
	if false {
		chosen = h
	} else {
		chosen = allDakuOn
	}

	words := make([]string, 0)
	for _, v := range chosen {
		words = append(words, v...)
	}

	// Shuffles list.
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	numCorrectFirstTry := 0
	t1 := time.Now()
	// Quiz the user for each word.
	for _, word := range words {
		fmt.Println(word)
		fmt.Print("Type in your answer:")
		complete := false
		tries := 0
		for !complete {
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			ans, err := correctAnswer(word)
			if err != nil {
				fmt.Println(err)
			}
			if input.Text() == "?" {
				fmt.Println(ans)
			} else if input.Text() == ans {
				fmt.Println("Correct!")
				complete = true
				if tries == 0 {
					numCorrectFirstTry += 1
				}
			} else {
				fmt.Println("Incorrect. Try again.")
				tries += 1
			}
		}
	}
	t2 := time.Now()

	fmt.Printf("Practice complete. Percent correct first try: %g\n", float64(numCorrectFirstTry)/float64(len(words)))
	fmt.Printf("Time elapsed: %v\n", t2.Sub(t1))
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
		"は": "ha",
		"ひ": "hi",
		"ふ": "fu",
		"へ": "he",
		"ほ": "ho",
		"ま": "ma",
		"み": "mi",
		"む": "mu",
		"め": "me",
		"も": "mo",
		"や": "ya",
		"ゆ": "yu",
		"よ": "yo",
		"ら": "ra",
		"り": "ri",
		"る": "ru",
		"れ": "re",
		"ろ": "ro",
		"わ": "wa",
		"を": "wo",
		"ん": "n",
		"が": "ga",
		"ぎ": "gi",
		"ぐ": "gu",
		"げ": "ge",
		"ご": "go",
		"ざ": "za",
		"じ": "ji",
		"ず": "zu",
		"ぜ": "ze",
		"ぞ": "zo",
		"だ": "da",
		"ぢ": "ji",
		"づ": "zu",
		"で": "de",
		"ど": "do",
		"ば": "ba",
		"び": "bi",
		"ぶ": "bu",
		"べ": "be",
		"ぼ": "bo",
		"ぱ": "pa",
		"ぴ": "pi",
		"ぷ": "pu",
		"ぺ": "pe",
		"ぽ": "po",
	}

	val, ok := hiraganaToLetters[input]
	if !ok {
		return "", fmt.Errorf("Input not found")
	}
	return val, nil
}
