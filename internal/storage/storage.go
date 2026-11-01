// Package storage содержит реализации хранилищ задач,
package storage

import (
	"errors"

	"github.com/alexander-pogorelov/go-basic-homeworks/internal/model"
)

// ErrTaskNotFound возвращается, если задача с указанным ID не найдена.
var ErrTaskNotFound = errors.New("задача не найдена")

// MemoryStorage хранит задачи в памяти.
type MemoryStorage struct {
	tasks  []*model.Task
	byID   map[int]*model.Task
	nextID int
}

// NewMemoryStorage создаёт пустое хранилище задач
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make([]*model.Task, 0),
		byID:   make(map[int]*model.Task),
		nextID: 1,
	}
}

// Add сохраняет задачу в хранилище, присваивая ей следующий свободный ID.
func (s *MemoryStorage) Add(task *model.Task) {
	task.SetID(s.nextID)
	s.tasks = append(s.tasks, task)
	s.byID[task.ID()] = task
	s.nextID++
}

// GetByID возвращает задачу с указанным ID.
func (s *MemoryStorage) GetByID(id int) (*model.Task, error) {
	t, ok := s.byID[id]
	if !ok {
		return nil, ErrTaskNotFound
	}
	return t, nil
}

// Delete удаляет задачу с указанным ID.
func (s *MemoryStorage) Delete(id int) error {
	if _, ok := s.byID[id]; !ok {
		return ErrTaskNotFound
	}
	delete(s.byID, id)
	for i, t := range s.tasks {
		if t.ID() == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return ErrTaskNotFound
}

// Update сохраняет изменения существующей задачи.
func (s *MemoryStorage) Update(task *model.Task) error {
	id := task.ID()

	if _, ok := s.byID[id]; !ok {
		return ErrTaskNotFound
	}

	s.byID[id] = task

	for i, t := range s.tasks {
		if t.ID() == id {
			s.tasks[i] = task
			return nil
		}
	}

	return ErrTaskNotFound
}

// GetAll возвращает копию слайса всех задач в порядке их добавления.
func (s *MemoryStorage) GetAll() []*model.Task {
	clone := make([]*model.Task, len(s.tasks))
	copy(clone, s.tasks)
	return clone
}

// FilterByStatus возвращает задачи с указанным статусом
func (s *MemoryStorage) FilterByStatus(status model.Status) []*model.Task {
	result := make([]*model.Task, 0)
	for _, t := range s.tasks {
		if t.Status() == status {
			result = append(result, t)
		}
	}
	return result
}
