package main

import (
	"fmt"

	"github.com/nickellingson/compiler/internal/lexicalAnalysis"
)


func main () {
	fmt.Println("start of main")
	lexicalAnalysis.Tokenize()
}