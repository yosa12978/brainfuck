package main

import (
	"fmt"
	"os"

	"github.com/yosa12978/brainfuck/internal/interpreter"
	"github.com/yosa12978/brainfuck/internal/lexer"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "error parsing arguments")
		os.Exit(1)
	}

	program_file := os.Args[1]
	program_bytes, err := os.ReadFile(program_file)
	if err != nil {
		panic(err)
	}
	program_runes := []rune(string(program_bytes))
	program := lexer.Tokenize(program_runes)
	interpreter := interpreter.NewBFInterpreter(program)
	interpreter.Run()
}
