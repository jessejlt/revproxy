package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

// Root is a user's root directory
type Root struct {
  URL  string `json:"url"`
  Name string `json:"name"`
}

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
func GetRoot(authToken string) (*http.Response, *Root, error) {

  apiDef := &api{method: "GET", endpoint: "/root", authToken: authToken}
  resp, body, err := proxyRequest(apiDef)
  if err != nil {
    return nil, nil, err
  }

  r := new(Root)
  err = json.Unmarshal(body, &r)
  if err != nil {
    return nil, nil, err
  }

  return resp, r, nil
}
