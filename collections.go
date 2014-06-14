package main

import (
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "log"
  "net/http"
)

// CollectionsHandler proxies /root to TODO
func CollectionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

  authToken := r.Header.Get("Authorization")
  if authToken == "" {
    http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
    return
  }

  resp, root, err := GetRoot(authToken)
  if err != nil {

    log.Print(err)
    statusCode := 500
    if resp != nil {
      statusCode = resp.StatusCode
    }

    errMessage := http.StatusText(statusCode)
    http.Error(w, errMessage, statusCode)
    return
  }

  js, err := json.Marshal(root)
  if err != nil {
    log.Printf("Failed to unmarshal upstream response %v", err)
    http.Error(w, "", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
