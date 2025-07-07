package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        // log the error and exit
        fmt.Println("Server failed:", err)
    }
}