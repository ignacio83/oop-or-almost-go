package main

import (
	"fmt"
	"ignacio83/oop-or-almost-go/ui"
	"log"
)

const logPrefix = "oop-or-almost: "

func init() {
	log.SetPrefix(logPrefix)
	log.SetFlags(0)
}

func main() {
	shape, err := ui.BuildShapeFromUserInputs()
	if err != nil {
		log.Fatal(err)
	}

	area := shape.Area()
	fmt.Printf("Area: %.2f\n", area)
}
