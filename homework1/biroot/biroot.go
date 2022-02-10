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
	fmt.Printf("Your equation: \u001B[1;33;44m(%.3f * x²) + (%.3f * x) + (%.3f) = 0\033[0m \n", a, b, c)
	analyze(a, b, c)
	D := discriminate(a, b, c)
	x1 = (complex(-b, 0) + cmplx.Sqrt(D)) / complex(2*a, 0)
	x2 = (complex(-b, 0) - cmplx.Sqrt(D)) / complex(2*a, 0)
	return
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("\u001B[31mError open file: " + err.Error() + "[0m")
	} else {
		fmt.Println("\u001B[3;32mSuccess open file: " + filename + "[0m")
	}
	defer file.Close()
	var coefs [3]float64
	for i := 0; i < 3; i++ {
		_, err := fmt.Fscanf(file, "%f", &coefs[i])
		if err != nil {
			fmt.Println("\u001B[31mError parse file: " + err.Error() + "[0m")
			return
		}
	}
	x1, x2 := calculate(coefs[0], coefs[1], coefs[2])
	fmt.Printf("\u001B[32mx1 = %f\nx2 = %f\n\u001B[0m", x1, x2)
}

func ReadArgs(str string) {

	var a, b, c float64

	if str == "" {
		fmt.Printf("Enter the coefficients a,b,c: ")
		_, err := fmt.Scan(&a, &b, &c)
		if err != nil {
			fmt.Println("\u001B[31mError scan args" + "[0m")
			os.Exit(1)
		}
		x1, x2 := calculate(a, b, c)
		fmt.Printf("\u001B[32mx1 = %f\nx2 = %f\n\u001B[0m", x1, x2)
	} else {
		readFile(str)
	}
}
