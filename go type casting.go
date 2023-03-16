package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Go type Casting/ Convertion/ Assertion")

	var number1 int = 55
	fmt.Printf(" Number1 Value = %v, Type = %T\n", number1, number1)

	var number2 float64 = float64(number1)
	fmt.Printf(" Number2 Value = %v, Type = %T\n", number2, number2)

	var number3 int = int(number2)
	fmt.Printf(" Number3 float64 to integer -> Value = %v, Type = %T\n", number3, number3)

	var number4 string = string(number3)
	fmt.Printf(" Number4 int to string, ASCII value -> Value = %v, Type = %T", number4, number4)

	var number6 string = strconv.Itoa(number3)
	fmt.Printf(" Number6*(Atoi -> string to int) Value = %v, Type = %T", number6, number6)

}
