MAIN_FILE = ./cmd/Brainfuck/main.go


run:
	go run $(MAIN_FILE)

build:
	go build -o ./bin/brainfuck.exe $(MAIN_FILE)