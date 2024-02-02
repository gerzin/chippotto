package main

import (
	"flag"
	"fmt"

	"github.com/gerzin/chippotto/interpreter"
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
