package lexer

type OpCode uint8

const (
	INCREMENT OpCode = iota
	DECREMENT
	SHIFT_LEFT
	SHIFT_RIGHT
	INPUT
	OUTPUT
	LOOP_START
	LOOP_END
)

func Tokenize(program []rune) []OpCode {
	bytecode := []OpCode{}
	for _, v := range program {
		switch v {
		case '.':
			bytecode = append(bytecode, OUTPUT)
		case ',':
			bytecode = append(bytecode, INPUT)
		case '+':
			bytecode = append(bytecode, INCREMENT)
		case '-':
			bytecode = append(bytecode, DECREMENT)
		case '>':
			bytecode = append(bytecode, SHIFT_RIGHT)
		case '<':
			bytecode = append(bytecode, SHIFT_LEFT)
		case '[':
			bytecode = append(bytecode, LOOP_START)
		case ']':
			bytecode = append(bytecode, LOOP_END)
		}
	}
	return bytecode
}
