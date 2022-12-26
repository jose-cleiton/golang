package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

type ZooData struct {
	Species []Species
}

func main() {
	file, err := os.Open("data/zoo_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var data ZooData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
