package main

import (
	"fmt"

	"github.com/nickellingson/compiler/internal/fileInputOutput"
	"github.com/nickellingson/compiler/internal/lexicalAnalysis"
)


func main () {
	fmt.Println("Start of main")

	// testFile, _ := filepath.Abs("../../precompiled/test1.txt")
	testFile := "../../precompiled/test1.txt"

	fmt.Println("Calling file input output")
	// Read file
	result := fileInputOutput.ReadFile(testFile)

	fmt.Println("Calling lexical analysis")
	// Lexical Analysis
	tokenList := lexicalAnalysis.Tokenize(result)

	fmt.Println(tokenList)
}