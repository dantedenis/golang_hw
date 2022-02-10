package biroot

import (
	"fmt"
	"math/cmplx"
	"os"
)

func discriminate(a, b, c float64) complex128 {
	return complex(b*b-4*a*c, 0)
}

func analyze(a, b, c float64) {
	if a == 0 {
		if b == 0 {
			fmt.Println("No roots")
			os.Exit(0)
		} else {
			if c == 0 {
				fmt.Println("Infinite number of root")
				os.Exit(0)
			}
			fmt.Println("Root: ", -c/b)
			os.Exit(0)
		}
	}
}

func calculate(a, b, c float64) (x1, x2 complex128) {
	analyze(a, b, c)
	D := discriminate(a, b, c)
	x1 = (complex(-b, 0) + cmplx.Sqrt(D)) / complex(2*a, 0)
	x2 = (complex(-b, 0) - cmplx.Sqrt(D)) / complex(2*a, 0)
	return
}

func readFile(filename string) (*complex128, *complex128, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error open file: " + err.Error())
		return nil, nil, err
	} else {
		fmt.Println("Success open file: " + filename)
	}
	defer file.Close()
	var coefs [3]float64
	for i := 0; i < 3; i++ {
		_, err := fmt.Fscanf(file, "%f", &coefs[i])
		if err != nil {
			fmt.Println("Error parse file: " + err.Error())
			return nil, nil, err
		}
	}
	//fmt.Printf("Your equation: %3f*x^2 + (%3f) + (%3f) = 0", coefs[0], coefs[1], coefs[2])
	var x1, x2 *complex128
	*x1, *x2 = calculate(coefs[0], coefs[1], coefs[2])
	return x1, x2, nil
}

func ReadArgs(str string) (x1, x2 *complex128, err error) {

	var a, b, c float64

	if str == "" {
		fmt.Printf("Enter the coefficients a,b,c: ")
		_, err := fmt.Scan(&a, &b, &c)
		if err != nil {
			fmt.Println("Error scan args")
			os.Exit(1)
		}
		fmt.Println(calculate(a, b, c))
	} else {
		x1, x2, err = readFile(str)
	}
	return x1, x2, err
}
