package main

import "regexp"

func main() {
	exampleLine := "The boy is going to school"
	print(wordCounter(exampleLine))
}

func wordCounter(line string) int {
	numbersOfWords := 0
	r := regexp.MustCompile("\\S+")
	for range r.FindAllString(line, -1) {
		numbersOfWords++
	}
	return numbersOfWords
}
