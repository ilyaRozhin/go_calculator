package main

import (
	"fmt"
	"strings"
)

// bracketRules - соблюдение правил расстановки скобок
func bracketRules(brackets *string) string {
	buffer := "()"
	size := len(*brackets)
	*brackets = strings.ReplaceAll(*brackets, buffer, "")
	if size == len(*brackets) {
		return *brackets
	}
	return bracketRules(brackets)
}

// checkBracket - функция проверки скобок
func checkBracket(line *string) (bool, string) {
	openBracket := 0
	closeBracket := 0
	var bracketMass string
	for _, value := range *line {
		if value == '(' {
			openBracket++
			bracketMass += "("
		} else if value == ')' {
			closeBracket++
			bracketMass += ")"
		}
	}
	if openBracket != closeBracket {
		return false, "error: number of opening brackets is not equal to the number of closing brackets"
	} else if lastBracket := bracketRules(&bracketMass); lastBracket != "" {
		return false, "error: parentheses rules broken, correct errors, extra parentheses -> " + lastBracket
	}
	return true, "" // ???
}

// checkExample - функция проверки на соблюдения всех правил ввода мат. выражений
func checkExample(line *string) (bool, string) {
	if *line == "" {
		return false, "error: line is empty"
	} else if ansBool, ansString := checkBracket(line); ansBool == false {
		return false, ansString
	}
	return true, "" // ???
}

// menu - основное меню приложения
func menu() {
	var example string
	fmt.Println("Please enter a new example:")
	_, err := fmt.Scan(&example)
	if err != nil {
		fmt.Println("Scan error,please repeat input!")
	} else if ansBool, ansString := checkExample(&example); ansBool == false {
		fmt.Println(ansString)
	} else {
		fmt.Println("The entered string is valid:")
		fmt.Print(example)
	}
}

func main() {
	menu()
}
