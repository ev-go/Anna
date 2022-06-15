package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	h := a / 30
	m := a*2 - (h * 60)
	fmt.Println("it is", h, "hours", m, "minutes.")
}
