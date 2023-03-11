package full_check

import (
	"strconv"
	"strings"
)

// CheckNum - функция проверки корректности чисел
func CheckNum(line *string) (bool, string) {
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
