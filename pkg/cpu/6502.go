package cpu

import (
	"nes-emulator/pkg/memory"
	"strconv"
)

type CPU6502 struct {
	bus     memory.Bus
	lookup  [256]Instruction
	a       uint8
	x       uint8
	y       uint8
	stkp    uint8
	pc      uint16
	status  uint8
	fetched uint8
	addrAbs uint16
	addrRel uint16
	opcode  uint8
	cycles  uint8
}

type Instruction struct {
	Name     string
	AddrName string
	Operate  func() uint8
	AddrMode func() uint8
	Cycles   uint8
}

type Flags6502 uint8

const (
	// Status register flags
	C Flags6502 = 1 << iota // Carry
	Z                       // Zero
	I                       // Interrupt Disable
	D                       // Decimal Mode (unused in NES)
	B                       // Break
	U                       // Unused
	V                       // Overflow
	N                       // Negative
)

func NewCPU6502(bus memory.Bus) *CPU6502 {
	cpu := &CPU6502{
		bus:     bus,
		a:       0x00,
		x:       0x00,
		y:       0x00,
		stkp:    0x00,
		pc:      0x00,
		status:  0x00,
		fetched: 0x00,
		addrAbs: 0x00,
		addrRel: 0x00,
		opcode:  0x00,
		cycles:  0,
	}
	cpu.lookup = [256]Instruction{
		{"BRK", "IMP", (*cpu).BRK, (*cpu).IMP, 7}, {"ORA", "IZX", (*cpu).ORA, (*cpu).IZX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"???", "IMP", (*cpu).NOP, (*cpu).IMP, 3}, {"ORA", "ZP0", (*cpu).ORA, (*cpu).ZP0, 3}, {"ASL", "ZP0", (*cpu).ASL, (*cpu).ZP0, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"PHP", "IMP", (*cpu).PHP, (*cpu).IMP, 3}, {"ORA", "IMM", (*cpu).ORA, (*cpu).IMM, 2}, {"ASL", "IMP", (*cpu).ASL, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"ORA", "ABS", (*cpu).ORA, (*cpu).ABS, 4}, {"ASL", "ABS", (*cpu).ASL, (*cpu).ABS, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6},
		{"BPL", "REL", (*cpu).BPL, (*cpu).REL, 0}, {"ORA", "IZY", (*cpu).ORA, (*cpu).IZY, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"ORA", "ZPX", (*cpu).ORA, (*cpu).ZPX, 4}, {"ASL", "ZPX", (*cpu).ASL, (*cpu).ZPX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"CLC", "IMP", (*cpu).CLC, (*cpu).IMP, 2}, {"ORA", "ABY", (*cpu).ORA, (*cpu).ABY, 4}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"ORA", "ABX", (*cpu).ORA, (*cpu).ABX, 4}, {"ASL", "ABX", (*cpu).ASL, (*cpu).ABX, 7}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7},
		{"JSR", "ABS", (*cpu).JSR, (*cpu).ABS, 6}, {"AND", "IZX", (*cpu).AND, (*cpu).IZX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"BIT", "ZP0", (*cpu).BIT, (*cpu).ZP0, 3}, {"AND", "ZP0", (*cpu).AND, (*cpu).ZP0, 3}, {"ROL", "ZP0", (*cpu).ROL, (*cpu).ZP0, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"PLP", "IMP", (*cpu).PLP, (*cpu).IMP, 4}, {"AND", "IMM", (*cpu).AND, (*cpu).IMM, 2}, {"ROL", "IMP", (*cpu).ROL, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"BIT", "ABS", (*cpu).BIT, (*cpu).ABS, 4}, {"AND", "ABS", (*cpu).AND, (*cpu).ABS, 4}, {"ROL", "ABS", (*cpu).ROL, (*cpu).ABS, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6},
		{"BMI", "REL", (*cpu).BMI, (*cpu).REL, 6}, {"AND", "IZY", (*cpu).AND, (*cpu).IZY, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"AND", "ZPX", (*cpu).AND, (*cpu).ZPX, 4}, {"ROL", "ZPX", (*cpu).ROL, (*cpu).ZPX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"SEC", "IMP", (*cpu).SEC, (*cpu).IMP, 2}, {"AND", "ABY", (*cpu).AND, (*cpu).ABY, 4}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"AND", "ABX", (*cpu).AND, (*cpu).ABX, 4}, {"ROL", "ABX", (*cpu).ROL, (*cpu).ABX, 7}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7},
		{"RTI", "IMP", (*cpu).RTI, (*cpu).IMP, 6}, {"EOR", "IZX", (*cpu).EOR, (*cpu).IZX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 3}, {"EOR", "ZP0", (*cpu).EOR, (*cpu).ZP0, 3}, {"LSR", "ZP0", (*cpu).LSR, (*cpu).ZP0, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"PHA", "IMP", (*cpu).PHA, (*cpu).IMP, 3}, {"EOR", "IMM", (*cpu).EOR, (*cpu).IMM, 2}, {"LSR", "IMP", (*cpu).LSR, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"JMP", "ABS", (*cpu).JMP, (*cpu).ABS, 3}, {"EOR", "ABS", (*cpu).EOR, (*cpu).ABS, 4}, {"LSR", "ABS", (*cpu).LSR, (*cpu).ABS, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6},
		{"BVC", "REL", (*cpu).BVC, (*cpu).REL, 2}, {"EOR", "IZY", (*cpu).EOR, (*cpu).IZY, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"EOR", "ZPX", (*cpu).EOR, (*cpu).ZPX, 4}, {"LSR", "ZPX", (*cpu).LSR, (*cpu).ZPX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"CLI", "IMP", (*cpu).CLI, (*cpu).IMP, 2}, {"EOR", "ABY", (*cpu).EOR, (*cpu).ABY, 4}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"EOR", "ABX", (*cpu).EOR, (*cpu).ABX, 4}, {"LSR", "ABX", (*cpu).LSR, (*cpu).ABX, 7}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7},
		{"RTS", "IMP", (*cpu).RTS, (*cpu).IMP, 6}, {"ADC", "IZX", (*cpu).ADC, (*cpu).IZX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 3}, {"ADC", "ZP0", (*cpu).ADC, (*cpu).ZP0, 3}, {"ROR", "ZP0", (*cpu).ROR, (*cpu).ZP0, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"PLA", "IMP", (*cpu).PLA, (*cpu).IMP, 4}, {"ADC", "IMM", (*cpu).ADC, (*cpu).IMM, 2}, {"ROR", "IMP", (*cpu).ROR, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"JMP", "IND", (*cpu).JMP, (*cpu).IND, 5}, {"ADC", "ABS", (*cpu).ADC, (*cpu).ABS, 4}, {"ROR", "ABS", (*cpu).ROR, (*cpu).ABS, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6},
		{"BVS", "REL", (*cpu).BVS, (*cpu).REL, 2}, {"ADC", "IZY", (*cpu).ADC, (*cpu).IZY, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"ADC", "ZPX", (*cpu).ADC, (*cpu).ZPX, 4}, {"ROR", "ZPX", (*cpu).ROR, (*cpu).ZPX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"SEI", "IMP", (*cpu).SEI, (*cpu).IMP, 2}, {"ADC", "ABY", (*cpu).ADC, (*cpu).ABY, 4}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"ADC", "ABX", (*cpu).ADC, (*cpu).ABX, 4}, {"ROR", "ABX", (*cpu).ROR, (*cpu).ABX, 7}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7},
		{"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"STA", "IZX", (*cpu).STA, (*cpu).IZX, 6}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"STY", "ZP0", (*cpu).STY, (*cpu).ZP0, 3}, {"STA", "ZP0", (*cpu).STA, (*cpu).ZP0, 3}, {"STX", "ZP0", (*cpu).STX, (*cpu).ZP0, 3}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 3}, {"DEY", "IMP", (*cpu).DEY, (*cpu).IMP, 2}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"TXA", "IMP", (*cpu).TXA, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"STY", "ABS", (*cpu).STY, (*cpu).ABS, 4}, {"STA", "ABS", (*cpu).STA, (*cpu).ABS, 4}, {"STX", "ABS", (*cpu).STX, (*cpu).ABS, 4}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 4},
		{"BCC", "REL", (*cpu).BCC, (*cpu).REL, 2}, {"STA", "IZY", (*cpu).STA, (*cpu).IZY, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"STY", "ZPX", (*cpu).STY, (*cpu).ZPX, 4}, {"STA", "ZPX", (*cpu).STA, (*cpu).ZPX, 4}, {"STX", "ZPY", (*cpu).STX, (*cpu).ZPY, 4}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 4}, {"TYA", "IMP", (*cpu).TYA, (*cpu).IMP, 2}, {"STA", "ABY", (*cpu).STA, (*cpu).ABY, 5}, {"TXS", "IMP", (*cpu).TXS, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 5}, {"STA", "ABX", (*cpu).STA, (*cpu).ABX, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5},
		{"LDY", "IMM", (*cpu).LDY, (*cpu).IMM, 2}, {"LDA", "IZX", (*cpu).LDA, (*cpu).IZX, 6}, {"LDX", "IMM", (*cpu).LDX, (*cpu).IMM, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"LDY", "ZP0", (*cpu).LDY, (*cpu).ZP0, 3}, {"LDA", "ZP0", (*cpu).LDA, (*cpu).ZP0, 3}, {"LDX", "ZP0", (*cpu).LDX, (*cpu).ZP0, 3}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 3}, {"TAY", "IMP", (*cpu).TAY, (*cpu).IMP, 2}, {"LDA", "IMM", (*cpu).LDA, (*cpu).IMM, 2}, {"TAX", "IMP", (*cpu).TAX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"LDY", "ABS", (*cpu).LDY, (*cpu).ABS, 4}, {"LDA", "ABS", (*cpu).LDA, (*cpu).ABS, 4}, {"LDX", "ABS", (*cpu).LDX, (*cpu).ABS, 4}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 4},
		{"BCS", "REL", (*cpu).BCS, (*cpu).REL, 2}, {"LDA", "IZY", (*cpu).LDA, (*cpu).IZY, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"LDY", "ZPX", (*cpu).LDY, (*cpu).ZPX, 4}, {"LDA", "ZPX", (*cpu).LDA, (*cpu).ZPX, 4}, {"LDX", "ZPY", (*cpu).LDX, (*cpu).ZPY, 4}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 4}, {"CLV", "IMP", (*cpu).CLV, (*cpu).IMP, 2}, {"LDA", "ABY", (*cpu).LDA, (*cpu).ABY, 4}, {"TSX", "IMP", (*cpu).TSX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 4}, {"LDY", "ABX", (*cpu).LDY, (*cpu).ABX, 4}, {"LDA", "ABX", (*cpu).LDA, (*cpu).ABX, 4}, {"LDX", "ABY", (*cpu).LDX, (*cpu).ABY, 4}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 4},
		{"CPY", "IMM", (*cpu).CPY, (*cpu).IMM, 2}, {"CMP", "IZX", (*cpu).CMP, (*cpu).IZX, 6}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"CPY", "ZP0", (*cpu).CPY, (*cpu).ZP0, 3}, {"CMP", "ZP0", (*cpu).CMP, (*cpu).ZP0, 3}, {"DEC", "ZP0", (*cpu).DEC, (*cpu).ZP0, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"INY", "IMP", (*cpu).INY, (*cpu).IMP, 2}, {"CMP", "IMM", (*cpu).CMP, (*cpu).IMM, 2}, {"DEX", "IMP", (*cpu).DEX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"CPY", "ABS", (*cpu).CPY, (*cpu).ABS, 4}, {"CMP", "ABS", (*cpu).CMP, (*cpu).ABS, 4}, {"DEC", "ABS", (*cpu).DEC, (*cpu).ABS, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6},
		{"BNE", "REL", (*cpu).BNE, (*cpu).REL, 2}, {"CMP", "IZY", (*cpu).CMP, (*cpu).IZY, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"CMP", "ZPX", (*cpu).CMP, (*cpu).ZPX, 4}, {"DEC", "ZPX", (*cpu).DEC, (*cpu).ZPX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"CLD", "IMP", (*cpu).CLD, (*cpu).IMP, 2}, {"CMP", "ABY", (*cpu).CMP, (*cpu).ABY, 4}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"CMP", "ABX", (*cpu).CMP, (*cpu).ABX, 4}, {"DEC", "ABX", (*cpu).DEC, (*cpu).ABX, 7}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7},
		{"CPX", "IMM", (*cpu).CPX, (*cpu).IMM, 2}, {"SBC", "IZX", (*cpu).SBC, (*cpu).IZX, 6}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"CPX", "ZP0", (*cpu).CPX, (*cpu).ZP0, 3}, {"SBC", "ZP0", (*cpu).SBC, (*cpu).ZP0, 3}, {"INC", "ZP0", (*cpu).INC, (*cpu).ZP0, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"INX", "IMP", (*cpu).INX, (*cpu).IMP, 2}, {"SBC", "IMM", (*cpu).SBC, (*cpu).IMM, 2}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"CPX", "ABS", (*cpu).CPX, (*cpu).ABS, 4}, {"SBC", "ABS", (*cpu).SBC, (*cpu).ABS, 4}, {"INC", "ABS", (*cpu).INC, (*cpu).ABS, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6},
		{"BEQ", "REL", (*cpu).BEQ, (*cpu).REL, 2}, {"SBC", "IZY", (*cpu).SBC, (*cpu).IZY, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"SBC", "ZPX", (*cpu).SBC, (*cpu).ZPX, 4}, {"INC", "ZPX", (*cpu).INC, (*cpu).ZPX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"SED", "IMP", (*cpu).SED, (*cpu).IMP, 2}, {"SBC", "ABY", (*cpu).SBC, (*cpu).ABY, 4}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"SBC", "ABX", (*cpu).SBC, (*cpu).ABX, 4}, {"INC", "ABX", (*cpu).INC, (*cpu).ABX, 7}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7},
	}

	return cpu
}

