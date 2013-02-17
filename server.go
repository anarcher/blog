package main

import (
    "net/http"
    "log"
)

func init() {
    http.Handle("/", http.FileServer(http.Dir("./public")))
}

func main() {
    log.Println("Starting server :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
