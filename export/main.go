package main

import (
    "fmt"
    "funcoes"
    
   
)

func main() {

    // A função "apply" recebe um inteiro "n" e uma função "f" como argumentos
    // e retorna o resultado da aplicação de "f" a "n"
    fmt.Println(funcoes.Apply(2, funcoes.MultiplyByTwo))
    // Output: 4
   
}
