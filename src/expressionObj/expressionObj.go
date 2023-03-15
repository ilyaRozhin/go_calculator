package expressionObj

import (
	"bufio"
	"fmt"
	"go_calculator/src/check"
	"go_calculator/src/operationObj"
	"os"
)

type ExpressionObj struct {
	str  string
	vars []variableObj
}

// withoutSpace позволяет избавиться от лишних пробелов в выражении
func withoutSpace(str *string) {
	var withoutSpace string
	for _, value := range *str {
		if value != ' ' {
			withoutSpace += string(value)
		}
	}
	*str = withoutSpace
}

// toString конвертирует пользовательскую строку в
// из LaTeX типа и других в обычную строку
func toString(str string) string {
	// Надо подумть как реализовать
	return str
}

// Calculate - основная функция для расчета пользовательского выражения
func (w ExpressionObj) Calculate() string {
	return operationObj.Calc(w.str)
}

// ScanUserInput ждем пользовательского ввода
func ScanUserInput(w *ExpressionObj) {
	var scan = bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a new example:")
	scan.Scan()
	w.str = scan.Text()
	withoutSpace(&(w.str))
	if err := scan.Err(); err != nil {
		fmt.Println("Scan error,please repeat input!")
	} else if ans, str := check.Valid(&(w.str)); ans == false { // Заменить потом str, на err
		fmt.Println(str)
	} else {
		fmt.Print("The entered string is valid: ")
		fmt.Println(w.str)

	}
}

// InitExpressionObj - инициализация экземпляра ExpressionObj,
// с помощью некоторой входной строки
func InitExpressionObj(w *ExpressionObj, newStr string) {
	newStr = toString(newStr)
	w.str = newStr
}

// ConvertToLatex - конвертирует строку содержащуюся в экзе-
// мпляре ExpressionObj в LaTeX форму
func ConvertToLatex(w *ExpressionObj) {
	// Надо подумать как реализовать
}

// AddVariables добавляет переменные к ExpressionObj
func AddVariables(varSlice []string) {
	// Инициализировать объекты типа variableObj
}