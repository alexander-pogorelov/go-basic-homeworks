// main package
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/alexander-pogorelov/go-basic-homeworks/internal/cli"
	"github.com/alexander-pogorelov/go-basic-homeworks/internal/model"
	"github.com/alexander-pogorelov/go-basic-homeworks/internal/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	storage.Seed(store)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		command, ok := cli.ReadLine(scanner)
		if !ok {
			fmt.Println("Ввод закончился, выходим.")
			return
		}

		switch command {
		case "1":
			handleListAll(store)
		case "2":
			if !handleListByID(store, scanner) {
				return
			}
		case "3":
			if !updateTaskStatus(store, scanner) {
				return
			}
		case "4":
			if !deleteTask(store, scanner) {
				return
			}
		case "5":
			if !handleListByStatus(store, scanner) {
				return
			}
		case "6":
			if !addTask(store, scanner) {
				return
			}
		case "7":
			fmt.Println("Выход.")
			return
		default:
			fmt.Printf("Неизвестная команда %q. Введите 1, 2, 3, 4, 5, 6 или 7.\n", command)
		}
	}
}

func printMenu() {
	fmt.Println()
	fmt.Println("Меню:")
	fmt.Println("  1 - показать все задачи")
	fmt.Println("  2 - показать задачу по ID")
	fmt.Println("  3 - изменить статус задачи")
	fmt.Println("  4 - удалить задачу")
	fmt.Println("  5 - отфильтровать задачи по статусам")
	fmt.Println("  6 - добавить новую задачу")
	fmt.Println("  7 - выйти")
	fmt.Print("Выберите пункт: ")
}

func handleListAll(store *storage.MemoryStorage) bool {
	fmt.Println()
	fmt.Println("Все задачи:")
	cli.PrintTasks(store.GetAll())
	return true
}

func handleListByID(store *storage.MemoryStorage, scanner *bufio.Scanner) bool {
	fmt.Println()
	fmt.Print("Введите ID задачи: ")
	id, ok := cli.ReadLine(scanner)
	if !ok {
		fmt.Println("Ввод закончился, выходим.")
		return false
	}
	n, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Вы ввели не число.")
		return true
	}
	t, err := store.GetByID(n)
	if err != nil {
		fmt.Println(err)
		return true
	}
	fmt.Println(t)

	return true
}

func updateTaskStatus(store *storage.MemoryStorage, scanner *bufio.Scanner) bool {
	fmt.Println()
	fmt.Print("Введите ID задачи: ")
	id, ok := cli.ReadLine(scanner)
	if !ok {
		fmt.Println("Ввод закончился, выходим.")
		return false
	}
	n, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Вы ввели не число.")
		return true
	}
	t, err := store.GetByID(n)
	if err != nil {
		fmt.Println(err)
		return true
	}

	fmt.Println()
	fmt.Print("Выберите статус: ")

	status, ok := cli.ReadLine(scanner)
	if !ok {
		fmt.Println("Ввод закончился, выходим.")
		return false
	}

	err = t.Update(t.Title(), status)
	if err != nil {
		fmt.Println(err)
		return true
	}

	fmt.Println("Задача успешно обновлена.")

	return true
}

func deleteTask(store *storage.MemoryStorage, scanner *bufio.Scanner) bool {
	fmt.Println()
	fmt.Print("Введите ID задачи: ")
	id, ok := cli.ReadLine(scanner)
	if !ok {
		fmt.Println("Ввод закончился, выходим.")
		return false
	}
	n, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Вы ввели не число.")
		return true
	}
	_, err = store.GetByID(n)
	if err != nil {
		fmt.Println(err)
		return true
	}

	err = store.Delete(n)
	if err != nil {
		fmt.Println(err)
		return true
	}

	fmt.Println("Задача успешно удалена.")

	return true
}

func handleListByStatus(store *storage.MemoryStorage, scanner *bufio.Scanner) bool {
	fmt.Println()
	fmt.Print("Выберите статус: ")

	status, ok := cli.ReadLine(scanner)
	if !ok {
		fmt.Println("Ввод закончился, выходим.")
		return false
	}

	if !model.IsStatusValid(status) {
		fmt.Println("Выбран невалидный статус")
		cli.PrintValidStatuses()
		return true
	}

	cli.PrintTasksByStatus(store, model.CastToStatus(status))

	return true
}

func addTask(store *storage.MemoryStorage, scanner *bufio.Scanner) bool {
	fmt.Println()
	fmt.Print("Введите название задачи: ")
	title, ok := cli.ReadLine(scanner)
	if !ok {
		fmt.Println("Ввод закончился, выходим.")
		return false
	}

	fmt.Print("Введите статус задачи: ")
	status, ok := cli.ReadLine(scanner)
	if !ok {
		fmt.Println("Ввод закончился, выходим.")
		return false
	}

	t, err := model.NewTask(title, status)

	switch {
	case err == nil:
		store.Add(t)
		fmt.Println("Задача успешно добавлена.")
	case errors.Is(err, model.ErrEmptyTitle):
		fmt.Println("Ошибка! Вы не ввели название задачи.")
	case errors.Is(err, model.ErrTitleTooLong):
		fmt.Println("Ошибка! Название задачи длиннее 100 символов.")
	case errors.Is(err, model.ErrInvalidStatus):
		fmt.Println("Ошибка! Невалидный статус задачи.")
		cli.PrintValidStatuses()
	default:
		fmt.Println("Ошибка:", err)
	}

	return true
}
