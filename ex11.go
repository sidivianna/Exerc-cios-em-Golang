package main

import (
	"fmt"
)

func main() {
	anon := 1991
	anof := 2022
	
	for {
		if anon > anof {
			break
		}
		fmt.Println(anon)
		anon++
	}
}

//Loop utilizando a sintaxe: for {}; para demonstrar os anos desde que voçê nasceu.