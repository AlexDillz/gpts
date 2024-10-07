// Создай структуру Task, которая будет содержать поля:

// Description (строка) — описание задания.
// Done (булево значение) — статус выполнения задания (выполнено или нет).
// Добавь метод Complete(), который будет отмечать задание как выполненное (Done = true).

// Затем напиши функцию, которая принимает список заданий (срез Task) и возвращает только те задания,
// которые еще не выполнены (Done == false).

// Требования:
// Создать структуру Task с нужными полями.
// Реализовать метод Complete() для изменения статуса задачи.
// Реализовать функцию, которая фильтрует список задач по полю Done.

package main

type Task struct {
	Description string
	Done        bool
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
