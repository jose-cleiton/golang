package estruturas

import (
	"fmt"
	
)

type Pessoa struct {
	Nome string
	Idade int
}

func (p Pessoa) Falar() {
	fmt.Println("Olá, meu nome é", p.Nome, "e tenho", p.Idade, "anos de idade")

}