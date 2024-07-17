package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "todolist/models"

    "github.com/gorilla/mux"
)

var tasks = make(map[string]models.Task)

// CreateTask creates a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)
    task.ID = time.Now().Format("20060102150405")
    tasks[task.ID] = task
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"id": task.ID})
}

// UpdateTask updates an existing task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]
    var updatedTask models.Task
    json.NewDecoder(r.Body).Decode(&updatedTask)
    if _, ok := tasks[id]; ok {
        tasks[id] = updatedTask
        w.WriteHeader(http.StatusNoContent)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}

// DeleteTask deletes a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]
    if _, ok := tasks[id]; ok {
        delete(tasks, id)
        w.WriteHeader(http.StatusNoContent)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}

// MarkTaskDone marks a task as done
func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]
    if task, ok := tasks[id]; ok {
        task.Completed = true
        tasks[id] = task
        w.WriteHeader(http.StatusNoContent)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}

// ListTasks lists tasks by status
func ListTasks(w http.ResponseWriter, r *http.Request) {
    status := r.URL.Query().Get("status")
    var filteredTasks []models.Task
    for _, task := range tasks {
        if status == "done" && task.Completed {
            filteredTasks = append(filteredTasks, task)
        } else if status != "done" && !task.Completed {
            filteredTasks = append(filteredTasks, task)
        }
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(filteredTasks)
}

// HealthCheck checks the health of the service
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}