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

func readFile(file string) {

}

func ReadArgs(str string) {

	var a, b, c float64

	if str == "" {
		fmt.Printf("Enter the coefficients a,b,c: ")
		_, err := fmt.Scan(&a, &b, &c)
		if err != nil {
			fmt.Println("Error scan args")
			os.Exit(1)
		}
		calculate(a, b, c)
	} else {
		readFile(str)
	}
}
