package main

import (
	"fmt"
	"sort"
)

func main() {
    numbers := []int{1, 3, 5, 7, 9, 11}

		fmt.Println(sort.SearchInts(numbers,7))

    var firstEven []int
		

    for _, number := range numbers {
        if number >= 7 {
            firstEven = append(firstEven, number)
            
        }
    }

    fmt.Println(firstEven) // expected output: 2
}
