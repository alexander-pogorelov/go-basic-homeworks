package model

import "slices"

// Status представляет статус задачи.
type Status string

const (
	// StatusNew — только что созданная задача.
	StatusNew Status = "новая"
	// StatusInProgress — задача в работе.
	StatusInProgress Status = "в работе"
	// StatusCompleted — задача завершена.
	StatusCompleted Status = "завершена"
	// StatusPaused — задача приостановлена.
	StatusPaused Status = "приостановлена"
)

var validStatuses = []Status{
	StatusNew,
	StatusInProgress,
	StatusCompleted,
	StatusPaused,
}

// Statuses возвращает копию списка всех допустимых статусов задачи.
func Statuses() []Status {
	return slices.Clone(validStatuses)
}

// IsStatusValid сообщает, является ли переданная строка допустимым статусом задачи.
func IsStatusValid(status string) bool {
	st := Status(status)
	for _, s := range validStatuses {
		if s == st {
			return true
		}
	}
	return false
}

// CastToStatus приводит строку к типу Status без проверки допустимости.
// Для проверки используйте IsStatusValid.
func CastToStatus(status string) Status {
	return Status(status)
}
