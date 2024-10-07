// Создай структуру Person, которая будет иметь следующие поля:

// Name (строка),
// Age (целое число),
// Occupation (строка).
// Напиши функцию, которая принимает список таких структур и возвращает список имен людей, которым больше 30 лет.

// Требования:
// Структура должна быть экспортируемой.
// Функция должна корректно обрабатывать пустой список.

package main

type Person struct {
	Name       string
	Age        int
	Occupation string
}

func GetNamesOlderThan30(people []Person) []string {
	var result []string
	for _, p := range people {
		if p.Age > 30 {
			result = append(result, p.Name)
		}
	}
	return result
}

// для теста +
// func main() {
// 	people := []Person{
// 		{Name: "Alice", Age: 25, Occupation: "Engineer"},
// 		{Name: "Bob", Age: 35, Occupation: "Doctor"},
// 		{Name: "Charlie", Age: 40, Occupation: "Teacher"},
// 	}

// 	names := GetNamesOlderThan30(people)
// 	fmt.Println(names) // Ожидаемый вывод: [Bob Charlie]
// }
