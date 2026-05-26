package main

import "fmt"

func main() {
	numbers := []float32{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	numbers = append(numbers, 1,2,3,4,5,6,7,8,8)
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// initializing size of slice

	num := make([]int, 3,5)   // there 3 is the length , and 5 is the capacity
	fmt.Println(num)
	fmt.Println(cap(num))  // capacity of slice
	fmt.Println(len(num))

	num = append(num, 5,6)
	fmt.Println(cap(num)) 
	num = append(num,6)
	fmt.Println(cap(num)) 
	fmt.Println(num)

}