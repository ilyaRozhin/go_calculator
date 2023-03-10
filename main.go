package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// operation - структура для хранения бинарной операции
type operation struct {
	firstNum  float64
	secondNum float64
	operation string
}

// initOperation - функция иницицализация нового экземпляра operation
func initOperation(line *string, w *operation) {
	numSep := "*+-^:/"
	var numberBuffer = ""
	for _, value := range *line {
		if strings.ContainsRune(numSep, value) {
			w.operation = string(value)
			w.firstNum, _ = strconv.ParseFloat(numberBuffer, 64) // Error strconv
			numberBuffer = ""
		} else if value == ')' {
			w.secondNum, _ = strconv.ParseFloat(numberBuffer, 64)
		} else if value >= '0' && value <= '9' || value == ',' || value == '.' {
			numberBuffer += string(value)
		}
	}
}

// calcFunc - расчет значения экземпляра operation
func (w operation) calcFunc() float64 {
	var result float64
	switch w.operation {
	case "+":
		result = w.firstNum + w.secondNum
	case "-":
		result = w.firstNum - w.secondNum
	case "*":
		result = w.firstNum * w.secondNum
	case "/":
		result = w.firstNum / w.secondNum
	default:
		result = w.firstNum * w.secondNum
	}
	return result
}

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

// checkNum - функция проверки корректности чисел
func checkNum(line *string) (bool, string) {
	numSep := "*+-^:/)("
	errorMessage := "error: mistakes in writing numbers -> "
	var errFlag = false
	var resultLine string
	for _, value := range *line {
		if (value <= '9' && value >= '0') || value == '.' || value == ',' { //|| (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') {
			resultLine += string(value)
		} else if strings.ContainsRune(numSep, value) {
			if len(resultLine) != 0 {
				if _, err := strconv.ParseFloat(resultLine, 64); err != nil {
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
	} else if ansBool, ansString := checkNum(line); ansBool == false {
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
