package main

import "fmt"

func main() {
	var a int
	ScanCount, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(ScanCount)
	h := a / 30
	m := a*2 - (h * 60)
	fmt.Println("it is", h, "hours", m, "minutes.")
}
