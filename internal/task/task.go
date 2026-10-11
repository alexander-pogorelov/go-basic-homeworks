// Package task содержит модель задач, допустимые статусы и операции над ними.
package task

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
)

const (
	// StatusNew — только что созданная задача.
	StatusNew = "новая"
	// StatusInProgress — задача в работе.
	StatusInProgress = "в работе"
	// StatusCompleted — задача завершена.
	StatusCompleted = "завершена"
	// StatusPaused — задача приостановлена.
	StatusPaused = "приостановлена"
)

// ErrEmptyTitle возвращается, если название задачи пустое.
var ErrEmptyTitle = errors.New("название задачи не может быть пустым")

// ErrTitleTooLong возвращается, если название длиннее 100 символов.
var ErrTitleTooLong = errors.New("название задачи не может быть длиннее 100 символов")

// ErrInvalidStatus возвращается, если статус не входит в список допустимых.
var ErrInvalidStatus = errors.New("недопустимый статус задачи")

var validStatuses = []string{StatusNew, StatusInProgress, StatusCompleted, StatusPaused}

var (
	ids      = []int{1, 2, 3, 4}
	titles   = []string{"Подготовить презентацию", "Сделать ДЗ", "Купить книгу", "Пересмотреть лекцию"}
	statuses = []string{StatusNew, StatusInProgress, StatusCompleted, StatusInProgress}
)

// IsStatusValid сообщает, входит ли статус в список допустимых.
func IsStatusValid(status string) bool {
	for i := range validStatuses {
		if validStatuses[i] == status {
			return true
		}
	}
	return false
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
	for _, id := range ids {
		if id > maxID {
			maxID = id
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
	ids = append(ids, nextID())
	titles = append(titles, title)
	statuses = append(statuses, status)
	return nil
}

// ValidStatuses возвращает копию списка допустимых статусов.
func ValidStatuses() []string {
	return slices.Clone(validStatuses)
}

// ID возвращает идентификатор задачи по индексу.
func ID(i int) int {
	return ids[i]
}

// Title возвращает название задачи по индексу.
func Title(i int) string {
	return titles[i]
}

// Status возвращает статус задачи по индексу.
func Status(i int) string {
	return statuses[i]
}

// Count возвращает количество задач.
func Count() int {
	return len(ids)
}
