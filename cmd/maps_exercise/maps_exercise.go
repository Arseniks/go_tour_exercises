package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordsCounter := make(map[string]int)
	for _, v := range strings.Fields(s) {
		wordsCounter[v]++
	}
	return wordsCounter
}

func main() {
	wc.Test(WordCount)
}
