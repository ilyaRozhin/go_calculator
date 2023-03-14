package bracket_expression

import (
	"fmt"
	exp "go_calculator/src/expression_in"
	"strings"
)

// CalcFullExp - основная функция для расчета пользовательского выражения
func CalcFullExp(userInput *string) string {
	*userInput = reduceBracketExp(userInput)
	var calcInput exp.Expression
	calcInput = exp.InitExpression(userInput, &calcInput)
	return fmt.Sprint(calcInput.CalcFunc())
}

// reduceBracketExp - функция, позволяющая поэтапно избавляться от скобок в выражении
func reduceBracketExp(userLine *string) string {
	var initStr string
	var inputFlag = false
	var bufferStr = *userLine
	var statements = "*+-^:/"
	for _, value := range *userLine {
		if value == '(' {
			if inputFlag && initStr != "" {
				//calcMini(&initStr, &bufferStr, "(")
				initStr = ""
			}
			inputFlag = true
		} else if value == ')' {
			if inputFlag {
				calcMini(&initStr, &bufferStr, ")")
			}
			inputFlag = false
		} else {
			if inputFlag && (value <= '9' && value >= '0' || value == '.' || value == ',') {
				initStr += string(value)
			} else if strings.ContainsRune(statements, value) && inputFlag {
				initStr += string(value)
			}
		}
	}
	//fmt.Println(bufferStr)
	if bracketIn(&bufferStr) {
		return reduceBracketExp(&bufferStr)
	} else {
		return bufferStr
	}
}

// calcMini - функция для расчета элементарного случая
// со скобками и его замены в выражении на результат
func calcMini(initStr *string, bufferStr *string, bracketVar string) {
	var buffer exp.Expression
	buffer = exp.InitExpression(initStr, &buffer)
	bufferResult := buffer.CalcFunc()
	*initStr = "(" + *initStr + bracketVar
	*bufferStr = strings.Replace(*bufferStr, *initStr, fmt.Sprint(bufferResult), 1)
	*initStr = ""
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
