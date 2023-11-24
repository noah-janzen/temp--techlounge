package main

import (
	"fmt"

	"github.com/thenativeweb/hallowelt/calculator"
)

func main() {
	sum := calculator.Add(21, 42)
	fmt.Println(sum)
}