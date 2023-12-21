package logic

import (
	"math/rand"
	"strconv"
)

const (
	Fire 	= 0 // Fuego vence a la planta (planta + 1) % 3 = 0
	Water 	= 1 // Agua vence al fuego (fuego + 1) % 3 = 1
	Plant 	= 2 // Planta vence al agua (agua + 1) % 3 = 2
)

// Estructura para dar resultado de cada ronda
type RoundResult struct {
	Message				string `json:"message"`
	ComputerChoice 		string `json:"computer_choice"`
	RoundResult 		string `json:"round_result"`
	ComputerChoiceInt 	int	   `json:"computer_choice_int"`
	ComputerScore 		string `json:"computer_score"`
	PlayerScore 		string `json:"player_score"`
}

// Mensajes de resultado ganador player
var winMessages = []string{
	"¡Ganaste!",
	"¡Bien hecho!",
	"¡Sigue así!",
	"¡Eres el mejor!",
}

// Mensajes de resultado ganador computer
var loseMessages = []string{
	"¡Perdiste!",
	"¡Lo siento!",
	"¡Sigue intentando!",
	"¡No te rindas!",
}

// Mensajes de resultado empate
var tieMessages = []string{
	"¡Empate!",
	"¡Vaya!",
	"¡Vamos de nuevo!",
	"¡Suerte para la próxima!",
}

// Variable para almacenar el resultado de cada ronda
var ComputerScore, PlayerScore int

// Funcion para juagar cada Ronda
func PlayRound(playerValue int) RoundResult {
	// Generar numero random para la eleccion del computer
	computerValue := rand.Intn(3)

	var computerChoice, round string
	var computerChoiceInt int

	// Mensaje dependiendo de lo que elija el computer
	switch computerValue {
		case Fire:
			computerChoiceInt = Fire
			computerChoice = "La computadora eligió al monstruo de fuego"
		case Water:
			computerChoiceInt = Water
			computerChoice = "La computadora eligió al monstruo de agua"
		case Plant:
			computerChoiceInt = Plant
			computerChoice = "La computadora eligió al monstruo de planta"
	}

	// Generar un numero aleatorio de 0 a 3 para el mensaje de resultado
	messageInt := rand.Intn(4)

	// Declarar una variable para contener el mensaje de resultado
	var message string

	if playerValue == computerValue {
		// Empate
		message = tieMessages[messageInt]
		round = "Empate"
	} else if playerValue == (computerValue + 1) % 3 {
		// Ganador
		message = winMessages[messageInt]
		round = "Ganaste"
		PlayerScore++
	} else {
		// Perdedor
		message = loseMessages[messageInt]
		round = "Perdiste"
		ComputerScore++
	}

	// Retornar el resultado de la ronda
	return RoundResult{
		Message: message,
		ComputerChoice: computerChoice,
		RoundResult: round,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore: strconv.Itoa(ComputerScore),
		PlayerScore: strconv.Itoa(PlayerScore),
	}
}