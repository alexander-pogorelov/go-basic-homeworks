// Package cli отвечает за ввод и вывод задач в консольном интерфейсе.
package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/alexander-pogorelov/go-basic-homeworks/internal/model"
	"github.com/alexander-pogorelov/go-basic-homeworks/internal/storage"
)

// ReadLine читает одну строку из сканера и обрезает пробелы по краям.
func ReadLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}

func printHeaders() {
	fmt.Printf("%-4s %-15s %-20s %-20s %s\n", "ID", "СТАТУС", "СОЗДАНА", "ОБНОВЛЕНА", "НАЗВАНИЕ")
}

// PrintTasks печатает все задачи.
func PrintTasks(tasks []*model.Task) {
	printHeaders()
	for _, t := range tasks {
		fmt.Println(t)
	}
}

// PrintValidStatuses печатает список допустимых статусов.
func PrintValidStatuses() {
	fmt.Println("Валидные статусы:")
	for _, status := range model.Statuses() {
		fmt.Printf(" - %s\n", status)
	}
}

// PrintTasksByStatus печатает задачи с указанным статусом.
func PrintTasksByStatus(s *storage.MemoryStorage, status model.Status) {
	tasks := s.FilterByStatus(status)
	if len(tasks) == 0 {
		fmt.Println("Задач с таким статусом не найдено.")
		return
	}
	PrintTasks(tasks)
}
