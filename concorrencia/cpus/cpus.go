package main

import "fmt"

type Array []int

func (a Array) Map(callback func(int) int) Array {
	if a == nil {
		return nil
	}

	result := make(Array, len(a))
	for i, v := range a {
		result[i] = callback(v)
	}
	return result
}

func main() {
	array := Array{1, 2, 3, 4, 5}
	
	
	fmt.Println(array.Map(func(a,b int) int {
			return a+b
		}
	)) // [2 4 6 8 10]
}
