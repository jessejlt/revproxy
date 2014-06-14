package main

import (
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

type api struct {
  method, endpoint, authToken string
}

func proxyRequest(apiDef *api) (*http.Response, []byte, error) {

  host := os.Getenv("PROXY_HOST")
  client := &http.Client{}

  req, err := http.NewRequest(apiDef.method, host+apiDef.endpoint, nil)
  if err != nil {
    log.Print(err)
    return nil, nil, err
  }

  req.Header.Add("Authorization", apiDef.authToken)

  resp, err := client.Do(req)
  if err != nil {
    log.Print(err)
    return nil, nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Print(err)
    return nil, nil, err
  }

  return resp, body, err
}

// GetRoot retrieves the user's root directory
func GetRoot(authToken string) (*http.Response, []byte, error) {

  apiDef := &api{method: "GET", endpoint: "/settings", authToken: authToken}
  return proxyRequest(apiDef)
}
