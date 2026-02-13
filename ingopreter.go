package main

import (
	"errors"
	"facelzz/ingopreter/internal/lexer"
	"log/slog"
	"os"
	"strings"
)

func main() {
	setupLogging()

	// get file name argument
	filePath, err := getFilePath()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// read file into the memory?
	lines, err := readFileByLines(filePath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// lexer: tokenization, line by line
	for _, line := range lines {
		tokens := lexer.Tokenize([]rune(line))
		if tokens != nil {
			slog.Info("line tokenized",
				"tokens", tokens)
		}
	}

	// parser: build a syntax tree

	// semantic analysis

	// execution
}

func setupLogging() {
	logLevel := slog.LevelInfo
	verbose := getArg("-v")
	if verbose {
		logLevel = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)
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

func readFileByLines(path string) ([]string, error) {
	slog.Debug("reading file",
		"path", path)

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.SplitAfter(string(bytes), "\n")
	slog.Debug("reading file",
		"bytes", bytes,
		"lines", lines)

	return lines, nil
}
