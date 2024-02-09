package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Tokenizer avanzado
type AdvancedTokenizer struct{}

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
	for _, token := range tokens {
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

func main() {
	// Texto pa tokenizar
	text := `The Battle of the Bulge, also known as the Ardennes Counteroffensive, was the last major German offensive campaign on the Western Front during World War II. It was launched through the densely forested Ardennes region of Wallonia in eastern Belgium, northeast France, and Luxembourg, towards the end of 1944. The surprise attack caught the Allied forces completely off guard. United States forces bore the brunt of the attack and incurred their highest casualties of any operation during the war. The battle also severely depleted Germany's armored forces, and they were largely unable to replace them. The Battle of the Bulge was the largest and bloodiest single battle fought by the United States in World War II and the third-deadliest campaign in American history. Following the Battle of the Bulge, the Allies launched a successful invasion of Germany, leading to the eventual surrender of Nazi forces.`

	// creo una instancia
	tokenizer := AdvancedTokenizer{}

	// tokenizo el texto
	tokens := tokenizer.Tokenize(text)

	// Palabras pa excluir del texto tokenizado
	stopWords := []string{"the", "was", "of", "and", "in", "to", "their", "they", "by", "any", "during", "also", "as", "on", "at", "towards", "end", "through", "towards", "towards", "off", "were", "unable", "largely", "them", "largest", "single", "fought", "third", "history", "following", "launched", "leading", "eventual"}

	// elimino las palabras que no aportan significado
	tokensWithoutStopWords := RemoveStopWords(tokens, stopWords)

	// cuento la frecuencia de cada token
	frequencies := FrequencyCounter(tokensWithoutStopWords)

	// imprimo todo
	fmt.Println("Token Frequencies:")
	for token, freq := range frequencies {
		fmt.Printf("%s: %d\n", token, freq)
	}
}
