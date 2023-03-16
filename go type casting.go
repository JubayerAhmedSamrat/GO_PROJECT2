package main

import "fmt"

func main() {


	fmt.Println("Go type casting")
	
	var number int = 55
	fmt.Printf("Value = %v, Type = %T\n", number, number)

	var number2 float64 = float64(number)
	fmt.Printf("Value = %v, Type = %T\n", number2, number2) 



}