package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {

}

func countLines(fileName string) (int, int, int) {
	var err error
	var numberOfLines int
	var numberOfCharacters int
	var numberOfWords int
	numberOfLines = 0
	numberOfCharacters = 0
	numberOfWords = 0

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// The ReadString function tells the program to read until the first occurrence of '\n' in the input
		// Multiple calls to ReadString() mean that you are reading a file line by line
		line, err := r.ReadString('\n')

		// The for loop terminates when you get the io.EOF error, signifying that there is nothing left to read from the input file
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}

		// Each time the bufio reader reads a new line, the value of the NumberOfLines variable is increased by one
		numberOfLines++
		r := regexp.MustCompile("\\S+")
		for range r.FindAllString(line, -1) {
			numberOfWords++
		}
		// Number of characters is implemented with the help of len() function that returns the number of characters in a given string
		numberOfCharacters += len(line)
	}
	return numberOfLines, numberOfWords, numberOfCharacters
}
