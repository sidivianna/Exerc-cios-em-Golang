package main

import (
	"fmt"
)

const (
	a = 2022 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(a, b, c, d, e)
}

