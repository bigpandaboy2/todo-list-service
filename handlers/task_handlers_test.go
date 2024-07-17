package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
    "todolist/models"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
    router := mux.NewRouter()
    router.HandleFunc("/api/todo-list/tasks", CreateTask).Methods("POST")

    task := models.Task{
        Title:    "Test Task",
        ActiveAt: time.Now(),
    }

    body, _ := json.Marshal(task)
    req, err := http.NewRequest("POST", "/api/todo-list/tasks", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateTask(t *testing.T) {
    router := mux.NewRouter()
    router.HandleFunc("/api/todo-list/tasks/{id}", UpdateTask).Methods("PUT")

    task := models.Task{
        Title:    "Test Task Updated",
        ActiveAt: time.Now(),
    }

    tasks["1"] = task
    body, _ := json.Marshal(task)
    req, err := http.NewRequest("PUT", "/api/todo-list/tasks/1", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteTask(t *testing.T) {
    router := mux.NewRouter()
    router.HandleFunc("/api/todo-list/tasks/{id}", DeleteTask).Methods("DELETE")

    tasks["1"] = models.Task{
        Title:    "Test Task",
        ActiveAt: time.Now(),
    }

    req, err := http.NewRequest("DELETE", "/api/todo-list/tasks/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestMarkTaskDone(t *testing.T) {
    router := mux.NewRouter()
    router.HandleFunc("/api/todo-list/tasks/{id}/done", MarkTaskDone).Methods("PUT")

    tasks["1"] = models.Task{
        Title:    "Test Task",
        ActiveAt: time.Now(),
    }

    req, err := http.NewRequest("PUT", "/api/todo-list/tasks/1/done", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestListTasks(t *testing.T) {
    router := mux.NewRouter()
    router.HandleFunc("/api/todo-list/tasks", ListTasks).Methods("GET")

    tasks["1"] = models.Task{
        Title:    "Test Task",
        ActiveAt: time.Now(),
    }

    req, err := http.NewRequest("GET", "/api/todo-list/tasks?status=active", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
}