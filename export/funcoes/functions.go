package funcoes
    


// A função "apply" recebe um inteiro "n" e uma função "f" como argumentos
// e retorna o resultado da aplicação de "f" a "n"
func Apply(n int, f func(int) int) int {
    return f(n)
}

// A função "multiplyByTwo" multiplica um número por 2
func MultiplyByTwo(n int) int {
    return n * 2
}
