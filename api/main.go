package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type AppStatus struct {
	State   string `json:"state"`
	Service string `json:"service_name"`
}

type SystemDetails struct {
	BuildVersion string `json:"build_version"`
	CurrentTime  string `json:"current_time"`
}

func sendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("[ERROR] Не удалось закодировать JSON-ответ: %v", err)
	}
}

func StatusController(w http.ResponseWriter, r *http.Request) {
	payload := AppStatus{
		State:   "running",
		Service: "Library System Backend",
	}
	sendJSON(w, http.StatusOK, payload)
}

func InformationController(w http.ResponseWriter, r *http.Request) {
	payload := SystemDetails{
		BuildVersion: "1.0.0-rc1",
		CurrentTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	sendJSON(w, http.StatusOK, payload)
}

func CatchAllNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte("Requested endpoint not found")); err != nil {
		log.Printf("[WARNING] Ошибка отправки 404: %v", err)
	}
}

func main() {
	// Маршрутизация
	http.HandleFunc("/healthcheck", StatusController)
	http.HandleFunc("/api/v1/info", InformationController)
	http.HandleFunc("/notfound", CatchAllNotFound)

	serverPort := ":8080"

	log.Printf("[INFO] Инициализация веб-сервера на порту %s...", serverPort)
	fmt.Printf("--> Пинг-сервис:  http://localhost%s/healthcheck\n", serverPort)
	fmt.Printf("--> Инфо-сервис:  http://localhost%s/api/v1/info\n", serverPort)
	fmt.Printf("--> Тест 404:     http://localhost%s/notfound\n", serverPort)

	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatalf("[FATAL] Критическая ошибка при работе сервера: %v", err)
	}
}
