package calculate

import "fmt"

func withoutCapitalize(rub int64, percent float64, time int64) {
	result := float64(rub) * percent * float64(time) / 100
	fmt.Printf("Result after %d years \u001B[4;34mwithout\u001B[0m capitalize: \u001B[37m%d\n\u001B[0m", time, rub+int64(result))
}

func withCapitalize(rub int64, percent float64, time int64) {
	for i := 0; i < int(time); i++ {
		result := float64(rub) * percent / 100
		rub += int64(result)
	}
	fmt.Printf("Result after %d years \u001B[4;34mwith\u001B[0m capitalize: \u001B[37m%d\n\u001B[0m", time, rub)
}

func Calculate(rub int64, percent float64, time int64) {
	withoutCapitalize(rub, percent, time)
	withCapitalize(rub, percent, time)
}
