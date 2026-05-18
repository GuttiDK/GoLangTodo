package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	mu     sync.Mutex
	todos  = make(map[int]Todo)
	nextID = 1
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		getTodoByID(w, id)
	case http.MethodPut:
		updateTodo(w, r, id)
	case http.MethodDelete:
		deleteTodo(w, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	list := make([]Todo, 0, len(todos))
	for _, t := range todos {
		list = append(list, t)
	}
	mu.Unlock()
	writeJSON(w, list)
}

func getTodoByID(w http.ResponseWriter, id int) {
	mu.Lock()
	t, ok := todos[id]
	mu.Unlock()
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	writeJSON(w, t)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &t); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	mu.Lock()
	t.ID = nextID
	nextID++
	todos[t.ID] = t
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	writeJSON(w, t)
}

func updateTodo(w http.ResponseWriter, r *http.Request, id int) {
	var payload Todo
	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	mu.Lock()
	t, ok := todos[id]
	if !ok {
		mu.Unlock()
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if payload.Title != "" {
		t.Title = payload.Title
	}
	t.Completed = payload.Completed
	todos[id] = t
	mu.Unlock()
	writeJSON(w, t)
}

func deleteTodo(w http.ResponseWriter, id int) {
	mu.Lock()
	_, ok := todos[id]
	if ok {
		delete(todos, id)
	}
	mu.Unlock()
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
