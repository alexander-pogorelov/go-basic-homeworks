// Package cli отвечает за ввод и вывод задач в консольном интерфейсе.
package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/alexander-pogorelov/go-basic-homeworks/internal/model"
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

// PrintTasks печатает список всех задач в табличном виде.
func PrintTasks() {
	printHeaders()
	for _, t := range model.Tasks() {
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
func PrintTasksByStatus(status string) {
	st := model.CastToStatus(status)
	printHeaders()
	found := false
	for _, t := range model.Tasks() {
		if st == t.Status() {
			found = true
			fmt.Println(t)
		}
	}
	if !found {
		fmt.Println("Задач с таким статусом не найдено.")
	}
}
