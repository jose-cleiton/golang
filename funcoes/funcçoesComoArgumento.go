package main

import "fmt"

// A função "apply" recebe um inteiro "n" e uma função "f" como argumentos
// e retorna o resultado da aplicação de "f" a "n"
var apply = func(n int, f func(int) int) int {
    return f(n)
}

func main() {
    // Definindo uma função anônima que multiplica um número por 2
    multiplyByTwo := func(n int) int {
        return n * 2
    }

    // Usando a função "apply" para aplicar a função anônima ao número 3
    resultado := apply(3, multiplyByTwo)
    fmt.Println(resultado) // 6
}
