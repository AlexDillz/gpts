// Напишите функцию CountVowels(s string) int, которая принимает строку s
// в качестве аргумента и возвращает количество гласных букв (a, e, i, o, u) в этой строке.
// Учтите как заглавные, так и строчные буквы.

// package main

// func CountVowels(s string) int {
// 	count := 0
// 	for _, r := range s{
// 		if r == 'a' || r == 'A' || r == 'e'|| r == 'E' || r == 'i' || r == 'I' || r == 'o' || r == 'O' || r == 'u' || r == 'U'{
// 			count++
// 		}
// 	}
// 	return count
// }

package main

import "fmt"

func CountVowels(s string) int {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	count := 0
	for _, r := range s {
		for _, v := range vowels {
			if r == v {
				count++
				break
			}
		}
	}
	return count
}

func main1() {
	testString := "Hello, World!"
	result := CountVowels(testString)
	fmt.Printf("The number of vowels in '%s' is: %d\n", testString, result)
}
