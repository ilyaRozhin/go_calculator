package expression_in

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Expression - структура для выражения в скобках
type Expression struct {
	numSlice  []float64
	statement []string
}

// InitExpression - функция иницицализации нового экземпляра expression
func (w Expression) InitExpression(line *string) Expression {
	numSep := "*+-^:/"
	var numberBuffer = ""
	var floatBuffer float64
	for index, value := range *line {
		if strings.ContainsRune(numSep, value) {
			w.statement = append(w.statement, string(value))
			floatBuffer, _ = strconv.ParseFloat(numberBuffer, 64) // Error strconv
			w.numSlice = append(w.numSlice, floatBuffer)
			numberBuffer = ""
		} else if index == (len(*line) - 1) {
			numberBuffer += string(value)
			floatBuffer, _ = strconv.ParseFloat(numberBuffer, 64)
			w.numSlice = append(w.numSlice, floatBuffer)
		} else if value >= '0' && value <= '9' || value == ',' || value == '.' {
			numberBuffer += string(value)
		}
	}
	return w
}

func (w Expression) SeeAllInf() {
	fmt.Println("NumSlice:", w.numSlice)
	fmt.Println("Statements:", w.statement)
	//fmt.Println("Result:", w.result)
}

func copy(w Expression) Expression {
	var newCopy Expression
	newCopy.numSlice = w.numSlice
	newCopy.statement = w.statement
	//newCopy.result = w.result
	return newCopy
}

// CalcFunc - расчет значения экземпляра operation
func (w Expression) CalcFunc() float64 {
	statementStr := "^*:/+-"
	buffer := copy(w)
	for _, value := range statementStr {
		if ans, result := resultCalc(string(value), &buffer); ans == true {
			return result
		}
	}
	return 0 // Реализовать ошибку
}

func resultCalc(numOp string, w *Expression) (bool, float64) {
	var result float64
	var bufferSlice []float64
	var indexesStatement []int
	var bufferIndex int
	var firstNum, secondNum float64
	for index, value := range w.statement {
		if numOp == value {
			indexesStatement = append(indexesStatement, index)
		}
	}
	for i := len(indexesStatement) - 1; i >= 0; i-- {
		bufferIndex = indexesStatement[i]
		firstNum = w.numSlice[bufferIndex]
		secondNum = w.numSlice[bufferIndex+1]
		switch numOp {
		case "+":
			result = firstNum + secondNum
		case "-":
			result = firstNum - secondNum
		case "*":
			result = firstNum * secondNum
		case "/", ":":
			result = firstNum / secondNum
		case "^":
			result = math.Pow(firstNum, secondNum)
		default:
			result = firstNum * secondNum
		}
		bufferSlice = w.numSlice
		if len(w.numSlice) == 2 {
			return true, result
		} else {
			w.numSlice = append(w.numSlice[:bufferIndex], result)
			if len(bufferSlice) > bufferIndex+2 {
				w.numSlice = append(w.numSlice, bufferSlice[bufferIndex+2:]...)
			}
			w.statement = append(w.statement[:bufferIndex], w.statement[bufferIndex+1:]...)
		}
	}
	return false, 0
}
