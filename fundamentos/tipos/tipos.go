package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// tipos de variáveis
	 // números inteiros
	 fmt.Println(1, 2, 1000)
	 fmt.Println("Literal inteiro é", 32000)
	 fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	 // sem sinal (só positivos)... uint8 uint16 uint32 uint64
	 var b byte = 255
	 fmt.Println("O byte é", b)
	 fmt.Println("O byte é", reflect.TypeOf(b))
	 // com sinal... int8 int16 int32 int64
	 i1 := math.MaxInt64
	 fmt.Println("O valor máximo do int é", i1)
	 fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))
	 var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	 fmt.Println("O rune é", i2)
	 fmt.Println("O tipo de i2 é", reflect.TypeOf(i2))
	 // números reais (float32, float64)
	 var x float32 = 49.99
	 fmt.Println("O tipo de x é", x)
	 fmt.Println("O tipo de x é", reflect.TypeOf(x))
	 fmt.Println("O tipo literal de 49.99 é ",reflect.TypeOf(49.99))
	 // boolean
	 bo := true
	 fmt.Println("O tipo de bo é", bo)
	 fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	 // string
	 s1 := "Olá meu nome é João"
	 fmt.Println(s1 + "!")
	 s2 := `Olá
	 meu
	 nome
	 é
	 João`
	 fmt.Println("O tamanho da string s2 é", len(s2))
	 // char???
	 // var x char = 'b'
	 char := 'a'
	 fmt.Println("O tipo de char é", char)
	 fmt.Println("O tipo de char é", reflect.TypeOf(char))

	 // string com múltiplas linhas
	 s3 := `Olá
	 meu
	 nome
	 é
	 João`
	 fmt.Println(s3)

	 
	   



	}