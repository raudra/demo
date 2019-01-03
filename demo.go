package main

import (
    "log"
    "net/http"
    "demo/config/routes"
)

func main() {
    router := routes.NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}