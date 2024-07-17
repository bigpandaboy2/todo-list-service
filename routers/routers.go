package routers

import (
    "todolist/handlers"

    "github.com/gorilla/mux"
    httpSwagger "github.com/swaggo/http-swagger"
)

// InitializeRouter initializes the router
func InitializeRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/api/todo-list/tasks", handlers.CreateTask).Methods("POST")
    router.HandleFunc("/api/todo-list/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    router.HandleFunc("/api/todo-list/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
    router.HandleFunc("/api/todo-list/tasks/{id}/done", handlers.MarkTaskDone).Methods("PUT")
    router.HandleFunc("/api/todo-list/tasks", handlers.ListTasks).Methods("GET")
    router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

    // Swagger endpoint
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    return router
}