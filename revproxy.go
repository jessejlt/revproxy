package main

// Playground for writing testable proxy APIs in golang

import (
  "github.com/julienschmidt/httprouter"
  "log"
  "net/http"
  "os"
  "unicode/utf8"
)

func validateStartup() {

  host := os.Getenv("PROXY_HOST")
  if utf8.RuneCountInString(host) == 0 {
    log.Fatalf("Missing environment PROXY_HOST")
  }
}

func main() {

  validateStartup()
  router := httprouter.New()
  router.GET("/root", CollectionsHandler)

  log.Fatal(http.ListenAndServe(":8080", router))
}
