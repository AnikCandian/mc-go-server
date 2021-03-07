package main

import (
	"positions"
	"fmt"
)

func main() {
	fmt.Println(decodePositions(encodePositions(1, 2, 3)))
}
