package main

import (
	"fmt"

	"github.com/nickellingson/compiler/internal/fileInputOutput"
	"github.com/nickellingson/compiler/internal/lexicalAnalysis"
)


func main () {
	fmt.Println("start of main")

	// testFile, _ := filepath.Abs("../../precompiled/test1.txt")
	testFile := "../../precompiled/test1.txt"

	// Read file
	fileInputOutput.ReadFile(testFile)

	// Lexical Analysis
	lexicalAnalysis.Tokenize()
}