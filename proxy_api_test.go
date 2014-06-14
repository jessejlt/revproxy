package main

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "os"
  "testing"
)

func getRootHandler(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintln(w, `{"foo": "bar"}`)
}

func TestGetRoot(t *testing.T) {

  s := httptest.NewServer(http.HandlerFunc(getRootHandler))
  defer s.Close()
  os.Setenv("PROXY_HOST", s.URL)

  _, _, err := GetRoot("")
  if err != nil {
    t.Error(err)
  }
}
