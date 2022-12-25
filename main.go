package main

import "fmt"

func main() {
	aa := make(map[int]int)

	aa[0] = 0
	aa[1] = 1
	aa[3] = 3

	b := aa[2]

	fmt.Printf("%v\n", b)

}
