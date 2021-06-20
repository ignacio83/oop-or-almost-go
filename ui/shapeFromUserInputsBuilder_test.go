package ui

import (
	"bytes"
	"ignacio83/oop-or-almost-go/geometry"
	"testing"
)

func TestBuildRectangle(t *testing.T) {
	reader := bytes.Buffer{}
	writer := bytes.Buffer{}

	reader.Write([]byte("R\n10\n10\n"))

	shapeFromUserInputsBuilder := ShapeFromUserInputFactory{Writer: &writer, Reader: &reader}
	shapeGot, _ := shapeFromUserInputsBuilder.Build()

	got := writer.String()
	want := "Choose the shape: (R)ectangle, (C)ircle or (T)riangle? Width: Height: "

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if !isRectangle(shapeGot) {
		t.Errorf("got %q want Rectangle", shapeGot)
	}
}

func TestBuildTriangle(t *testing.T) {
	reader := bytes.Buffer{}
	writer := bytes.Buffer{}

	reader.Write([]byte("T\n10\n20\n"))

	shapeFromUserInputsBuilder := ShapeFromUserInputFactory{Writer: &writer, Reader: &reader}
	shapeGot, _ := shapeFromUserInputsBuilder.Build()

	got := writer.String()
	want := "Choose the shape: (R)ectangle, (C)ircle or (T)riangle? Base: Height: "

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if !isTriangle(shapeGot) {
		t.Errorf("got %q want Triangle", shapeGot)
	}
}

func TestBuildCircle(t *testing.T) {
	reader := bytes.Buffer{}
	writer := bytes.Buffer{}

	reader.Write([]byte("C\n5\n"))

	shapeFromUserInputsBuilder := ShapeFromUserInputFactory{Writer: &writer, Reader: &reader}
	shapeGot, _ := shapeFromUserInputsBuilder.Build()

	got := writer.String()
	want := "Choose the shape: (R)ectangle, (C)ircle or (T)riangle? Radius: "

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if !isCircle(shapeGot) {
		t.Errorf("got %q want Circle", shapeGot)
	}
}

func isRectangle(shape interface{}) bool {
	switch shape.(type) {
	case *geometry.Rectangle:
		return true
	default:
		return false
	}
}

func isTriangle(shape interface{}) bool {
	switch shape.(type) {
	case *geometry.Triangle:
		return true
	default:
		return false
	}
}

func isCircle(shape interface{}) bool {
	switch shape.(type) {
	case *geometry.Circle:
		return true
	default:
		return false
	}
}
