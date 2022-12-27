package main

import (

	"fmt"

	"github.com/jose-cleiton/data"
	


)


func main() {
	allLionsOlderThan12 := data.GetAnimalsOlderThan(data.Data, "lions", 12)


	allTigersOlderThan5 := data.GetAnimalsOlderThan(data.Data, "tigers", 5)

	

	fmt.Println(allLionsOlderThan12)


	fmt.Println(allTigersOlderThan5)
}


	

	


	


    