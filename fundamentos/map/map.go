package main

import "fmt"

type Person struct {
  Name string
  Sex  string
  Age  int
}





func main() {
	

	data := []struct {
		Name string
		Sex  string
		Age  int
	}{
		{"Faustino", "male", 7},
		{"Maria", "female", 9},
		{"João", "male", 10},
		{"Ana", "female", 12},
		{"Pedro", "male", 14},
	}

	people := map[string]Person{
		"Faustino": {Name: "Faustino", Sex: "male", Age: 7},
		"Maria":    {Name: "Maria", Sex: "female", Age: 9},
		"João":     {Name: "João", Sex: "male", Age: 10},
		"Ana":      {Name: "Ana", Sex: "female", Age: 12},
		"Pedro":    {Name: "Pedro", Sex: "male", Age: 14},
	}
	

	fmt.Println("data:", data)
	fmt.Println("people:", people)


}
