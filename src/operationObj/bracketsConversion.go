package operationObj

import (
	"fmt"
	"strings"
)

// calcFullExp - основная функция для расчета пользовательского выражения
func calcFullExp(str *string) string {
	*str = reduceBracketExp(str)
	var calcInput operationIn
	if statementIn(str) {
		InitOperation(str, &calcInput)
		return fmt.Sprint(calcInput.CalcFunc())
	}
	return *str
}

// reduceBracketExp - функция, позволяющая поэтапно избавиться от скобок в выражении
func reduceBracketExp(userLine *string) string {
	var operationStr string
	var inputFlag = false
	var bufferStr = *userLine
	var validCharacters = "0123456789.,*+-^:/"
	for _, value := range *userLine {
		if value == '(' {
			if inputFlag && operationStr != "" {
				operationStr = ""
			}
			inputFlag = true
		} else if value == ')' {
			if inputFlag {
				calcMini(&operationStr, &bufferStr, ")")
			}
			inputFlag = false
		} else {
			if inputFlag && (strings.ContainsRune(validCharacters, value)) {
				operationStr += string(value)
			}
		}
	}
	if bracketIn(&bufferStr) {
		return reduceBracketExp(&bufferStr)
	} else {
		return bufferStr
	}
}

// calcMini - функция для расчета элементарного случая
// со скобками и его замены в выражении на результат
func calcMini(operationStr *string, bufferStr *string, bracketVar string) {
	var buffer operationIn
	InitOperation(operationStr, &buffer)
	bufferResult := fmt.Sprint(buffer.CalcFunc())
	*operationStr = "(" + *operationStr + bracketVar
	*bufferStr = strings.Replace(*bufferStr, *operationStr, bufferResult, 1)
	*operationStr = ""
}

// bracketIn - проверка на наличие скобок в выражении
func bracketIn(str *string) bool {
	for _, value := range *str {
		if value == ')' || value == '(' {
			return true
		}
	}
	return false
}

// statementIn позволяет проверить выражение на наличие операторов
func statementIn(str *string) bool {
	statements := "*+-^:/"
	for _, value := range *str {
		if strings.ContainsRune(statements, value) {
			return true
		}
	}
	return false
}

// Calc - экспортируемая функция для расчета выражения
func Calc(str string) string {
	return calcFullExp(&str)
}
