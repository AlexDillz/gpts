// Создай интерфейс Worker, который будет содержать метод Work() string. У структур Person и Robot (которую нужно создать)
// должен быть метод Work, который возвращает описание их работы. Для Person это может быть профессия,
// а для Robot — "I am building something."

// Создай функцию, которая принимает список объектов, реализующих интерфейс Worker, и вызывает метод Work() для каждого из них.

// Требования:
// Определить интерфейс Worker.
// Определить структуру Robot с нужным методом.
// Реализовать функцию, которая принимает срез интерфейсов и выводит результат работы каждого объекта.

package main

type Worker interface {
	Work() string
}

type Person struct {
	Name       string
	Occupation string
}

type Robot struct {
	Model string
}

func (p Person) Work() string {
	return p.Occupation
}

func (r Robot) Work() string {
	return "I am building something."
}

func SliceInterface(people []Worker) []string {
	var results []string
	for _, worker := range people {
		results = append(results, worker.Work())
	}
	return results
}
