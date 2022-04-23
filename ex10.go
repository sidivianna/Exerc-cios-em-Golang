package main

import (
	"fmt"
)

func main() {
	anon := 1991
	anof := 2022
	for anon <= anof {
		fmt.Println(anon)
		anon++
	}
}

// Loop utilizando a sentaxe: for condition {}; utilizando-a para demonstrar os anos desde que voçê nasceu.