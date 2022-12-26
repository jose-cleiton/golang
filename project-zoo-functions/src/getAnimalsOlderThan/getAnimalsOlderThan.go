package Getanimalsolderthan

func GetAnimalsOlderThan(speciesName string, minAge int) bool {
  species, err := getSpeciesByIds(speciesName)
  if err != nil {
    return false
  }
  if len(species) == 0 {
    return false
  }
  for _, s := range species {
    for _, r := range s.Residents {
      if r.Age < minAge {
        return false
      }
    }
  }
  return true
}

