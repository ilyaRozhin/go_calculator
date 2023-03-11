package main

import (
	"bufio"
	"fmt"
	Expr "go_calculator/src/expression_in"
	Check "go_calculator/src/full_check"
	"os"
)

// menu - основное меню приложения
func menu() {
	var example string
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a new example:")
	scanner.Scan()
	example = scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error,please repeat input!")
	} else if ansBool, ansString := Check.CheckExample(&example); ansBool == false {
		fmt.Println(ansString)
	} else {
		fmt.Println("The entered string is valid:")
		fmt.Print(example)
	}
}

func main() {
	//menu()
	//strExp := "1*2+3*4+6*5+19/32+643*0.1212"
	strExp := "1*2+3*4+6*5"
	var exp1 Expr.Expression
	exp1 = exp1.InitExpression(&strExp)
	result := exp1.CalcFunc()
	fmt.Println(result)
	exp1.SeeAllInf()
}
