// main package
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/alexander-pogorelov/go-basic-homeworks/internal/cli"
	"github.com/alexander-pogorelov/go-basic-homeworks/internal/model"
	"github.com/alexander-pogorelov/go-basic-homeworks/internal/storage"
)

var ErrInvalidID = errors.New("ID должен быть целым числом")

func main() {
	store := storage.NewMemoryStorage()
	storage.Seed(store)
	scanner := bufio.NewScanner(os.Stdin)

	if err := run(store, scanner); err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Println("Ввод закончился, выходим.")
			return
		}
		fmt.Println("Ошибка:", err)
	}
}

func run(store *storage.MemoryStorage, scanner *bufio.Scanner) error {
	for {
		printMenu()
		command, err := readLine(scanner)
		if err != nil {
			return err
		}
		switch command {
		case "1":
			if err := handleListAll(store); err != nil {
				return err
			}
		case "2":
			if err := handleListByID(store, scanner); err != nil {
				return err
			}
		case "3":
			if err := updateTaskStatus(store, scanner); err != nil {
				return err
			}
		case "4":
			if err := deleteTask(store, scanner); err != nil {
				return err
			}
		case "5":
			if err := handleListByStatus(store, scanner); err != nil {
				return err
			}
		case "6":
			if err := addTask(store, scanner); err != nil {
				return err
			}
		case "7":
			fmt.Println("Выход.")
			return nil
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

func readLine(scanner *bufio.Scanner) (string, error) {
	s, ok := cli.ReadLine(scanner)
	if !ok {
		return "", io.EOF
	}
	return s, nil
}

func readTaskID(scanner *bufio.Scanner) (int, error) {
	raw, err := readLine(scanner)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(raw)
	if err != nil {
		return 0, ErrInvalidID
	}
	return id, nil
}

func handleListAll(store *storage.MemoryStorage) error {
	fmt.Println()
	fmt.Println("Все задачи:")
	cli.PrintTasks(store.GetAll())
	return nil
}

func handleListByID(store *storage.MemoryStorage, scanner *bufio.Scanner) error {
	fmt.Println()
	fmt.Print("Введите ID задачи: ")
	n, err := readTaskID(scanner)
	if err != nil {
		return handleError(err)
	}
	t, err := store.GetByID(n)
	if err != nil {
		return handleError(err)
	}
	cli.PrintTasks([]*model.Task{t})
	return nil
}

func updateTaskStatus(store *storage.MemoryStorage, scanner *bufio.Scanner) error {
	fmt.Println()
	fmt.Print("Введите ID задачи: ")
	n, err := readTaskID(scanner)
	if err != nil {
		return handleError(err)
	}
	t, err := store.GetByID(n)
	if err != nil {
		return handleError(err)
	}
	fmt.Println()
	fmt.Print("Выберите статус: ")
	status, err := readLine(scanner)
	if err != nil {
		return handleError(err)
	}
	if !model.IsStatusValid(status) {
		fmt.Println("Выбран невалидный статус")
		cli.PrintValidStatuses()
		return nil
	}
	if err := t.Update(t.Title(), status); err != nil {
		return handleError(err)
	}
	fmt.Println("Задача успешно обновлена.")
	return nil
}

func deleteTask(store *storage.MemoryStorage, scanner *bufio.Scanner) error {
	fmt.Println()
	fmt.Print("Введите ID задачи: ")
	n, err := readTaskID(scanner)
	if err != nil {
		return handleError(err)
	}
	if err := store.Delete(n); err != nil {
		return handleError(err)
	}
	fmt.Println("Задача успешно удалена.")
	return nil
}

func handleListByStatus(store *storage.MemoryStorage, scanner *bufio.Scanner) error {
	fmt.Println()
	fmt.Print("Выберите статус: ")
	status, err := readLine(scanner)
	if err != nil {
		return handleError(err)
	}
	if !model.IsStatusValid(status) {
		fmt.Println("Выбран невалидный статус")
		cli.PrintValidStatuses()
		return nil
	}
	cli.PrintTasksByStatus(store, model.CastToStatus(status))
	return nil
}

func addTask(store *storage.MemoryStorage, scanner *bufio.Scanner) error {
	fmt.Println()
	fmt.Print("Введите название задачи: ")
	title, err := readLine(scanner)
	if err != nil {
		return handleError(err)
	}

	fmt.Print("Введите статус задачи: ")
	status, err := readLine(scanner)
	if err != nil {
		return handleError(err)
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
	return nil
}

func handleError(err error) error {
	switch {
	case errors.Is(err, io.EOF):
		return err
	case errors.Is(err, ErrInvalidID):
		fmt.Println("Вы ввели не число.")
		return nil
	case errors.Is(err, storage.ErrTaskNotFound):
		fmt.Println("Задача не найдена")
		return nil
	default:
		fmt.Println(err)
		return nil
	}
}
