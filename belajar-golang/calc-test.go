package main

import (
	"fmt"
	"belajar-golang/Calculation"
)

func main()  {
	resultSum := Calculation.Sum(1, 9)
	resultSubtract := Calculation.Subtract(11, 1)
	resultMultiply := Calculation.Multiply(5, 6)
	resultDivide := Calculation.Divide(20, 2)

	fmt.Println(resultSum, resultSubtract, resultMultiply, resultDivide)
}