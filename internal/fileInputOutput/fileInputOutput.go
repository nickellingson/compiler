package fileInputOutput

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(fileName string) []byte {
	fmt.Println("Start of readfile")
	// Read contents
	dat, err := os.ReadFile(fileName)
	check(err)
	return dat
}