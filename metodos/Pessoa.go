package main

import (
	"fmt"
	"math"
	"time"
)

const horasPorAno = 8766

type Pessoa struct {
	Nome       string
	Nascimento time.Time
	Senha      string
}

func (c Pessoa) mostrarNome() {
	fmt.Println("Nome:", c.Nome)
	idadeEmHoras := math.Round(time.Now().Sub(c.Nascimento).Hours())
	idadeEmAnos := idadeEmHoras / horasPorAno
	fmt.Printf("Idade: %.0f anos\n", idadeEmAnos)
}

func criarPessoa(nome string, nascimento time.Time, senha string) Pessoa {
	return Pessoa{nome, nascimento, senha}
}


func main() {
	pessoa1 := criarPessoa("Joao", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "minhasenha")

	
	pessoa1.mostrarNome()
}
