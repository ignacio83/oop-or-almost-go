package ui

import (
	"bufio"
	"errors"
	"fmt"
	"ignacio83/oop-or-almost-go/geometry"
	"io"
	"strconv"
)

const chooseShapeMessage = "Choose the shape: (R)ectangle, (C)ircle or (T)riangle? "

var builderByFirstLetter = map[string]ShapeFromScannerBuilder{
	"R": RectangleFromScannerBuilder{},
	"C": CircleFromScannerBuilder{},
	"T": TriangleFromScannerBuilder{},
}

type ShapeFromUserInputFactory struct {
	Writer io.Writer
	Reader io.Reader
}

func (s *ShapeFromUserInputFactory) Build() (geometry.Shape, error) {
	fmt.Fprintf(s.Writer, chooseShapeMessage)
	scanner := bufio.NewScanner(s.Reader)
	scanner.Scan()
	shapeInput := scanner.Text()

	if len(shapeInput) == 0 {
		return nil, errors.New("shape is required")
	}
	builder := builderByFirstLetter[shapeInput]
	if builder == nil {
		return nil, errors.New("invalid shape")
	}
	return builder.build(s.Writer, scanner)
}

func requestRequiredFloat64Input(writer io.Writer, scanner *bufio.Scanner, label string) (float64, error) {
	fmt.Fprintf(writer, "%s: ", label)
	scanner.Scan()
	inputText := scanner.Text()

	if len(inputText) == 0 {
		return 0, fmt.Errorf("%s is required", label)
	}
	input, err := strconv.ParseFloat(inputText, 64)
	if err != nil {
		return 0, errors.New("width must be a float number")
	}
	return input, nil
}

type ShapeFromScannerBuilder interface {
	build(writer io.Writer, scanner *bufio.Scanner) (geometry.Shape, error)
}

type RectangleFromScannerBuilder struct {
}

func (r RectangleFromScannerBuilder) build(writer io.Writer, scanner *bufio.Scanner) (geometry.Shape, error) {
	width, err := requestRequiredFloat64Input(writer, scanner, "Width")
	if err != nil {
		return nil, err
	}
	height, err := requestRequiredFloat64Input(writer, scanner, "Height")
	if err != nil {
		return nil, err
	}
	return &geometry.Rectangle{Width: width, Height: height}, nil
}

type TriangleFromScannerBuilder struct {
}

func (r TriangleFromScannerBuilder) build(writer io.Writer, scanner *bufio.Scanner) (geometry.Shape, error) {
	base, err := requestRequiredFloat64Input(writer, scanner, "Base")
	if err != nil {
		return nil, err
	}
	height, err := requestRequiredFloat64Input(writer, scanner, "Height")
	if err != nil {
		return nil, err
	}
	return &geometry.Triangle{Base: base, Height: height}, nil
}

type CircleFromScannerBuilder struct {
}

func (r CircleFromScannerBuilder) build(writer io.Writer, scanner *bufio.Scanner) (geometry.Shape, error) {
	radius, err := requestRequiredFloat64Input(writer, scanner, "Radius")
	if err != nil {
		return nil, err
	}
	return &geometry.Circle{Radius: radius}, nil
}
