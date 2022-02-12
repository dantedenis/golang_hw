package executer

import (
	"errors"
	"fmt"
	"homework3/stack"
	"os"
	"strconv"
	"strings"
)

func isOperator(ch uint8) bool {
	return strings.ContainsAny(string(ch), "+-*/")
}

func isOperand(ch uint8) bool {
	return ch >= '0' && ch <= '9' || ch == '.'
}

func getWeight(str string) int {
	switch str {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return -1
}

func getPriority(str1 string, str2 string) bool {
	return getWeight(str1) >= getWeight(str2)
}

func toPostfix(str string) []string {
	isUnar := true
	result := ""
	st := stack.New()
	length := len(str)
	for i := 0; i < length; i++ {
		ch := string(str[i])
		if !isOperator(str[i]) || isUnar {
			j := i
			number := ""
			if isUnar && str[i] == '-' {
				j++
				number += "-"
			}
			for ; j < length && isOperand(str[j]); j++ {
				number += string(str[j])
			}
			result += " " + number
			i = j - 1
			isUnar = false
		} else {
			for !st.IsEmpty() {
				top := st.Top().(string)
				if !getPriority(top, ch) {
					break
				}
				result += " " + top
				st.Pop()
			}
			st.Push(ch)
			isUnar = true
		}
	}
	for !st.IsEmpty() {
		str := st.Pop().(string)
		result += " " + str
	}
	result = strings.TrimSpace(result)
	return strings.Split(result, " ")
}

func operation(op string, a, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("\u001B[1;31mDivizion by zero\u001B[0m")
		} else {
			return a / b, nil
		}
	}
	return 0, errors.New("\u001B[1;31mEmpty...\u001B[0m")
}

func calculate(lines []string) (interface{}, error) {
	st := stack.New()
	for _, line := range lines {
		if isOperand(line[0]) || line[0] == '-' && len(line) != 1 {
			st.Push(line)
		} else {
			a, err := strconv.ParseFloat(st.Pop().(string), 64)
			if err != nil {
				return nil, errors.New("\u001B[1;31mError Parse string to float\u001B[0m")
			}
			b, err := strconv.ParseFloat(st.Pop().(string), 64)
			if err != nil {
				return nil, errors.New("\u001B[1;31mError Parse string to float\u001B[0m")
			}
			result, err := operation(line, b, a)
			if err != nil {
				return nil, err
			}
			st.Push(strconv.FormatFloat(result, 'f', 6, 64))
		}
	}
	//st.PrintAll()
	return st.Pop(), nil
}

func Exec(str string) {
	result := toPostfix(str)
	//fmt.Println(result, len(result))
	answer, err := calculate(result)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("\u001B[1;36mAnswer: \n%s = %s\n\u001B[0m", str, answer)
}
