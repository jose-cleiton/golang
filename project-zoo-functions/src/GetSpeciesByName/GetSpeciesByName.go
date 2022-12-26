package main

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

type SpeciesArray []Species

// declaração da função pública
func GetSpeciesByName(name string) (SpeciesArray, error) {
  // ler o arquivo JSON e armazenar os dados em uma variável do tipo map[string]interface{}
  data, err := ioutil.ReadFile("data/zoo_data.json")
  if err != nil {
    return Species{}, err
  }
  var dataMap map[string]interface{}
  json.Unmarshal(data, &dataMap)

  // ler a lista de espécies do mapa de dados
  speciesList, ok := dataMap["species"].([]map[string]interface{})
  if !ok {
    return Species{}, fmt.Errorf("error parsing species data")
  }

  // procurar a espécie com o nome especificado
  for _, speciesData := range speciesList {
    nameStr, ok := speciesData["name"].(string)
    if !ok {
      return Species{}, fmt.Errorf("error parsing species name")
    }
    if nameStr == name {
      id, ok := speciesData["id"].(string)
      if !ok {
        return Species{}, fmt.Errorf("error parsing species id")
      }
      popularity, ok := speciesData["popularity"].(float64)
      if !ok {
        return Species{}, fmt.Errorf("error parsing species popularity")
      }
      location, ok := speciesData["location"].(string)
      if !ok {
        return Species{}, fmt.Errorf("error parsing species location")
      }
      availability, ok := speciesData["availability"].([]interface{})
      if !ok {
        return Species{}, fmt.Errorf("error parsing species availability")
      }
      availabilityStr := make([]string, len(availability))
      for i, a := range availability {
        availabilityStr[i], ok = a.(string)
        if !ok {
          return Species{}, fmt.Errorf("error parsing species availability element")
        }
      }
      residents, ok := speciesData["residents"].([]interface{})
      if !ok {
        return Species{}, fmt.Errorf("error parsing species residents")
      }
      residentsStruct := make([]struct {
        Name string `json:"name"`
        Sex  string `json:"sex"`
        Age  int    `json:"age"`
      }, len(residents))
      for i, r := range residents {
