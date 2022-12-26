package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)

type Species struct {
  Id   int    `json:"id"`
  Name string `json:"name"`
}
type SpeciesArray []Species

func getSpeciesByIds(ids ...int) (SpeciesArray, error) {
  // ler o arquivo JSON e armazenar os dados em um array de espécies
  data, err := ioutil.ReadFile("data/zoo_data.json")
  if err != nil {
    return nil, err
  }
  var speciesArray SpeciesArray
  json.Unmarshal(data, &speciesArray)

  // verificar se o array de ids é vazio
  if len(ids) == 0 {
    return []Species{}, nil
  }

  // criar o array de retorno da função
  var result SpeciesArray

  // adicionar as espécies correspondentes aos ids passados como parâmetro
  for _, id := range ids {
    found := false
    for _, species := range speciesArray {
      if species.Id == id {
        result = append(result, species)
        found = true
        break
      }
    }
    if !found {
      return nil, fmt.Errorf("species with id %d not found", id)
    }
  }

  return result, nil
}

func main() {
  // chamar a função getSpeciesByIds com diferentes conjuntos de ids
  species, err := getSpeciesByIds()
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("species with no id:", species)
  }

  species, err = getSpeciesByIds(1)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("species with id 1:", species)
  }

  species, err = getSpeciesByIds(1, 2, 3)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("species with ids 1, 2, and 3:", species)
  }

  species, err = getSpeciesByIds(4, 5, 6)
  if err != nil {
    fmt.Println(err)
  }
}