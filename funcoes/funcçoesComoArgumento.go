package main

import (
	
	"fmt"
	

)

// A função "apply" recebe um inteiro "n" e uma função "f" como argumentos
// e retorna o resultado da aplicação de "f" a "n"
var apply = func(n int, f func(int) int) int {
    return f(n)
}

var multiplyByTwo = func(n int) int {
		return n * 2
}
func main() {
    // Usando a função "apply" para aplicar a função anônima ao número 3
 
    fmt.Println( apply(3, multiplyByTwo)) // 6
}
