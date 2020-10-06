package main

import (
	"fmt"

	"github.com/marcocipri/game-pooling-engine/strcon"
)

// Hello : here you tell us what Salutation is
// Printer : what is this?
// Greet : describe what this function does
// CreateMessage : describe what this function does
func Hello() string {
	fmt.Println("hello world")
	strcon.SwapCase("test")

	fmt.Println(strcon.SwapCase("test"))
	fmt.Println(strcon.SwapCase(strcon.Simple("")))

	return "Hello, world."
}
