package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// operation - структура для хранения бинарной операции
type operation struct {
	numSlice  []float64
	statement []string
	result    float64
}

type expression struct {
	strExpression string
}

// initOperation - функция иницицализации нового экземпляра operation
func initOperation(line *string, w *operation) {
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
}

func resultCalc(numOp string, w *operation) bool {
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
		if len(bufferSlice) == 2 {
			w.result = result
			fmt.Println(w.numSlice)
			fmt.Println(w.statement)
			fmt.Println(w.result)
			return true
		} else {
			w.numSlice = append(w.numSlice[:bufferIndex], result)
			if len(bufferSlice) > bufferIndex+2 {
				w.numSlice = append(w.numSlice, bufferSlice[bufferIndex+2:]...)
			}
			w.statement = append(w.statement[:bufferIndex], w.statement[bufferIndex+1:]...)
		}
	}
	return false
}

// calcFunc - расчет значения экземпляра operation
func calcFunc(w *operation) {
	statementStr := "^*:/+-"
	for _, value := range statementStr {
		if resultCalc(string(value), w) {
			break
		}
	}
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
	//strExp := "1*2+3*4+6*5+19/32+643*0.2213"
	//var exp1 operation
	//initOperation(&strExp, &exp1)
	//calcFunc(&exp1)
	//fmt.Println(exp1.result)
}
