// Package storage содержит реализации хранилищ задач,
package storage

import (
	"fmt"

	"github.com/alexander-pogorelov/go-basic-homeworks/internal/model"
)

// Seed заполняет хранилище демонстрационными задачами.
func Seed(s *MemoryStorage) {
	seed := []struct{ title, status string }{
		{"Подготовить презентацию", string(model.StatusNew)},
		{"Сделать ДЗ", string(model.StatusInProgress)},
		{"Купить книгу", string(model.StatusCompleted)},
		{"Пересмотреть лекцию", string(model.StatusInProgress)},
	}
	for _, item := range seed {
		t, err := model.NewTask(item.title, item.status)
		if err != nil {
			panic(fmt.Sprintf("seed task %q: %v", item.title, err))
		}
		s.Add(t)
	}
}
