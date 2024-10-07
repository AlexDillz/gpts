// Напиши функцию, которая будет инвертировать строку. То есть функция должна принимать строку и возвращать новую строку, в которой символы идут в обратном порядке.

// Указания:
// Для работы с символами строки используй срез rune (чтобы корректно обрабатывать многобайтовые символы, такие как юникод).
// Используй цикл для инвертирования среза рун.
// Верни строку, собранную из инвертированного среза рун.

package main

import "fmt"

func ReverseStrings(s string) string {
	var reverse []rune
	for _, r := range s {
		reverse = append([]rune{r}, reverse...)
	}
	return string(reverse)
}

func main3() {
	testString := "Hello, World!"
	reverse := ReverseStrings(testString)
	fmt.Println(reverse)
}
