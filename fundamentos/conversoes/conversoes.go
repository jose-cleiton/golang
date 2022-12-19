package main

import "fmt"

func main() {

	x:= 2.4
	y:= 2

	// não é possível converter int para float
	// fmt.Println(x / y)

	fmt.Println(x / float64(y))
 
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

}