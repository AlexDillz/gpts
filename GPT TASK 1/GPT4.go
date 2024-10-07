// Теперь давай немного усложним задачу.

// Создай новый интерфейс Report, который будет иметь метод Generate() string, возвращающий строку с отчетом.

// Добавь к структуре Task поле Owner, представляющее собой строку, указывающую, кто отвечает за выполнение этой задачи.

// Создай новую структуру TaskReport, которая будет реализовывать интерфейс Report. Структура должна содержать срез задач 
// и метод Generate(),
// который возвращает строку, содержащую список всех задач и их статус (выполнена или нет) в виде отчета.

package main

import (
	"fmt"
)

type Report interface {
	Generate() string
}

type Task struct {
	Description string
	Done        bool
	Owner       string
}

type TaskReport struct {
	Tasks []Task
}	

// Метод для выполнения задачи
func (t *Task) Complete() {
	t.Done = true
}

// Функция для фильтрации невыполненных задач
func IncompleteTasks(tasks []Task) []Task {
	var incomplete []Task
	for _, task := range tasks {
		if !task.Done { // Если задача не выполнена
			incomplete = append(incomplete, task)
		}
	}
	return incomplete
}

func (tr TaskReport) Generate() string{
	report := ""
	for _, task := range tr.Tasks{
		status := "Incomplete"
		if task.Done {
			status = "Complete"
		}
		report += fmt.Sprintf("Task: %s, Status: %s, Owner: %s\n", task.Description, status, task.Owner)
	}
	return report
}