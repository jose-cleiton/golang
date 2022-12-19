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
	fmt.Printf("O valor de x é %.2f", x)

// imprimir variáveis como no c
	a := 1
	b := 1.9999
	c := false
	d := "opa!"
	fmt.Printf("%d %f %.1f %t %s", a, b, b, c, d)
}


