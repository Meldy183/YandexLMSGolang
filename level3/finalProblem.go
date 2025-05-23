package main

import (
	"fmt"
	"sort"
	"strings"
)

func AnalyzeText(text string) {
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ToLower(text)
	//fmt.Println(text)
	mas := strings.Split(text, " ")
	//fmt.Println(mas)
	length := len(mas)
	fmt.Printf("Количество слов: %d\n", length)
	mapX := CountingSort(mas)
	cnt := 0
	for range mapX {
		cnt++
	}
	fmt.Printf("Количество уникальных слов: %d\n", cnt)
	top := getTopWords(mapX, 5)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", top[0], mapX[top[0]])
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for i := range 5 {
		fmt.Printf("\"%s\": %d раз\n", top[i], mapX[top[i]])
	}
}

type x struct {
	key string
	val int
}

func getTopWords(wordMap map[string]int, n int) []string {
	ans := make([]string, 0)
	da := sorting(wordMap)
	for _, word := range da {
		if n == 0 {
			break
		}
		n--
		ans = append(ans, word.key)
	}
	return ans
}

func sorting(wordMap map[string]int) []x {
	ans := make([]x, 0)
	for k, v := range wordMap {
		ans = append(ans, x{k, v})
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].val == ans[j].val {
			return ans[i].key < ans[j].key
		}
		return ans[i].val > ans[j].val
	})
	return ans
}
