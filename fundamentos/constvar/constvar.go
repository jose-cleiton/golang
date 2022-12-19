package main

import "math"


func main() {
	const pi float64 = 3.14
	const pi2 = 3.1415
	var raio = 3.2

	area:= pi * math.Pow(raio, 2)


	println("Area da circunferencia é", area)
}