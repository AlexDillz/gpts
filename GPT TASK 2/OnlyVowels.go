package main

import "fmt"

func OnlyVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y', 'A', 'E', 'I', 'O', 'U', 'Y'}
	var result []rune
	for _, r := range s {
		for _, v := range vowels {
			if v == r {
				result = append(result, r)
				break
			}
		}
	}
	return string(result)
}

func main2() {
	testString := "Hello, World!"
	result := OnlyVowels(testString)
	fmt.Printf(result)
}
