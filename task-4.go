package main

import "fmt"

func main() {
	var m map[string]int
	for _, word := range []string{"hello", "world", "from", "the", "best", "language",
		"in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// вызов, попытки взаимодействовать с неинициализированной мапой приводит к панике.
// Избежать паники можно 2 способами, всегда инициализировать мапу или проверять мапу на nil
