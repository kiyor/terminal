package main

import (
	"github.com/kiyor/terminal"
	"github.com/kiyor/terminal/color"
)

func main() {
	terminal.Stdout.Color("y").
		Print("Hello world").Nl().
		Reset().
		Colorf("@{kW}Hello world\n")

	color.Print("@rHello world")
	color.Print("@{216 8}Hello world")
}
