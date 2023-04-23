package main

import (
	"errors"
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
	runeSliceWord2 := []rune(word)
	words2 := words
	//формируем второй массив

	//for _, val := range words {
	//	vlr := []rune(val)
	//	if len(vlr) >= 4 {
	//		vlr = vlr[:len(vlr)-1]
	//	}
	//	words2 = append(words2, string(vlr))
	//}

	for len(runeSliceWord) >= 3 {
		if contains(words, string(runeSliceWord)) {
			ch <- map[int]float64{index: float64(len(string(runeSliceWord))) * 0.01}
			return
		}
		runeSliceWord = runeSliceWord[:len(runeSliceWord)-1]
	}

	for {
		words2, isChanged := createWords2Slice(&words2)
		if !isChanged {
			break
		}
		if contains(words2, string(runeSliceWord2)) {
			ch <- map[int]float64{index: float64(len(string(runeSliceWord2))) * 0.007}
			return
		}
	}
	ch <- map[int]float64{index: 0.0}
}

func createWords2Slice(words2 *[]string) ([]string, bool) {

	var isChanged bool
	for i, val := range *words2 {
		vlr := []rune(val)
		isChanged = false
		if len(vlr) >= 4 {
			(*words2)[i] = string(vlr[:len(vlr)-1])
			isChanged = true
		}
	}
	return *words2, isChanged
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
func convertPhrasesMapToSlice(mp map[string]int64) ([]string, map[string]int64) {
	var res []string
	maxFrequencyMap := make(map[string]int64)
	for key, freq := range mp {
		res = append(res, key)

		mfKey, _ := getKeyFromMap(maxFrequencyMap)
		if mfKey != "" {
			if maxFrequencyMap[mfKey] < freq {
				maxFrequencyMap[key] = freq
				delete(maxFrequencyMap, mfKey)
			}
		} else {
			maxFrequencyMap[key] = freq
		}
	}
	return res, maxFrequencyMap
}

func deleteFromSlice(slice []string, val string) []string {
	for i, v := range slice {
		if v == val {
			// Use append to remove the element at index i
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}
func getKeyFromMap(mp map[string]int64) (string, error) {
	for key := range mp {
		return key, nil
	}
	return "", errors.New("не найдено значение с переданным ключом")
}

func getSortList() []map[string]int64 {
	inpMp := map[string]int64{
		"Без своей порции кофе утром я просто не могу начать день":                               100,
		"Кофе - это не только напиток, это мой образ жизни":                                      200,
		"Я увлечена кофе как искусством и наслаждаюсь каждой чашечкой":                           300,
		"Каждый глоток кофе - это для меня настоящий райский кайф":                               400,
		"Как же прекрасно начинать свое утро с ароматной чашечки кофе":                           400,
		"Я не могу себе представить свою жизнь без кофе - это моя зависимость":                   401,
		"Чай - это не просто напиток, это настоящая медитация и умиротворение для меня":          400,
		"Я люблю чай за его способность расслабить и укрепить мой организм":                      200,
		"Чашка горячего чая - это мой способ побаловать себя в конце дня":                        300,
		"Как же великолепно наслаждаться чашечкой ароматного чая в уютной атмосфере своего дома": 100,
		"кофе это мое утро":      150,
		"кофе это моя страсть":   200,
		"чай это зеленый листик": 300,
		"чай это напиток богов":  12,
	}

	phrases, maxFreqMap := convertPhrasesMapToSlice(inpMp)
	phrase, err := getKeyFromMap(maxFreqMap)
	if err != nil {
		panic(err.Error())
	}
	phrases = deleteFromSlice(phrases, phrase)

	var res = []map[string]int64{maxFreqMap}

	for range phrases {
		phrase = mostSimilarPhrase(phrase, phrases)
		res = append(res, map[string]int64{phrase: inpMp[phrase]})
		phrases = deleteFromSlice(phrases, phrase)
	}
	return res
}

func main() {
	fmt.Println(getSortList())
}
