package main

import (
	"fmt"
	"math"
	"/lib/strings"
)


func main() {
	x := math.Sum(1,2,3)
	y := strings.Join([]string{"a", "b", "c"}, ",")
	fmt.Println(x, y)
}
