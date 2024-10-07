// Напиши функцию, которая удаляет все гласные буквы из строки. Функция должна принимать строку,
// удалять из неё все гласные (как строчные, так и заглавные) и возвращать результат.

// Описание задачи:

// На вход поступает строка.
// На выходе — строка без гласных букв.
// Шаги для решения:
// Создай срез рун для строки.
// Пробегись по каждой руне в строке и проверь, является ли она гласной.
// Используй условие, чтобы исключить гласные.
// Возвращай строку без гласных.

package main

import "fmt"

func DeleteVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y', 'A', 'E', 'I', 'O', 'U', 'Y'}
	var result []rune

	for _, r := range s {
		isVowel := false
		for _, v := range vowels {
			if r == v {
				isVowel = true
				break
			}
		}
		if !isVowel {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	testString := "Hello, World!"
	result := DeleteVowels(testString)
	fmt.Println(result)
}
