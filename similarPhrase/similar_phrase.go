package main

import (
	"fmt"
	"math"
	"strings"
)

// Функция для разделения фразы на отдельные слова
func tokenize(phrase string) []string {
	return strings.Split(strings.ToLower(phrase), " ")
}

// Функция для создания слайса слов из списка фраз
func createWordSlice(phrases []string) []string {
	wordSlice := make([]string, 0)
	wordSet := make(map[string]bool)
	for _, phrase := range phrases {
		words := tokenize(phrase)
		for _, word := range words {
			if !wordSet[word] {
				wordSlice = append(wordSlice, word)
				wordSet[word] = true
			}
		}
	}
	return wordSlice
}

func getVector(phrase string, wordSlice []string) []float64 {
	words := tokenize(phrase)
	vector := make([]float64, len(wordSlice))
	vectorCh := make(chan map[int]float64, len(wordSlice))

	for i := range wordSlice {
		go getVectorVal(words, wordSlice[i], vectorCh, i)
	}

	for range wordSlice {
		for i, v := range <-vectorCh {
			vector[i] = v
		}
	}

	return vector
}

func getVectorVal(words []string, word string, ch chan<- map[int]float64, index int) {
	runeSliceWord := []rune(word)
	for len(runeSliceWord) >= 3 {
		if contains(words, string(runeSliceWord)) {
			ch <- map[int]float64{index: float64(len(string(runeSliceWord))) * 0.01}
			return
		}
		runeSliceWord = runeSliceWord[:len(runeSliceWord)-1]
	}
	ch <- map[int]float64{index: 0.0}
}

// Функция для проверки, содержится ли слово в списке
func contains(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

// Функция для вычисления косинусного сходства между двумя векторами
func cosineSimilarity(vectorA, vectorB []float64) float64 {
	dotProduct := dot(vectorA, vectorB)
	magnitudeA := magnitude(vectorA)
	magnitudeB := magnitude(vectorB)
	return dotProduct / (magnitudeA * magnitudeB)
}

// Функция для вычисления скалярного произведения двух векторов
func dot(vectorA, vectorB []float64) float64 {
	var dotProduct float64
	for i := range vectorA {
		dotProduct += vectorA[i] * vectorB[i]
	}
	return dotProduct
}

// Функция для вычисления модуля вектора
func magnitude(vector []float64) float64 {
	var sum float64
	for _, v := range vector {
		sum += v * v
	}
	return math.Sqrt(sum)
}

// Функция для нахождения наиболее похожей фразы
func mostSimilarPhrase(searchPhrase string, phrases []string) string {
	wordSet := createWordSlice(phrases)
	searchVector := getVector(searchPhrase, wordSet)
	highestSimilarity := -1.0
	mostSimilarPhrase := ""
	for _, phrase := range phrases {
		phraseVector := getVector(phrase, wordSet)
		similarity := cosineSimilarity(searchVector, phraseVector)
		if similarity > highestSimilarity {
			highestSimilarity = similarity
			mostSimilarPhrase = phrase
		}
	}
	return mostSimilarPhrase
}

func main() {
	phrases := []string{
		"кофе это мое утро",
		"кофе это моя страсть",
		"чай это зеленый листик",
		"чай это напиток богов",
	}
	fmt.Println(mostSimilarPhrase("чай бог", phrases))
}
