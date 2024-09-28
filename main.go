package main

import (
    "log"
    "net/http"
    "blog/api"
)

func main() {
    http.HandleFunc("/", handler.Handler)
    log.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}