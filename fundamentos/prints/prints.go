package main

import "fmt"

func main() {

	
	fmt.Print("Mesma")
	fmt.Println(" linha.")
	fmt.Println(" Nova linha.")

	// imprimindo caracteres numericos com strings
	// e concatenando com variaveis
	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
}


