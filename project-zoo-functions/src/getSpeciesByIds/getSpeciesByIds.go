package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
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

type ZooData struct {
  Species []Species `json:"species"`
}

func getSpeciesByIds(ids ...string) ([]Species, error) {
  // ler o arquivo JSON e armazenar os dados em uma variável do tipo ZooData
  data, err := ioutil.ReadFile("./data/zoo_data.json")
  if err != nil {
    return nil, err
  }
  var dataMap ZooData
  json.Unmarshal(data, &dataMap)

  // criar o array de retorno da função
  var result []Species

  // adicionar as espécies correspondentes aos ids passados como parâmetro
  for _, id := range ids {
    found := false
    for _, species := range dataMap.Species {
      if species.Id == id {
        found = true
        result = append(result, species)
        break
      }
    }
    if !found {
      return nil, fmt.Errorf("species with id %s not found", id)
    }
  }
  return result, nil
}

func main() {
  id1 := "0938aa23-f153-4937-9f88-4858b24d6bce"
  id2 := "ef3778eb-2844-4c7c-b66c-f432073e1c6b"

  // chamar a função getSpeciesByIds passando os ids como parâmetros
  species, err := getSpeciesByIds(id1, id2)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("espécies com id", id1, "e", id2, ":", species)
  }
}
