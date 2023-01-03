package main

import "fmt"

func divisao(a, b int) (quociente, resto int) {
    quociente = a / b
    resto = a % b
    return
}

func main() {
    quociente, resto := divisao(10, 3)
    fmt.Println(quociente, resto) // 3 1
}



