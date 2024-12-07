package main

import (
	"fmt"
	"os"

	"github.com/StephenButtolph/canoto"
)

const (
	scalarsFile            = "./internal/proto/examples/scalars.go"
	largestFieldNumberFile = "./internal/proto/examples/largest_field_number.go"
)

var files = []string{
	scalarsFile,
	largestFieldNumberFile,
}

func main() {
	for _, file := range files {
		if err := canoto.Generate(file); err != nil {
			fmt.Println("Failed to generate canoto:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Successfully generated canoto")
}