func (c *CPU6502) Read(addr uint16) uint8 {
	return c.bus.Read(addr, false) // readOnly = false
}

func (c *CPU6502) Write(addr uint16, data uint8) {
	c.bus.Write(addr, data)
}

func (c *CPU6502) Fetch() uint8 {
	if c.lookup[c.opcode].AddrName != "IMP" {
		c.fetched = c.Read(c.addrAbs)
	}
	return c.fetched
}

func (c *CPU6502) Clock() {
	if c.cycles == 0 {
		c.opcode = c.Read(c.pc)
		c.pc++
		c.cycles = c.lookup[c.opcode].Cycles
		additionalCycle1 := c.lookup[c.opcode].AddrMode()
		additionalCycle2 := c.lookup[c.opcode].Operate()

		c.cycles += (additionalCycle1 & additionalCycle2)
	}
	c.cycles--
}

func (c *CPU6502) Complete() bool {
	return c.cycles == 0
}

func (c *CPU6502) Reset() {
	c.a = 0
	c.x = 0
	c.y = 0
	c.stkp = 0xFD
	c.status = 0x00 | uint8(U)
	c.addrAbs = 0xFFFC
	lo := uint16(c.Read(c.addrAbs + 0))
	hi := uint16(c.Read(c.addrAbs + 1))
	c.pc = (hi << 8) | lo
	c.addrRel = 0x0000
	c.addrAbs = 0x0000
	c.fetched = 0x00
	c.cycles = 8
}

