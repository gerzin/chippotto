package interpreter

import (
	"encoding/binary"
	"fmt"
	"os"
)

func Disassemble(instruction uint16) string {
	unknown := func(instr uint16) string {
		return fmt.Sprintf("NOOP;; UNKNOWN INSTRUCTION %X", instr)
	}

	switch instruction & 0xF000 {
	case 0x0:
		switch instruction & 0x000F {
		case 0x0:
			return "CLS"
		case 0xE:
			return "RET"
		default:
			nnn := instruction & 0x0FFF
			return fmt.Sprintf("SYS %X", nnn)
		}
	case 0x1000:
		nnn := instruction & 0x0FFF
		return fmt.Sprintf("JMP %X", nnn)
	case 0x2000:
		nnn := instruction & 0x0FFF
		return fmt.Sprintf("CALL %X", nnn)
	case 0x3000:
		V := (instruction & 0x0F00) >> 8
		KK := (instruction & 0x00FF)
		return fmt.Sprintf("SE V%X, %X", V, KK)
	case 0x4000:
		V := (instruction & 0x0F00) >> 8
		KK := (instruction & 0x00FF)
		return fmt.Sprintf("SNE V%X, %X", V, KK)
	case 0x5000:
		VX := (instruction & 0x0F00) >> 8
		VY := (instruction & 0x00F0) >> 4
		return fmt.Sprintf("SE V%X, V%X", VX, VY)
	case 0x6000:
		V := (instruction & 0x0F00) >> 8
		KK := instruction & 0x00FF
		return fmt.Sprintf("LD V%X, %X", V, KK)
	case 0x7000:
		V := (instruction & 0x0F00) >> 8
		KK := instruction & 0x00FF
		return fmt.Sprintf("ADD V%X, %X", V, KK)
	case 0x8000:
		VX := (instruction & 0x0F00) >> 8
		VY := (instruction & 0x00F0) >> 4
		i := uint8(instruction & 0x000F)
		switch i {
		case 0x0:
			return fmt.Sprintf("LD V%X, V%X", VX, VY)
		case 0x1:
			return fmt.Sprintf("OR V%X, V%X", VX, VY)
		case 0x2:
			return fmt.Sprintf("AND V%X, V%X", VX, VY)
		case 0x3:
			return fmt.Sprintf("XOR V%X, V%X", VX, VY)
		case 0x4:
			return fmt.Sprintf("ADD V%X, V%X", VX, VY)
		case 0x5:
			return fmt.Sprintf("SUB V%X, V%X", VX, VY)
		case 0x6:
			return fmt.Sprintf("SHR V%X {, V%X}", VX, VY)
		case 0x7:
			return fmt.Sprintf("SUBN V%X, V%X", VX, VY)
		case 0xE:
			return fmt.Sprintf("SHL V%X {, V%X}", VX, VY)
		default:
			return unknown(instruction)
		}
	case 0x9000:
		VX := (instruction & 0x0F00) >> 8
		VY := (instruction & 0x00F0) >> 4
		return fmt.Sprintf("SNE V%X, V%X", VX, VY)
	case 0xA000:
		nnn := instruction & 0x0FFF
		return fmt.Sprintf("LD I, %X", nnn)
	case 0xB000:
		nnn := instruction & 0x0FFF
		return fmt.Sprintf("JP V0, %X", nnn)
	case 0xC000:
		V := (instruction & 0x0F00) >> 8
		KK := instruction & 0x00FF
		return fmt.Sprintf("RND V%X, %X", V, KK)
	case 0xD000:
		X := (instruction & 0x0F00) >> 8
		Y := (instruction & 0x00F0) >> 4
		N := instruction & 0x000F
		return fmt.Sprintf("DRW V%X, V%X, %X", X, Y, N)
	case 0xE000:
		switch instruction & 0x00FF {
		case 0x9E:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("SKP V%X", x)
		case 0xA1:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("SKNP V%X", x)
		default:
			return unknown(instruction)
		}
	case 0xF000:
		switch instruction & 0x00FF {
		case 0x07:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("LD V%X, DT", x)
		case 0x0A:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("LD V%X, K", x)
		case 0x15:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("LD DT, V%X", x)
		case 0x18:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("LD ST, V%X", x)
		case 0x1E:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("ADD I, V%X", x)
		case 0x29:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("LD F, V%X", x)
		case 0x33:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("LD B, V%X", x)
		case 0x55:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("LD [I], V%X", x)
		case 0x65:
			x := (instruction % 0x0F00) >> 8
			return fmt.Sprintf("LD V%X, [I]", x)
		default:
			return unknown(instruction)
		}
	default:
		return unknown(instruction)
	}
}

func DisassembleRom(filePath string) []string {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	resultArray := make([]string, len(fileBytes)/2)

	for i := 0; i < len(fileBytes)-1; i += 2 {
		// Extract a uint16 value from the current pair of bytes
		value := binary.LittleEndian.Uint16(fileBytes[i : i+2])

		// Apply your function to the uint16 value
		result := Disassemble(value)

		// Store the result in the new array
		resultArray[i/2] = result
	}

	return resultArray

}
