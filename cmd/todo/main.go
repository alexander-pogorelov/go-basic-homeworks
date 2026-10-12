// main package
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	StatusNew        = "новая"
	StatusInProgress = "в работе"
	StatusCompleted  = "завершена"
	StatusPaused     = "приостановлена"
)

var validStatuses = []string{StatusNew, StatusInProgress, StatusCompleted, StatusPaused}

var (
	ids      = []int{1, 2, 3, 4}
	titles   = []string{"Подготовить презентацию", "Сделать ДЗ", "Купить книгу", "Пересмотреть лекцию"}
	statuses = []string{StatusNew, StatusInProgress, StatusCompleted, StatusInProgress}
)

func printHeaders() {
	fmt.Printf("%-4s %-15s %s\n", "ID", "СТАТУС", "НАЗВАНИЕ")
}

func printTasks() {
	printHeaders()
	for i := range ids {
		fmt.Printf("%-4d %-15s %s\n", ids[i], statuses[i], titles[i])
	}
}

func printValidStatuses() {
	fmt.Println("Валидные статусы:")
	for i := range validStatuses {
		fmt.Printf(" - %s\n", validStatuses[i])
	}
}

func isStatusValid(status string) bool {
	for i := range validStatuses {
		if validStatuses[i] == status {
			return true
		}
	}
	return false
}

func printTasksByStatus(status string) {
	printHeaders()
	found := false
	for i := range ids {
		if status == statuses[i] {
			found = true
			fmt.Printf("%-4d %-15s %s\n", ids[i], statuses[i], titles[i])
		}
	}
	if !found {
		fmt.Println("Задач с таким статусом не найдено.")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

menu:
	for {
		fmt.Println()
		fmt.Println("Меню:")
		fmt.Println("  1 - показать все задачи")
		fmt.Println("  2 - отфильтровать задачи по статусам")
		fmt.Println("  3 - выйти")
		fmt.Print("Выберите пункт: ")

		if !scanner.Scan() {
			fmt.Println()
			fmt.Println("Ввод закончился, выходим.")
			break menu
		}

		command := scanner.Text()

		switch command {
		case "1":
			fmt.Println()
			fmt.Println("Все задачи:")
			printTasks()
		case "2":
			fmt.Println()
			fmt.Print("Выберите статус: ")
			if !scanner.Scan() {
				fmt.Println()
				fmt.Println("Ввод закончился, выходим.")
				break menu
			}
			status := scanner.Text()
			if !isStatusValid(status) {
				fmt.Println("Выбран невалидный статус")
				printValidStatuses()
				continue
			}
			printTasksByStatus(status)
		case "3":
			fmt.Println("Выход.")
			break menu
		default:
			fmt.Printf("Неизвестная команда %q. Введите 1, 2 или 3.\n", command)
		}
	}
}
