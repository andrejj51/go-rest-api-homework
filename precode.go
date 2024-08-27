package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

// Обработчик для получения всех задач:
func getTasks(w http.ResponseWriter, r *http.Request) {
	// сериализуем данные из слайса tasks
	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// запись в заголовок
	w.Header().Set("Content-Type", "application/json")

	// статус OK
	w.WriteHeader(http.StatusOK)

	// запись сериализованных данных json в тело ответа
	w.Write(resp)
}

// Обработчик для отправки задачи на сервер:

// Обработчик для получения задачи по ID:

// Обработчик удаления задачи по ID:

func main() {
	r := chi.NewRouter()

	// Обработчик для получения всех задач:
	r.Get("/tasks", getTasks)

	// Обработчик для отправки задачи на сервер:
	//r.Post("/taksks", postTask)

	// Обработчик для получения задачи по ID:
	//r.Get("/tasks/{id}", getTask)

	// Обработчик удаления задачи по ID:
	//r.Delete("/tasks/{id}", deleteTask)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
