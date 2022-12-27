package main

import (

	"fmt"

	"github.com/jose-cleiton/data"
	


)


func main() {
	zoo:=data.Data
	
	allLionsOlderThan12 := data.GetAnimalsOlderThan(zoo, "lions", 12)


	allTigersOlderThan5 := data.GetAnimalsOlderThan(zoo, "tigers", 5)

	

	fmt.Println(allLionsOlderThan12)


	fmt.Println(allTigersOlderThan5)

	fmt.Println(data.GetSpeciesByIds(zoo, "0938aa23-f153-4937-9f88-4858b24d6bce"))
}


	

	


	


    