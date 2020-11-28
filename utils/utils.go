package utils

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

func HexDecode(s string) []byte {
	v, err := hex.DecodeString(s)
	if err != nil {
		errStr := fmt.Sprintf("failed to decode hex %s: %v", s, err)
		panic(errStr)
	}
	return v
}

// BuildCorpus builds corpus for a given text
func BuildCorpus(text string) map[rune]float64 {
	c := make(map[rune]float64)
	for _, char := range text {
		c[char] = c[char] + 1
	}
	total := utf8.RuneCountInString(text)
	for char := range c {
		c[char] = c[char] / float64(total)
	}
	return c
}

// CorpusFromFrile builds corpus from a file
func CorpusFromFrile(name string) map[rune]float64 {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		panic("failed to open corpus file: " + err.Error())
	}
	return BuildCorpus(string(text))
}

// GetDefaultCorpus returns a default corpus built from
// "Alice's Adventures in Wonderland by Lewis Carroll"
// https://www.gutenberg.org/ebooks/11
func GetDefaultCorpus() map[rune]float64 {
	pwd, _ := os.Getwd()
	corpusFilePath := pwd + "/../fixtures/alice.txt"
	c := CorpusFromFrile(corpusFilePath)
	return c
}

// ScoreEnglish calculate a score for an english word according to the Corpus
func ScoreEnglish(text string, c map[rune]float64) float64 {
	var score float64
	for _, char := range text {
		score += c[char]
	}
	total := utf8.RuneCountInString(text)
	return score / float64(total)
}
