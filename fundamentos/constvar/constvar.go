package main

import (
	"fmt"
	"math"
)


func main() {
	const pi float64 = 3.14
	const pi2 = 3.1415
	var raio = 3.2

	var (
		nome = "João"
		idade = 32
	)

	const (
		a = 1
		b = 2
	)

	const e, f bool = true, false

	g,h,i := 2, false, "epa!"



	area:= pi * math.Pow(raio, 2)


fmt.Println("Nome:", nome, "Idade:", idade)

	println("Area da circunferencia é", area)
fmt.Println("Valor de a:", a, "Valor de b:", b)	
fmt.Println("Valor de e:", e, "Valor de f:", f)
fmt.Println("Valor de pi:", pi, "Valor de pi2:", pi2)
fmt.Println("Valor de g:", g, "Valor de h:", h, "Valor de i:", i)
}