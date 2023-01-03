package main

import (
	"fmt"
	"math/math"
    "funcoes"
    "strings"
)

func main() {

   
    x := math.Sum(1,2,3)
    fmt.Println(x)
    a:=funcoes.Apply(2, funcoes.MultiplyByTwo)
    fmt.Println(a)
    j:= strings.Join([]string{"a", "b", "c"}, ",")
    fmt.Println(j)
    
   
}
