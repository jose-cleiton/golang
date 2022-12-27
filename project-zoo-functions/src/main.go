package main

import "fmt"

// Defina a estrutura Animal com os campos Name e Age
type Animal struct {
    Name string
    Age  int
}

func main() {
    // Crie uma lista de animais de exemplo
    animals := []Animal{
        {Name: "Lion", Age: 10},
        {Name: "Tiger", Age: 5},
        {Name: "Bear", Age: 7},
        {Name: "Gorilla", Age: 3},
    }

    // Imprima a lista de animais
    for _, a := range animals {
        fmt.Println(a.Name, a.Age)
    }
}