func (c *CPU6502) IRQ() {
	if !c.GetFlag(I) {
		c.Write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
		c.stkp--
		c.Write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
		c.stkp--

		c.SetFlag(B, false)
		c.SetFlag(U, true)
		c.SetFlag(I, true)
		c.Write(0x0100+uint16(c.stkp), c.status)
		c.stkp--

		c.addrAbs = 0xFFFE
		lo := uint16(c.Read(c.addrAbs + 0))
		hi := uint16(c.Read(c.addrAbs + 1))
		c.pc = (hi << 8) | lo
		c.cycles = 7

	}
}
func (c *CPU6502) NMI() {
	c.Write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
	c.stkp--
	c.Write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
	c.stkp--

	c.SetFlag(B, false)
	c.SetFlag(U, true)
	c.SetFlag(I, true)
	c.Write(0x0100+uint16(c.stkp), c.status)
	c.stkp--

	c.addrAbs = 0xFFFA
	lo := uint16(c.Read(c.addrAbs + 0))
	hi := uint16(c.Read(c.addrAbs + 1))
	c.pc = (hi << 8) | lo
	c.cycles = 8
}

func (c *CPU6502) LoadProgram(instructions []string, offset uint16) {
	for i, instruction := range instructions {
		converted, _ := strconv.ParseUint(instruction, 16, 8)
		c.Write(offset+uint16(i), uint8(converted))
	}
}
