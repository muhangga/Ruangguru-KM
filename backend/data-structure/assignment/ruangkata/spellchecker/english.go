package spellchecker

import (
	"bufio"
	"os"
	"strings"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/ruangkata/downloader"
)

const url string = "https://gist.githubusercontent.com/fikriauliya/25c107ae057d3dc87abdb5dfb2f330b6/raw/639d54126b27c57651ac42ef8ece3c5c92a3d76b/en"
const filePath string = "./english-words.csv"

type spellchecker struct {
	words map[string]bool
}

func parseFile() (map[string]bool, error) {
	words := make(map[string]bool)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		words[word] = true
	}

	return words, nil
}

func NewEnglishSpellChecker() (SpellChecker, error) {
	err := downloader.Download(url, filePath)
	if err != nil {
		return nil, err
	}

	words, err := parseFile()
	if err != nil {
		return nil, err
	}

	return &spellchecker{words}, nil
}

func (s *spellchecker) CheckWord(word string) bool {

	spell := strings.ToLower(word)

	if _, ok := s.words[spell]; ok {
		return true
	}

	return false
}

func (s *spellchecker) CheckSentence(sentence string) (validWords []string, invalidWords []string) {

	// split valid and invalid words
	// hello world this is a sentence tapi mengandung kata2 nggak bener
	words := strings.Split(sentence, " ")

	for _, word := range words {
		validWords = append(validWords, word)

		if !s.CheckWord(word) {
			invalidWords = append(invalidWords, word)
		}


	}

	return validWords, invalidWords
}