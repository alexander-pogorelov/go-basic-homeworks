// main package
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/alexander-pogorelov/go-basic-homeworks/internal/cli"
	"github.com/alexander-pogorelov/go-basic-homeworks/internal/model"
	"github.com/alexander-pogorelov/go-basic-homeworks/internal/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	storage.Seed(store)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("Меню:")
		fmt.Println("  1 - показать все задачи")
		fmt.Println("  2 - отфильтровать задачи по статусам")
		fmt.Println("  3 - добавить новую задачу")
		fmt.Println("  4 - выйти")
		fmt.Print("Выберите пункт: ")

		command, ok := cli.ReadLine(scanner)
		if !ok {
			fmt.Println("Ввод закончился, выходим.")
			return
		}

		switch command {
		case "1":
			fmt.Println()
			fmt.Println("Все задачи:")
			cli.PrintTasks(store)
		case "2":
			fmt.Println()
			fmt.Print("Выберите статус: ")
			status, ok := cli.ReadLine(scanner)
			if !ok {
				fmt.Println("Ввод закончился, выходим.")
				return
			}
			cli.PrintTasksByStatus(store, status)

		case "3":
			fmt.Println()
			fmt.Print("Введите название задачи: ")
			title, ok := cli.ReadLine(scanner)
			if !ok {
				fmt.Println("Ввод закончился, выходим.")
				return
			}
			fmt.Print("Введите статус задачи: ")
			status, ok := cli.ReadLine(scanner)
			if !ok {
				fmt.Println("Ввод закончился, выходим.")
				return
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

		case "4":
			fmt.Println("Выход.")
			return

		default:
			fmt.Printf("Неизвестная команда %q. Введите 1, 2, 3 или 4.\n", command)
		}
	}
}
