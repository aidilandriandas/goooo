package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Jenkins CI/CD Docker Compose!")
}

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("Server running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
