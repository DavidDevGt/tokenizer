package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Tokenizer avanzado
type AdvancedTokenizer struct {}

// usamos regex y tokenizamos por palabras incluyendo los signos de puntuación y apostrofes
func (t AdvancedTokenizer) Tokenize(text string) []string {
	regExp := regexp.MustCompile(`[\w']+|[.,!?;]`)
	matches := regExp.FindAllString(text, -1)

	for i, match := range matches {
		matches[i] = strings.ToLower(match)
	}

	return matches
}

// Esto cuenta la frecuencia de las palabras
func FrequencyCounter(tokens []string) map[string]int {
	frequencies := make(map[string]int)
	for _, token:= range tokens {
		frequencies[token]++
	}
	return frequencies
}

// esto quita palabras que no aportan significado
func RemoveStopWords(tokens []string, stopWords []string) []string {
	filteredTokens := []string{}
	for _, token := range tokens {
		if !contains(stopWords, token) {
			filteredTokens = append(filteredTokens, token)
		}
	}
	return filteredTokens
}

// Esto verifica si una palabra está contenida en una lista
func contains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

func main() {}