// Package model содержит модель задач, допустимые статусы и операции над ними.
package model

import (
	"errors"
	"fmt"
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

// NewTask создаёт задачу без ID (ID присваивается хранилищем).
func NewTask(title string, status string) (*Task, error) {
	title = strings.TrimSpace(title)
	status = strings.TrimSpace(status)

	if err := validateTask(title, status); err != nil {
		return nil, err
	}
	now := time.Now()
	return &Task{
		title:     title,
		status:    CastToStatus(status),
		createdAt: now,
		updatedAt: now,
	}, nil
}

// Update обновляет title и status задачи
func (t *Task) Update(title, status string) error {
	title = strings.TrimSpace(title)
	status = strings.TrimSpace(status)
	if err := validateTask(title, status); err != nil {
		return err
	}
	t.title = title
	t.status = CastToStatus(status)
	t.updatedAt = time.Now()
	return nil
}

// String возвращает строковое представление задачи для вывода.
func (t *Task) String() string {
	return fmt.Sprintf("%-4d %-15s %-20s %-20s %s", t.id, t.status, t.createdAt.Format(time.RFC822), t.updatedAt.Format(time.RFC822), t.title)
}

func validateTitle(title string) error {
	if title == "" {
		return ErrEmptyTitle
	}
	if utf8.RuneCountInString(title) > 100 {
		return ErrTitleTooLong
	}
	return nil
}

func validateStatus(status string) error {
	if !IsStatusValid(status) {
		return ErrInvalidStatus
	}
	return nil
}

func validateTask(title, status string) error {
	if err := validateTitle(title); err != nil {
		return fmt.Errorf("некорректное название задачи: %w", err)
	}
	if err := validateStatus(status); err != nil {
		return fmt.Errorf("некорректный статус задачи: %w", err)
	}
	return nil
}

// Status геттер для поля status задачи.
func (t *Task) Status() Status {
	return t.status
}

// Title геттер для поля title задачи.
func (t *Task) Title() string {
	return t.title
}

// SetID сеттер для поля
func (t *Task) SetID(id int) {
	t.id = id
}

// ID геттер для поля ID задачи.
func (t *Task) ID() int {
	return t.id
}
