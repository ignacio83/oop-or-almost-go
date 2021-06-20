package main

import (
	"fmt"
	"ignacio83/oop-or-almost-go/ui"
	"log"
	"os"
)

const logPrefix = "oop-or-almost: "

func init() {
	log.SetPrefix(logPrefix)
	log.SetFlags(0)
}

func main() {
	shapeFromUserInputFactory := ui.ShapeFromUserInputFactory{Writer: os.Stdout, Reader: os.Stdin}
	shape, err := shapeFromUserInputFactory.Build()
	if err != nil {
		log.Fatal(err)
	}

	area := shape.Area()
	fmt.Printf("Area: %.2f\n", area)
}
