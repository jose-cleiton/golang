package main



import "fmt"

var sum = func(a,b int)int {
	return a + b
}

var numbers = []int{}

var frutas = []string{"Banan", "Pera", "Uva", "Abacati"}


func main() {

	numbers = append(numbers, 1,2,3,4,5,6,7,8,9,10)

	for indice, item := range frutas {
		fmt.Println(indice, item)

		for l, letras:= range item{
			fmt.Println(l, letras)
		}
	}

	

}