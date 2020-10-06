package main

import (
	"fmt"

	"github.com/marcocipri/game-pooling-engine/strcon"
)

func main() {
	fmt.Println("hello world")
	strcon.SwapCase("test")

	fmt.Println(strcon.SwapCase("test"))
	fmt.Println(strcon.SwapCase(strcon.Simple("")))

}
