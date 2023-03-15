package check

import (
	"strconv"
	"strings"
)

// bracketRules - соблюдение правил расстановки скобок
func bracketRules(brackets *string) string {
	rule := "()"
	size := len(*brackets)
	*brackets = strings.ReplaceAll(*brackets, rule, "")
	if size == len(*brackets) {
		return *brackets
	}
	return bracketRules(brackets)
}

// checkBracket - функция проверки скобок
func checkBracket(line *string) (bool, string) {
	openBracket := 0
	closeBracket := 0
	var bracketLine string
	for _, value := range *line {
		if value == '(' {
			openBracket++
			bracketLine += "("
		} else if value == ')' {
			closeBracket++
			bracketLine += ")"
		}
	}
	if openBracket != closeBracket {
		return false, "error: number of opening brackets is not equal to the number of closing brackets"
	} else if lastBrackets := bracketRules(&bracketLine); lastBrackets != "" {
		return false, "error: parentheses rules broken, correct errors, extra parentheses -> " + lastBrackets
	}
	return true, "Successful" // ???
}

// checkNum - функция проверки корректности чисел
func checkNum(line *string) (bool, string) {
	statements := "*+-^:/)("
	errMessage := "error: mistakes in writing numbers -> "
	var errFlag = false
	var resLine string
	for _, value := range *line {
		if (value <= '9' && value >= '0') || value == '.' || value == ',' {
			resLine += string(value)
		} else if strings.ContainsRune(statements, value) {
			if len(resLine) != 0 {
				if _, err := strconv.ParseFloat(resLine, 64); err != nil {
					errMessage += resLine + " "
					errFlag = true
				}
				resLine = ""
			}
		}
	}
	if errFlag {
		return false, errMessage
	}
	return true, "Successful"
}

// Valid - функция проверки на соблюдения всех правил ввода мат. выражений
func Valid(str *string) (bool, string) {

	if *str == "" {
		return false, "error: line is empty"
	} else if ansBr, strBr := checkBracket(str); ansBr == false {
		return false, strBr
	} else if ansNum, strNum := checkNum(str); ansNum == false {
		return false, strNum
	}

	return true, "Successful"
}
