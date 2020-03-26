package main

import (
	"fmt"
	"strings"
)

// Split Mutiple Delimiters
func Split(r rune) bool {
	return r == ',' || r == '.' || r == ' '
}

func wordCount(str string) {
	var mapWord map[string]int
	mapWord = make(map[string]int, 100)
	strSlice := strings.FieldsFunc(str, Split)
	for _, value := range strSlice {
		_, ok := mapWord[value]
		if ok {
			mapWord[value]++
		} else {
			mapWord[value] = 1
		}
	}
	for key, value := range mapWord {
		fmt.Printf("Word: %s\t\tCount: %d\n", key, value)
	}
}

func main() {
	str := "Hello World, This is an sample text for Hello World."
	wordCount(str)
}
