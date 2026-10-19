package main

import (
	"fmt"
	"log"
	"os"

	"library-app/library"
	"library-app/storage"
)

func main() {
	const storagePath = "books.json"

	appCatalog, err := storage.ImportDataFromFile(storagePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("[INFO] Локальное хранилище отсутствует. Инициализирована чистая база данных.")
			appCatalog = library.New()
		} else {
			log.Fatalf("[FATAL] Не удалось инициализировать приложение из файла: %v", err)
		}
	} else {
		fmt.Println("[SUCCESS] Данные книжного фонда успешно импортированы.")
	}

	LaunchConsoleInterface(appCatalog, storagePath)
}
