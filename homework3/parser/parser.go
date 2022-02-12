package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func checkParse(str string) (string, error) {
	regPat, err := regexp.Compile("\\s+")
	if err != nil {
		return "", errors.New("\u001B[1;31mError compile regex\u001B[0m")
	}
	str = regPat.ReplaceAllString(str, "")
	regPat, err = regexp.Compile("^-?\\d+(/.\\d+)?([+-/*//]-?\\d+(/.\\d+)?){0,}$")
	if err != nil {
		return "", errors.New("\u001B[1;31mError compile regex\u001B[0m")
	}
	if !regPat.MatchString(str) {
		return "", errors.New("\u001B[1;31mNot a mathemical expresion\u001B[0m")
	}
	return str, nil
}

func getLine() (string, error) {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		return "", err
	}
	return in.Text(), nil
}

func Start() (string, bool) {
	fmt.Printf("Please enter your equation: ")
	str, err := getLine()
	if err != nil {
		fmt.Println(err.Error())
		return str, false
	}
	str, err = checkParse(str)
	if err != nil {
		fmt.Println(err.Error())
		return str, false
	}
	fmt.Println("Counting.....")
	return str, true
}
