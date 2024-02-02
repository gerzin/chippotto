package interpreter

import (
	"fmt"
	"os"
)

const GfxSize uint16 = 0x144
const ramSize uint = 0x1000

var fontSet = []uint8{
	0xF0, 0x90, 0x90, 0x90, 0xF0,
	0x20, 0x60, 0x20, 0x20, 0x70,
	0xF0, 0x10, 0xF0, 0x80, 0xF0,
	0xF0, 0x10, 0xF0, 0x10, 0xF0,
	0x90, 0x90, 0xF0, 0x10, 0x10,
	0xF0, 0x80, 0xF0, 0x10, 0xF0,
	0xF0, 0x80, 0xF0, 0x90, 0xF0,
	0xF0, 0x10, 0x20, 0x40, 0x40,
	0xF0, 0x90, 0xF0, 0x90, 0xF0,
	0xF0, 0x90, 0xF0, 0x10, 0xF0,
	0xF0, 0x90, 0xF0, 0x90, 0x90,
	0xE0, 0x90, 0xE0, 0x90, 0xE0,
	0xF0, 0x80, 0x80, 0x80, 0xF0,
	0xE0, 0x90, 0x90, 0x90, 0xE0,
	0xF0, 0x80, 0xF0, 0x80, 0xF0,
	0xF0, 0x80, 0xF0, 0x80, 0x80,
}

type Interpreter struct {
	gfx        [GfxSize]uint8
	ram        [ramSize]uint8
	delayTimer uint8
	soundTimer uint8
	pc         uint16
	drawScreen bool
	vx         [16]uint8
	key        [16]uint8
	stack      [16]uint16
	sp         uint16
	iv         uint16
}

func NewInterpreter() Interpreter {
	interpreter := Interpreter{
		pc:         0x200,
		drawScreen: true,
	}

	copy(interpreter.ram[:], fontSet)

	return interpreter
}

func (p *Interpreter) LoadRom(path string) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		return
	}

	if len(fileBytes) > int(ramSize)-0x200 {
		fmt.Fprintln(os.Stderr, "Rom file too big: ", len(fileBytes), " bytes")
	}
	copy(p.ram[0x200:], fileBytes)
}

func (p *Interpreter) Step() {
	opcode := (uint16(p.ram[p.pc]) << 8) | uint16(p.ram[p.pc+1])

	switch opcode & 0xF000 {
	case 0x0000: // CLR
		for i := range p.gfx {
			p.gfx[i] = 0
		}
		p.drawScreen = true
		p.pc += 2

	default:
		fmt.Printf("Invalid opcode %X\n", opcode)
	}

	if p.delayTimer > 0 {
		p.delayTimer--
	}

	if p.soundTimer > 0 {
		p.soundTimer--
	}
}

func (p *Interpreter) Draw(callback func([GfxSize]uint8)) {
	if p.drawScreen {
		callback(p.gfx)
		p.drawScreen = false
	}

}

func (p Interpreter) Beep(callback func()) {
	if p.soundTimer == 1 {
		callback()
	}
}
