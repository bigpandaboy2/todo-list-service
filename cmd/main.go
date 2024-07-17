package main

import (
    "log"
    "net/http"
    "todolist/routers"

    _ "todolist/docs" 

    httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
    router := routers.InitializeRouter()

    // Swagger endpoint
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    log.Fatal(http.ListenAndServe(":8080", router))
}
