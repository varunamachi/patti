package main

import (
	"fmt"

	"github.com/varunamachi/patti/pt/dbg"
)

func main() {
	nums := []int{
		1,
		5,
		10,
		25,
		104,
		200,
		999,
		1000,
	}
	for _, n := range nums {
		val := dbg.ToStrNum(n)
		fmt.Printf("%s - %s\n", val.ID, val.Name)
	}
}
