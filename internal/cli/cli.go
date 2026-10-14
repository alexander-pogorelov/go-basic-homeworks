// Package cli отвечает за ввод и вывод задач в консольном интерфейсе.
package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/alexander-pogorelov/go-basic-homeworks/internal/task"
)

// ReadLine читает одну строку из сканера и обрезает пробелы по краям.
func ReadLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}

func printHeaders() {
	fmt.Printf("%-4s %-15s %s\n", "ID", "СТАТУС", "НАЗВАНИЕ")
}

// PrintTasks печатает список всех задач в табличном виде.
func PrintTasks() {
	printHeaders()
	for i := range task.Count() {
		fmt.Printf("%-4d %-15s %s\n", task.ID(i), task.Status(i), task.Title(i))
	}
}

// PrintValidStatuses печатает список допустимых статусов.
func PrintValidStatuses() {
	fmt.Println("Валидные статусы:")
	for _, status := range task.ValidStatuses() {
		fmt.Printf(" - %s\n", status)
	}
}

// PrintTasksByStatus печатает задачи с указанным статусом.
func PrintTasksByStatus(status string) {
	printHeaders()
	found := false
	for i := range task.Count() {
		if status == task.Status(i) {
			found = true
			fmt.Printf("%-4d %-15s %s\n", task.ID(i), task.Status(i), task.Title(i))
		}
	}
	if !found {
		fmt.Println("Задач с таким статусом не найдено.")
	}
}
