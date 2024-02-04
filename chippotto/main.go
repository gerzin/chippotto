package main

import (
	"chippotto/interpreter"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var (
		romPath       string
		disassemble   bool
		disOutputPath string
	)

	flag.StringVar(&romPath, "rom", "", "path of the rom file to run")
	flag.BoolVar(&disassemble, "disassemble", false, "disassemble a rom")
	flag.StringVar(&disOutputPath, "output", "./a.ch8.asm", "output path for the disassembled file")
	flag.Parse()

	fmt.Printf("Loading rom from: %s\n", romPath)

	if disassemble {
		const defaultOutputName string = "a.ch8.asm"

		fmt.Printf("Disassembling %s into %s\n", romPath, disOutputPath)
		asmCode := interpreter.DisassembleRom(romPath)
		dir, file := filepath.Split(disOutputPath)

		if file == "" {
			file = defaultOutputName
		}

		newPath := filepath.Join(dir, file)

		// Write the content to the file
		err := os.WriteFile(newPath, []byte(strings.Join(asmCode, "\n")), 0644)
		if err != nil {
			panic(fmt.Sprintf("Error saving the output to %s", newPath))
		}

		absPath, _ := filepath.Abs(newPath)
		fmt.Printf("Output written to: %s\n", absPath)

	} else {
		var chip8 = interpreter.NewInterpreter()
		chip8.LoadRom(romPath)
	}

}
