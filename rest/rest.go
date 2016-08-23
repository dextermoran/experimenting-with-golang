package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Payload struct {
  Stuff Data
}

type Data struct {
  Fruit Fruits
  Veggie Veggetables
}

type Fruits map[string]int
type Veggetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
  response, err := getJsonResponse()
  if err != nil {
    panic(err)
  }

  fmt.Fprintf(w, string(response))
  fmt.Printf("json served! \n")
}

func getJsonResponse() ([]byte, error) {
  fruits := make(map[string]int)
  fruits["Apples"] = 25
  fruits["Oranges"] = 10

  vegetables := make(map[string]int)
  vegetables["Carrots"] = 21
  vegetables["Peppers"] = 0

  d := Data{fruits, vegetables}
  p := Payload{d}

  return json.MarshalIndent(p, "", "  ")
}

func main() {
    http.HandleFunc("/", serveRest)
    http.ListenAndServe("localhost:3000", nil)
}
