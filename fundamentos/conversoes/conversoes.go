package main

import (
	"fmt"
	"strconv"
)

func main() {

	x:= 2.4
	y:= 2

	// não é possível converter int para float
	// fmt.Println(x / y)

	fmt.Println(x / float64(y))
 
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado
	fmt.Println("Teste " + string(97))

	fmt.Println("Teste " + string(rune(97)))
	// int para string  correto
	fmt.Println("Teste " + strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	

	

}