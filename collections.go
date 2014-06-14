package main

import (
  "github.com/julienschmidt/httprouter"
  "log"
  "net/http"
)

// CollectionsHandler proxies /appsettings/collections to TODO
func CollectionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

  authToken := r.Header.Get("Authorization")
  resp, body, err := GetRoot(authToken)
  if err != nil {

    log.Print(err)
    statusCode := 500
    if resp != nil {
      statusCode = resp.StatusCode
    }

    errMessage := http.StatusText(statusCode)
    if body != nil {
      errMessage = string(body[:])
    }
    http.Error(w, errMessage, statusCode)
  }

}
