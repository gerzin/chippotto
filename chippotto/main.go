package main

import (
	"flag"
	"fmt"

	"chippotto/interpreter"
)

func main() {
	var (
		romPath string
	)

	flag.StringVar(&romPath, "rom", "", "path of the rom file to run")
	flag.Parse()

	fmt.Printf("Loading rom from: %s\n", romPath)
	var interpreter = interpreter.NewInterpreter()
	interpreter.LoadRom(romPath)
}
