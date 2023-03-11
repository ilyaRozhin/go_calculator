package full_check

import (
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

// CheckBracket - функция проверки скобок
func CheckBracket(line *string) (bool, string) {
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
