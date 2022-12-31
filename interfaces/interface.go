package main

import "fmt"

type Leitor interface {
	Ler() string
}


type Escritor interface {
	Escrever(texto string)

}

type Arquivo struct {
	nome   string
	conteudo string
}

func (a *Arquivo) Ler() string {
	return a.conteudo
}


func (a *Arquivo) Escrever(texto string) {
	a.conteudo = texto
	a.nome = texto
	return
}
func imprimirConteudo(l Leitor) {
	fmt.Println(l.Ler())
}


func main() {



	arquivo := &Arquivo{"arquivo.txt", "Conteudo do arquivo"}
	imprimirConteudo(arquivo)
}