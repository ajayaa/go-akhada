package main

import "fmt"
import "strings"

func unique_words(str string) map[string]bool {
	m := make(map[string]bool)
	words := strings.Split(str, " ")
	for _, word := range words {
		if word == "" {
			fmt.Println("found an empty string")
			continue
		} else {
			m[word] = true
		}
	}
	return m
}

func main() {
	m := unique_words("aj   cj fs aj cj fs sf")
	for key, _ := range m {
		fmt.Println(key)
	}
}
