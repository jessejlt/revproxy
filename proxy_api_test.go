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
  fmt.Fprintln(w, `{"url":"http://foo.com","name":"bar"}`)
}

func TestGetRoot(t *testing.T) {

  s := httptest.NewServer(http.HandlerFunc(getRootHandler))
  defer s.Close()
  os.Setenv("PROXY_HOST", s.URL)

  _, root, err := GetRoot("")
  if err != nil {
    t.Error(err)
  }

  if root.Name != "bar" {
    t.Errorf("Expected root.Name = bar, received %v", root.Name)
  }
}
