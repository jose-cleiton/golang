package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int){

	for i := 0; i < qtde; i++ {

		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}


func main() {


	go fale("Maria", "Pq vc não fala comigo?", 3)
	go fale("João", "Só posso falar depois de vc!", 1)

	fale("Pedro", "Entendi!!!", 10)
}