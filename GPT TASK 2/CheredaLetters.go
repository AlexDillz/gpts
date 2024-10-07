// Напиши функцию, которая преобразует заданную строку так, чтобы каждая буква чередовалась между заглавной и строчной.
// Если символ — это не буква, оставь его без изменений.

// Пример работы функции:

// Ввод: "Go is great!"
// Вывод: "Go Is GrEaT!"

package main

import (
	"fmt"
	"unicode"
)

func CheredaLetters(s string) string {
	spisok := []rune(s)
	var result []rune
	count := 0
	for _, r := range spisok {
		if r == ' ' {
			result = append(result, r) // Добавляем пробел в результат
			continue                   // Переходим к следующей рунe
		}
		if count%2 == 0 { // Если позиция четная
			result = append(result, unicode.ToLower(r)) // Преобразуем в нижний регистр
		} else { // Если позиция нечетная
			result = append(result, unicode.ToUpper(r)) // Преобразуем в верхний регистр
		}
		count++ // Увеличиваем счетчик
	}

	return string(result) // Возвращаем результат в виде строки
}

func main() {
	testString := "avada ke davra"
	result := CheredaLetters(testString)
	fmt.Println(result)
}
