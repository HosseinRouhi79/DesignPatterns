package main

import (
	"designPatterns/Patterns/bridge"
)

func main() {

	e := &bridge.Epson{}
	mac := &bridge.Mac{}
	mac.SetPrinter(e)
	mac.Print()
}
