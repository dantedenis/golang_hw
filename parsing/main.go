package main

import "fmt"

func def1() {
	fmt.Println("defer1")
}

func def2() {
	fmt.Println("defer2")
}

func test() {
	defer def1()
	defer def2()
	fmt.Println("kek")
}

func main() {
	test()
}
