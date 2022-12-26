package zoo

type Resident struct {
  Name string
  Sex  string
  Age  int
}

type Species struct {
  ID          string
  Name        string
  Popularity  int
  Location    string
  Availability []string
  Residents   []Resident
}

type Employee struct {
  ID          string
  FirstName   string
  LastName    string
  Managers    []string
  ResponsibleFor []string
}

type Hours struct {
  Open  int
  Close int
}

type Prices struct {
  Adult  float64
  Senior float64
  Child  float64
}

type Zoo struct {
  Species    []Species
  Employees  []Employee
  Hours      map[string]Hours
  Prices     Prices
}

func GetSpeciesByIds(zoo Zoo, ids ...string) []Species {
  // Declare uma variável para armazenar as espécies encontradas
  var species []Species

  // Percorra a lista de espécies do zoológico
  for _, s := range zoo.Species {
    // Verifique se o ID da espécie atual está presente na lista de IDs fornecidos
    for _, id := range ids {
      if s.ID == id {
        // Se estiver, adicione a espécie ao array de espécies encontradas
        species = append(species, s)
      }
    }
  }

  // Retorne o array de espécies encontradas
  return species
}
