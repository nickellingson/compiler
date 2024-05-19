package lexicalAnalysis

import (
	"fmt"
	"strings"
)


func Tokenize (input []byte) []string {

	fmt.Println("Start of tokenize")

	// hash map for lexemes to token
	lexToToken := map[string]string{
		"print": "PRINT",
		"(":   "LAPREN",
		"\"": "QUOTE",
		")": "RPAREN",
	}

	index := 0
	var res strings.Builder
	var tokensList = []string{}

	// loop through all the bytes in file
	// find keywords, identifiers, seperators, operators, or literals
	for index < len(input) {
		c := input[index]

		token, ok := lexToToken[string(c)]

		// for literals need to append to string until next keyword, identifier, seperator, or operator comes by
		res.WriteString(string(c))

		// check lexeme from hashmap (byte)
		if ok {

			// literal 
			s := res.String()[:res.Len()-1]

			if len(s) > 0 {
				tokensList = append(tokensList, s)
			}
			tokensList = append(tokensList, token)
			res.Reset()
		}
		

		token2, ok2 := lexToToken[res.String()]

		// check lexeme from hashmap (string)
		if ok2 {
			tokensList = append(tokensList, token2)
			res.Reset()
		}
		index += 1
	}

	return tokensList
}