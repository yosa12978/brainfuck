package interpreter

import (
	"fmt"
	"os"

	"github.com/yosa12978/brainfuck/internal/lexer"
	"github.com/yosa12978/brainfuck/internal/stack"
)

type Interpreter interface {
	Run()
}

const (
	memsize    int = 30000
	stackdepth int = 1000
)

type BFInterpreter struct {
	ip      int
	s       *stack.Stack
	memory  []uint8
	mp      int
	program []lexer.OpCode
}

func NewBFInterpreter(p []lexer.OpCode) Interpreter {
	intp := &BFInterpreter{
		ip:      0,
		s:       stack.NewStack(stackdepth),
		memory:  make([]uint8, memsize),
		mp:      0,
		program: p,
	}
	for i := 0; i < len(intp.memory); i++ {
		intp.memory[i] = 0
	}
	return intp
}

func (bfi *BFInterpreter) Run() {
	for bfi.ip < len(bfi.program) {
		instr := bfi.program[bfi.ip]
		bfi.ip++
		switch instr {

		case lexer.INCREMENT:
			bfi.memory[bfi.mp]++

		case lexer.DECREMENT:
			bfi.memory[bfi.mp]--

		case lexer.SHIFT_RIGHT:
			if bfi.mp == memsize-1 {
				bfi.mp = 0
			} else {
				bfi.mp++
			}

		case lexer.SHIFT_LEFT:
			if bfi.mp == 0 {
				bfi.mp = memsize - 1
			} else {
				bfi.mp--
			}

		case lexer.OUTPUT:
			fmt.Printf("%c", bfi.memory[bfi.mp])

		case lexer.INPUT:
			fmt.Scanf("%c", &bfi.memory[bfi.mp])

		case lexer.LOOP_START:
			close := bfi.findEnd(bfi.ip - 1)
			if close == -1 {
				fmt.Fprintf(os.Stderr, "syntax error")
				os.Exit(1)
			}
			brac := stack.NewBraces(bfi.ip-1, close)
			bfi.s.Push(brac)
			if bfi.memory[bfi.mp] == 0 {
				bfi.ip = bfi.s.Pop().End + 1
			}

		case lexer.LOOP_END:
			if !bfi.s.IsEmpty() {
				bfi.ip = bfi.s.Pop().Start
			}
		}
	}
}

func (bfi *BFInterpreter) findEnd(start_addr int) int {
	var count int = 0
	i := bfi.ip - 1
	for i < len(bfi.program) {
		if bfi.program[i] == lexer.LOOP_END {
			count--
		}
		if bfi.program[i] == lexer.LOOP_START {
			count++
		}
		if count == 0 {
			return i
		}
		i++
	}
	return -1
}
