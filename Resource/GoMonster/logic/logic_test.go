package logic

import (
	"testing"
	"fmt"
)

func TestPlayRound(t *testing.T){
	for i := 0; i < 3; i++ {
		ronda := PlayRound(i)

		switch i {
		case 0:
			fmt.Println("Fuego")
		case 1:
			fmt.Println("Agua")
		case 2:
			fmt.Println("Planta")
		}

		fmt.Printf("Mensaje: %s\n", ronda.Message)
		fmt.Printf("ComputerChoice: %s\n", ronda.ComputerChoice)
		fmt.Printf("RoundResult: %s\n", ronda.RoundResult)
		fmt.Printf("ComputerChoiceInt: %d\n", ronda.ComputerChoiceInt)
		fmt.Printf("ComputerScore: %s\n", ronda.ComputerScore)
		fmt.Printf("PlayerScore: %s\n", ronda.PlayerScore)

		fmt.Println("--------------------------------------------------")
	}
}