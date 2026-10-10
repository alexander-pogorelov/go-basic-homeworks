// main package
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	StatusNew        = "новая"
	StatusInProgress = "в работе"
	StatusCompleted  = "завершена"
	StatusPaused     = "приостановлена"
)

var ErrEmptyTitle = errors.New("название задачи не может быть пустым")
var ErrTitleTooLong = errors.New("название задачи не может быть длиннее 100 символов")
var ErrInvalidStatus = errors.New("недопустимый статус задачи")

var validStatuses = []string{StatusNew, StatusInProgress, StatusCompleted, StatusPaused}

var (
	ids      = []int{1, 2, 3, 4}
	titles   = []string{"Подготовить презентацию", "Сделать ДЗ", "Купить книгу", "Пересмотреть лекцию"}
	statuses = []string{StatusNew, StatusInProgress, StatusCompleted, StatusInProgress}
)

func readLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
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
	if !isStatusValid(status) {
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

func addTask(title, status string) error {
	if err := validateTask(title, status); err != nil {
		return err
	}
	ids = append(ids, nextID())
	titles = append(titles, title)
	statuses = append(statuses, status)
	return nil
}

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

	for {
		fmt.Println()
		fmt.Println("Меню:")
		fmt.Println("  1 - показать все задачи")
		fmt.Println("  2 - отфильтровать задачи по статусам")
		fmt.Println("  3 - добавить новую задачу")
		fmt.Println("  4 - выйти")
		fmt.Print("Выберите пункт: ")
		command, success := readLine(scanner)
		if !success {
			fmt.Println("Ввод закончился, выходим.")
			return
		}
		switch command {
		case "1":
			fmt.Println()
			fmt.Println("Все задачи:")
			printTasks()
		case "2":
			fmt.Println()
			fmt.Print("Выберите статус: ")
			status, success := readLine(scanner)
			if !success {
				fmt.Println("Ввод закончился, выходим.")
				return
			}
			if !isStatusValid(status) {
				fmt.Println("Выбран невалидный статус")
				printValidStatuses()
				continue
			}
			printTasksByStatus(status)
		case "3":
			fmt.Println()
			fmt.Print("Введите название задачи: ")
			title, success := readLine(scanner)
			if !success {
				fmt.Println("Ввод закончился, выходим.")
				return
			}
			fmt.Print("Введите статус задачи: ")
			status, success := readLine(scanner)
			if !success {
				fmt.Println("Ввод закончился, выходим.")
				return
			}
			err := addTask(title, status)
			switch {
			case err == nil:
				fmt.Println("Задача успешно добавлена.")
			case errors.Is(err, ErrEmptyTitle):
				fmt.Println("Ошибка! Вы не ввели название задачи.")
			case errors.Is(err, ErrTitleTooLong):
				fmt.Println("Ошибка! Название задачи длиннее 100 символов.")
			case errors.Is(err, ErrInvalidStatus):
				fmt.Println("Ошибка! Невалидный статус задачи.")
				printValidStatuses()
			default:
				fmt.Println("Ошибка:", err)
			}
		case "4":
			fmt.Println("Выход.")
			return
		default:
			fmt.Printf("Неизвестная команда %q. Введите 1, 2, 3 или 4.\n", command)
		}
	}
}
