package operationObj

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// operationIn - структура для выполнения элементарных операций
type operationIn struct {
	numSlice   []float64
	statements []string
}

// clear - функция для очистки эксземпляра структуры operationIn
func clear(w *operationIn) {
	w.numSlice = []float64{}
	w.statements = []string{}

}

// copy дает возможность копировать элементы operationIn
func copy(w operationIn) operationIn {
	var newOp operationIn
	newOp.numSlice = w.numSlice
	newOp.statements = w.statements
	return newOp
}

// resultCalc, позволяет расчитать результат
// арифметического выражения для определенного
// бинарного оператора
func resultCalc(statement string, w *operationIn) (bool, float64) {
	var result float64
	var bufferSlice []float64
	var indexesStatement []int
	var bufferIndex int
	var firstNum, secondNum float64
	for index, value := range w.statements {
		if statement == value {
			indexesStatement = append(indexesStatement, index)
		}
	}
	for i := len(indexesStatement) - 1; i >= 0; i-- {
		bufferIndex = indexesStatement[i]
		firstNum = w.numSlice[bufferIndex]
		secondNum = w.numSlice[bufferIndex+1]

		result = binaryOperation(firstNum, secondNum, statement)

		bufferSlice = w.numSlice
		if len(w.numSlice) == 2 {
			return true, result
		} else {
			cutBinaryOperation(w, result, bufferIndex, &bufferSlice)
		}
	}
	return false, 0
}

// cutBinaryOperation позволяет вырезать отработанную бинарную операцию,
// то есть изъять операнды из numSlice и оператор из statements
func cutBinaryOperation(w *operationIn, result float64, bufferIndex int, bufferSlice *[]float64) {
	w.numSlice = append(w.numSlice[:bufferIndex], result)
	if len(*bufferSlice) > bufferIndex+2 {
		w.numSlice = append(w.numSlice, (*bufferSlice)[bufferIndex+2:]...)
	}
	w.statements = append(w.statements[:bufferIndex], w.statements[bufferIndex+1:]...)
}

// binaryOperation расчитывает результат бинарной операции
func binaryOperation(firstNum, secondNum float64, numOp string) float64 {
	var result float64
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
	return result
}

// InitOperation - функция иницицализации нового экземпляра operationIn
func InitOperation(line *string, w *operationIn) {
	clear(w)
	statements := "*+-^:/"
	var numberBuffer = ""
	var floatBuffer float64
	for index, value := range *line {
		if strings.ContainsRune(statements, value) {
			w.statements = append(w.statements, string(value))
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
}

// SeeAllInf - мето для проверки существующих данных
// внутри экзмпляра структуры operationIn
func (w operationIn) SeeAllInf() {
	fmt.Println("NumSlice:", w.numSlice)
	fmt.Println("Statements:", w.statements)
	//fmt.Println("Result:", w.result)
}

// CalcFunc - расчет значения operationIn
func (w operationIn) CalcFunc() float64 {
	statements := "^*:/+-"
	buffer := copy(w)
	for _, value := range statements {
		if ans, result := resultCalc(string(value), &buffer); ans == true {
			return result
		}
	}
	return 0 // Реализовать ошибку
}
