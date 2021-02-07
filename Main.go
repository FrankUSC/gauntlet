package main

import (
	"fmt"
	"strconv"
)

var asdf int = 2

func main() {
	var i int = 3
	j := strconv.Itoa(i)
	fmt.Println(j)
	fmt.Printf("%v, %T !\n", j, j)
}
