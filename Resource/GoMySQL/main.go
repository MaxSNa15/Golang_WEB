package main

import (
	"GoMySQL/data"
)

func main() {
	data.Connect()
	data.Close()
}