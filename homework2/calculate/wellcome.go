package calculate

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseStr(str string) (dep int64, percent float64, time int64, err error) {
	str = strings.TrimSpace(str)
	regPattern := regexp.MustCompile("\\s+")
	str = regPattern.ReplaceAllString(str, " ")
	strNumbers := strings.Split(str, " ")
	if len(strNumbers) != 3 {
		return 0, 0, 0, errors.New("Too much numbers")
	}
	dep, err = strconv.ParseInt(strNumbers[0], 10, 64)
	if err != nil || dep <= 0 {
		return 0, 0, 0, errors.New("parse INTEGER")
	}
	percent, err = strconv.ParseFloat(strNumbers[1], 64)
	if err != nil || percent > 100 || percent <= 0 {
		return 0, 0, 0, errors.New("parse FLOAT")
	}
	time, err = strconv.ParseInt(strNumbers[2], 10, 64)
	if err != nil || time <= 0 {
		return 0, 0, 0, errors.New("parse INTEGER")
	}
	return dep, percent, time, nil
}

func getLine() (string, error) {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		return "", err
	}
	return in.Text(), nil
}

func Welcome() (int64, float64, int64) {
	var deposit int64
	var percent float64
	var time int64
	fmt.Printf("Enter 3 values: \n")
	for {
		fmt.Printf("The amount of the deposit (U_INTEGER) and annual percentage (REAL) and deposit term, years (U_INTEGER):")
		str, err := getLine()
		if err != nil || str == "" {
			fmt.Println("\u001B[31mError input value(scan)" + "[0m")
			continue
		}
		deposit, percent, time, err = parseStr(str)
		if err != nil {
			fmt.Println("\u001B[31mError input value(parser): " + err.Error() + "[0m")
			continue
		} else {
			break
		}
	}
	//fmt.Println(deposit, percent)
	return deposit, percent, time
}
