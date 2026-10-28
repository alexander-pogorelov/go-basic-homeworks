// Package model содержит модель задач, допустимые статусы и операции над ними.
package model

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
	"unicode/utf8"
)

// ErrEmptyTitle возвращается, если название задачи пустое.
var ErrEmptyTitle = errors.New("название задачи не может быть пустым")

// ErrTitleTooLong возвращается, если название длиннее 100 символов.
var ErrTitleTooLong = errors.New("название задачи не может быть длиннее 100 символов")

// ErrInvalidStatus возвращается, если статус не входит в список допустимых.
var ErrInvalidStatus = errors.New("недопустимый статус задачи")

// Task — задача с названием, статусом и временными метками создания и обновления.
type Task struct {
	id        int
	title     string
	status    Status
	createdAt time.Time
	updatedAt time.Time
}

// NewTask создаёт задачу с указанными названием и статусом.
// Возвращает ошибку, если название или статус не прошли валидацию.
func NewTask(title string, status string) (*Task, error) {
	if err := validateTask(title, status); err != nil {
		return nil, err
	}
	task := &Task{
		id:        nextID(),
		title:     title,
		status:    CastToStatus(status),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	return task, nil
}

// String возвращает строковое представление задачи для вывода.
func (t Task) String() string {
	return fmt.Sprintf("%-4d %-15s %-20s %-20s %s", t.id, t.status, t.createdAt.Format(time.RFC822), t.updatedAt.Format(time.RFC822), t.title)
}

var tasks []Task

func init() {
	tasks = append(tasks, mustTask("Подготовить презентацию", string(StatusNew)))
	tasks = append(tasks, mustTask("Сделать ДЗ", string(StatusInProgress)))
	tasks = append(tasks, mustTask("Купить книгу", string(StatusCompleted)))
	tasks = append(tasks, mustTask("Пересмотреть лекцию", string(StatusInProgress)))
}

func mustTask(title, status string) Task {
	t, err := NewTask(title, status)
	if err != nil {
		panic(fmt.Sprintf("seed task %q: %v", title, err))
	}
	return *t
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
	t, err := NewTask(title, status)
	if err != nil {
		return err
	}
	tasks = append(tasks, *t)
	return nil
}

// Tasks возвращает копию списка всех задач.
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
