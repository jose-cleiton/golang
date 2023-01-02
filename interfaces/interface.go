package main

import "fmt"

// type Leitor interface {
// 	Ler() string
// 	Imprimir() 
// }


type Escritor interface {
	Escrever(texto string)

}

type Livro struct {
	autor   string
	titulo  string
	paginas []string
}


type Arquivo struct {
	nome   string
	conteudo string
}


func (l *Livro) Escrever(texto string) {
	l.paginas = append(l.paginas, texto)
}


func (self *Arquivo) Ler() string {
	return self.conteudo
}

func (self *Arquivo) Imprimir() {
	fmt.Println(self.nome)
	fmt.Println(self.conteudo)

}


func (a *Arquivo) Escrever(texto string) {
	a.conteudo = texto
	a.nome = texto
	return
}





func main() {

	var arquivo = &Arquivo{"arquivo.txt", "Conteudo do arquivo"}
  arquivo.Imprimir()

}