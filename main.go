package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

// checkNums - функция проверки корректности чисел
func checkNums(line *string) (bool, string) {
	numsSep := "*+-^:/)("
	errorMessage := "error: mistakes in writing numbers -> "
	var errFlag = false
	var resultLine string
	for _, value := range *line {
		if (value <= '9' && value >= '0') || value == '.' || value == ',' || (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') {
			resultLine += string(value)
		} else if strings.ContainsRune(numsSep, value) {
			if len(resultLine) != 0 {
				if _, err := strconv.Atoi(resultLine); err != nil {
					errorMessage += resultLine + " "
					errFlag = true
				}
				resultLine = ""
			}
		}
	}
	if errFlag {
		return false, errorMessage
	}
	return true, "Successful"
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
	var withoutSpace string
	for _, value := range *line {
		if value != ' ' {
			withoutSpace += string(value)
		}
	}
	*line = withoutSpace
	if *line == "" {
		return false, "error: line is empty"
	} else if bracketBool, bracketString := checkBracket(line); bracketBool == false {
		return false, bracketString
	} else if ansBool, ansString := checkNums(line); ansBool == false {
		return false, ansString
	}
	return true, "" // ???
}

// menu - основное меню приложения
func menu() {
	var example string
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a new example:")
	scanner.Scan()
	example = scanner.Text()
	if err := scanner.Err(); err != nil {
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
