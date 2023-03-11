package full_check

// CheckExample - функция проверки на соблюдения всех правил ввода мат. выражений
func CheckExample(line *string) (bool, string) {
	var withoutSpace string
	for _, value := range *line {
		if value != ' ' {
			withoutSpace += string(value)
		}
	}
	*line = withoutSpace
	if *line == "" {
		return false, "error: line is empty"
	} else if bracketBool, bracketString := CheckBracket(line); bracketBool == false {
		return false, bracketString
	} else if ansBool, ansString := CheckNum(line); ansBool == false {
		return false, ansString
	}
	return true, "" // ???
}
