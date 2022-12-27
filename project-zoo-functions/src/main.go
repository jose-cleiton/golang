package main

import (

	"fmt"

	"github.com/jose-cleiton/data"

)


func main() {

	for _, species := range data.Data.Species {
    fmt.Println(species.Name)
}


	

	
}

	


    