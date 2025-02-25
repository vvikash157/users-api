package controllers

import (
	"Login/models"
	"Login/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var task = models.Task{}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdTask, err := service.CreateTask(task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	status := r.URL.Query().Get("status")

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	tasks, totalRecords, err := service.GetTask(page, pageSize, status, task)
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"tasks":        tasks,
		"totalRecords": totalRecords,
		"page":         page,
		"pageSize":     pageSize,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task, err := service.GetTaskById(params["id"], task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task, err := service.UpdateTask(params["id"], updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("task extracted : ", params["id"])
	taskID, err := strconv.Atoi(params["id"])
	fmt.Println("task extracted as integer : ", taskID)
	if err != nil {
		http.Error(w, `{"message": "Invalid task ID"}`, http.StatusBadRequest)
		return
	}

	if err := service.DeleteTask(taskID); err != nil {
		http.Error(w, `{"message": "Failed to delete task"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})

}
