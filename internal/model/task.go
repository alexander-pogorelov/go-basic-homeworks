// Package model содержит модель задач, допустимые статусы и операции над ними.
package model

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
)

// ErrEmptyTitle возвращается, если название задачи пустое.
var ErrEmptyTitle = errors.New("название задачи не может быть пустым")

// ErrTitleTooLong возвращается, если название длиннее 100 символов.
var ErrTitleTooLong = errors.New("название задачи не может быть длиннее 100 символов")

// ErrInvalidStatus возвращается, если статус не входит в список допустимых.
var ErrInvalidStatus = errors.New("недопустимый статус задачи")

type Task struct {
	id     int
	title  string
	status Status
}

func NewTask(id int, title string, status Status) Task {
	return Task{
		id:     id,
		title:  title,
		status: status,
	}
}

func (t Task) String() string {
	return fmt.Sprintf("%-4d %-15s %s", t.id, t.status, t.title)
}

var tasks = []Task{
	NewTask(1, "Подготовить презентацию", StatusNew),
	NewTask(2, "Сделать ДЗ", StatusInProgress),
	NewTask(3, "Купить книгу", StatusCompleted),
	NewTask(4, "Пересмотреть лекцию", StatusInProgress),
}

func validateTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return ErrEmptyTitle
	}
	if utf8.RuneCountInString(title) > 100 {
		return ErrTitleTooLong
	}
	return nil
}

func validateStatus(status string) error {
	status = strings.TrimSpace(status)
	if !IsStatusValid(status) {
		return ErrInvalidStatus
	}
	return nil
}

func validateTask(title, status string) error {
	err := validateTitle(title)
	if err != nil {
		return fmt.Errorf("некорректное название задачи: %w", err)
	}
	err = validateStatus(status)
	if err != nil {
		return fmt.Errorf("некорректный статус задачи: %w", err)
	}
	return nil
}

func nextID() int {
	maxID := 0
	for _, task := range tasks {
		if task.id > maxID {
			maxID = task.id
		}
	}
	return maxID + 1
}

// AddTask добавляет задачу с указанными названием и статусом.
func AddTask(title, status string) error {
	title = strings.TrimSpace(title)
	status = strings.TrimSpace(status)
	if err := validateTask(title, status); err != nil {
		return err
	}
	tasks = append(tasks, NewTask(nextID(), title, CastToStatus(status)))
	return nil
}

func Tasks() []Task {
	return slices.Clone(tasks)
}

// Status геттер для поля status задачи.
func (t Task) Status() Status {
	return t.status
}

// Title геттер для поля title задачи.
func (t Task) Title() string {
	return t.title
}
