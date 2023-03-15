package main

import (
	"fmt"
	exp "go_calculator/src/expressionObj"
)

func main() {
	var new_exp exp.ExpressionObj
	exp.ScanUserInput(&new_exp)
	fmt.Println(new_exp.Calculate())
}
