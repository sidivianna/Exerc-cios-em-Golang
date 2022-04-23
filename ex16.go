package main

import (
	"fmt"
)

func main() {
	esportefavortito := "futebol"

	switch esportefavortito {
	case "futebol":
		fmt.Println("Brasil")

	case "voleibol":
		fmt.Println("EUA")

	case "judô":
		fmt.Println("Japão")
	}

}



// Crie um programa que utilize a declaração switch, onde o switch statemant seja uma  variável do tipo string e identificador "esportefavortito".