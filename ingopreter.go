package main

import (
	"errors"
	"facelzz/ingopreter/lexer"
	"log"
	"os"
	"strings"
)

func main() {
	// get file name argument
	filePath, err := getFilePath()
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	verbose := getArg("-v")

	// read file into the memory?
	lines, err := readFileByLines(filePath, verbose)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	// lexer: tokenization, line by line
	for _, line := range lines {
		tokens := lexer.Tokenize([]rune(line), verbose)
		if tokens != nil {
			log.Println(tokens)
		}
	}

	// parser: build a syntax tree

	// semantic analysis

	// execution
}

func getArg(agrName string) bool {
	for _, arg := range os.Args[1:] {
		if arg == agrName {
			return true
		}
	}

	return false
}

func getFilePath() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("usage: ingopreter <file>")
	}
	filePath := os.Args[1]

	return filePath, nil
}

func readFileByLines(path string, verbose bool) ([]string, error) {
	if verbose {
		log.Println("File to read:", path)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if verbose {
		log.Println("Bytes:", bytes)
	}
	lines := strings.SplitAfter(string(bytes), "\n")
	if verbose {
		log.Println("Lines:", lines)
	}

	return lines, nil
}
