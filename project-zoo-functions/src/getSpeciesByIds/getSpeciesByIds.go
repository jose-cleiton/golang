package main

import (
  
  "fmt"
  "github.com/jose-cleiton/fofunc"

 
)

type Species struct {
  Id          string   `json:"id"`
  Name        string   `json:"name"`
  Popularity  int      `json:"popularity"`
  Location    string   `json:"location"`
  Availability []string `json:"availability"`
  Residents   []struct {
    Name string `json:"name"`
    Sex  string `json:"sex"`
    Age  int    `json:"age"`
  } `json:"residents"`
}



func main() {
  id1 := "0938aa23-f153-4937-9f88-4858b24d6bce"
  id2 := "ef3778eb-2844-4c7c-b66c-f432073e1c6b"

  // chamar a função getSpeciesByIds passando os ids como parâmetros
  species, err :=   area.GetSpeciesByIds(id1, id2)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("espécies com id", id1, "e", id2, ":", species)
  }
}

