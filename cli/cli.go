package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library-app/library"
	"library-app/storage"
)

func readStringInput(scanner *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := scanner.ReadString('\n')
	return strings.TrimSpace(text)
}

func LaunchConsoleInterface(catalog *library.Library, saveFilePath string) {
	consoleScanner := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- СИСТЕМА УЧЕТА КНИГ ---")
		fmt.Println("1) Зарегистрировать новую книгу")
		fmt.Println("2) Добавить карточку читателя")
		fmt.Println("3) Оформить выдачу книги")
		fmt.Println("4) Оформить возврат книги")
		fmt.Println("5) Вывести весь книжный фонд")
		fmt.Println("6) Список зарегистрированных читателей")
		fmt.Println("0) Завершить работу")
		fmt.Print("Выберите действие: ")

		var menuChoice string
		fmt.Scanln(&menuChoice)
		menuChoice = strings.TrimSpace(menuChoice)

		commandCode, err := strconv.Atoi(menuChoice)
		if err != nil {
			fmt.Println("Ошибка: некорректная команда. Попробуйте еще раз.")
			continue
		}

		switch commandCode {
		case 1:
			bookTitle := readStringInput(consoleScanner, "Введите название книги: ")
			bookAuthor := readStringInput(consoleScanner, "Укажите автора: ")
			yearRaw := readStringInput(consoleScanner, "Год издания: ")

			publishYear, err := strconv.Atoi(yearRaw)
			if err != nil {
				fmt.Println("Ошибка: некорректный формат года.")
				continue
			}

			if _, err := catalog.RegisterBook(bookTitle, bookAuthor, publishYear); err != nil {
				fmt.Printf("Не удалось добавить книгу: %v\n", err)
			} else {
				fmt.Println("Успешно: книга внесена в базу данных.")
			}

		case 2:
			firstName := readStringInput(consoleScanner, "Имя посетителя: ")
			lastName := readStringInput(consoleScanner, "Фамилия посетителя: ")

			if _, err := catalog.RegisterReader(firstName, lastName); err != nil {
				fmt.Printf("Не удалось зарегистрировать пользователя: %v\n", err)
			} else {
				fmt.Println("Успешно: карточка читателя создана.")
			}

		case 3:
			targetBookStr := readStringInput(consoleScanner, "Укажите ID книги: ")
			targetReaderStr := readStringInput(consoleScanner, "Укажите ID читателя: ")

			idBook, _ := strconv.Atoi(targetBookStr)
			idReader, _ := strconv.Atoi(targetReaderStr)

			if err := catalog.IssueBookToReader(idBook, idReader); err != nil {
				fmt.Printf("Отказ в выдаче: %v\n", err)
			} else {
				fmt.Println("Успешно: книга закреплена за читателем.")
			}

		case 4:
			targetBookStr := readStringInput(consoleScanner, "Укажите ID возвращаемой книги: ")
			idBook, _ := strconv.Atoi(targetBookStr)

			if err := catalog.ReturnBook(idBook); err != nil {
				fmt.Printf("Ошибка при возврате: %v\n", err)
			} else {
				fmt.Println("Успешно: книга принята обратно.")
			}

		case 5:
			fmt.Println("\n=== Список доступной литературы ===")
			allBooks := catalog.GetAllBooks()
			if len(allBooks) == 0 {
				fmt.Println("Каталог пуст.")
			}
			for _, bookItem := range allBooks {
				fmt.Printf(" • %v\n", bookItem)
			}

		case 6:
			fmt.Println("\n=== Список посетителей ===")
			allReaders := catalog.GetAllReaders()
			if len(allReaders) == 0 {
				fmt.Println("Нет зарегистрированных читателей.")
			}
			for _, person := range allReaders {
				fmt.Printf(" • %v\n", person)
			}

		case 0:
			fmt.Println("Синхронизация локальной базы данных... Выход.")
			if err := storage.ExportDataToFile(saveFilePath, catalog); err != nil {
				fmt.Printf("Критическая ошибка при сохранении данных: %v\n", err)
			}
			return

		default:
			fmt.Println("Выбран неизвестный пункт меню. Повторите попытку.")
		}
	}
}
