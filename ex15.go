package main

import (
	"fmt"
)

func main() {
	tamx := 0

	switch {
	case tamx == 0:
		fmt.Println("x top")
	case tamx == 1:
		fmt.Println("x médio")
	case tamx == 2:
		fmt.Println("x ruin")
	}
}



// Crie um programa que utilize a declaração switch, sem switch statement especifícado.