package model

import "slices"

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

func Statuses() []Status {
	return slices.Clone(validStatuses)
}

func IsStatusValid(status string) bool {
	st := Status(status)
	for _, s := range validStatuses {
		if s == st {
			return true
		}
	}
	return false
}

func CastToStatus(status string) Status {
	return Status(status)
}
