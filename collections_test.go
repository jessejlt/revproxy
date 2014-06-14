package main

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "os"
  "testing"
)

func handler(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintln(w, `{"url":"http://foo.com","name":"bar"}`)
}

func wrapper(w http.ResponseWriter, r *http.Request) {
  CollectionsHandler(w, r, nil)
}

func TestCollectionsHandler(t *testing.T) {

  // Setup response handler for outbound request
  proxyServer := httptest.NewServer(http.HandlerFunc(handler))
  defer proxyServer.Close()
  os.Setenv("PROXY_HOST", proxyServer.URL)

  // Setup route wrapper
  routeServer := httptest.NewServer(http.HandlerFunc(wrapper))
  defer routeServer.Close()

  // Request
  uri := routeServer.URL + "/root"
  client := &http.Client{}
  req, err := http.NewRequest("GET", uri, nil)
  if err != nil {
    t.Error(err)
  }

  req.Header.Set("Authorization", "---")
  resp, err := client.Do(req)
  if err != nil {
    t.Error(err)
  }

  defer resp.Body.Close()
  if resp.StatusCode != 200 {
    t.Errorf("Expected StatusCode=200, recieved %v", resp.StatusCode)
  }
}

func TestCollectionsRequiresAuth(t *testing.T) {

  s := httptest.NewServer(http.HandlerFunc(wrapper))
  defer s.Close()

  resp, err := http.Get(s.URL)
  if err != nil {
    t.Error(err)
  }

  if resp.StatusCode != 401 {
    t.Errorf("Expected StatusCode=401, recieved %v", resp.StatusCode)
  }
}
