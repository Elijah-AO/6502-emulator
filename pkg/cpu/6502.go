package cpu

type CPU6502 struct {
	bus Bus
}

func NewCPU6502() *CPU6502 {
	return &CPU6502{}
}

func (c *CPU6502) ConnectBus(bus Bus) {
	c.bus = bus
}

func (c *CPU6502) Read(addr uint16) uint8 {
	return c.bus.Read(addr, false) // readOnly = false
}

func (c *CPU6502) Write(addr uint16, data uint8) {
	c.bus.Write(addr, data)
}

// TODO: Implement the get and set flag methods

func (c *CPU6502) GetFlag(flag Flags6502) bool {
	return true
}

func (c *CPU6502) SetFlag(flag Flags6502, value bool) {
	return
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

// Registers
var a uint8 = 0x00
var x uint8 = 0x00
var y uint8 = 0x00
var stkp uint8 = 0x00
var pc uint16 = 0x0000
var status uint8 = 0x00

// Addressing modes
func (c *CPU6502) IMP() uint8 { return 0 }
func (c *CPU6502) IMM() uint8 { return 0 }
func (c *CPU6502) ZP0() uint8 { return 0 }
func (c *CPU6502) ZPX() uint8 { return 0 }
func (c *CPU6502) ZPY() uint8 { return 0 }
func (c *CPU6502) REL() uint8 { return 0 }
func (c *CPU6502) ABS() uint8 { return 0 }
func (c *CPU6502) ABX() uint8 { return 0 }
func (c *CPU6502) ABY() uint8 { return 0 }
func (c *CPU6502) IND() uint8 { return 0 }
func (c *CPU6502) IZX() uint8 { return 0 }
func (c *CPU6502) IZY() uint8 { return 0 }

/*
const (
	// Status register flags
	C Flags6502 = 1 << iota
	Z
	I
	D
	B
	U
	V
	N
)
*/
// TODO: Refactor the instructions to the instructions.go

// Instructions
func (c *CPU6502) ADC() uint8 { return 0 }
func (c *CPU6502) AND() uint8 { return 0 }
func (c *CPU6502) ASL() uint8 { return 0 }
func (c *CPU6502) BCC() uint8 { return 0 }
func (c *CPU6502) BCS() uint8 { return 0 }
func (c *CPU6502) BEQ() uint8 { return 0 }
func (c *CPU6502) BIT() uint8 { return 0 }
func (c *CPU6502) BMI() uint8 { return 0 }
func (c *CPU6502) BNE() uint8 { return 0 }
func (c *CPU6502) BPL() uint8 { return 0 }
func (c *CPU6502) BRK() uint8 { return 0 }
func (c *CPU6502) BVC() uint8 { return 0 }
func (c *CPU6502) BVS() uint8 { return 0 }
func (c *CPU6502) CLC() uint8 { return 0 }
func (c *CPU6502) CLD() uint8 { return 0 }
func (c *CPU6502) CLI() uint8 { return 0 }
func (c *CPU6502) CLV() uint8 { return 0 }
func (c *CPU6502) CMP() uint8 { return 0 }
func (c *CPU6502) CPX() uint8 { return 0 }
func (c *CPU6502) CPY() uint8 { return 0 }
func (c *CPU6502) DEC() uint8 { return 0 }
func (c *CPU6502) DEX() uint8 { return 0 }
func (c *CPU6502) DEY() uint8 { return 0 }
func (c *CPU6502) EOR() uint8 { return 0 }
func (c *CPU6502) INC() uint8 { return 0 }
func (c *CPU6502) INX() uint8 { return 0 }
func (c *CPU6502) INY() uint8 { return 0 }
func (c *CPU6502) JMP() uint8 { return 0 }
func (c *CPU6502) JSR() uint8 { return 0 }
func (c *CPU6502) LDA() uint8 { return 0 }
func (c *CPU6502) LDX() uint8 { return 0 }
func (c *CPU6502) LDY() uint8 { return 0 }
func (c *CPU6502) LSR() uint8 { return 0 }
func (c *CPU6502) NOP() uint8 { return 0 }
func (c *CPU6502) ORA() uint8 { return 0 }
func (c *CPU6502) PHA() uint8 { return 0 }
func (c *CPU6502) PHP() uint8 { return 0 }
func (c *CPU6502) PLA() uint8 { return 0 }
func (c *CPU6502) PLP() uint8 { return 0 }
func (c *CPU6502) ROL() uint8 { return 0 }
func (c *CPU6502) ROR() uint8 { return 0 }
func (c *CPU6502) RTI() uint8 { return 0 }
func (c *CPU6502) RTS() uint8 { return 0 }
func (c *CPU6502) SBC() uint8 { return 0 }
func (c *CPU6502) SEC() uint8 { return 0 }
func (c *CPU6502) SED() uint8 { return 0 }
func (c *CPU6502) SEI() uint8 { return 0 }
func (c *CPU6502) STA() uint8 { return 0 }
func (c *CPU6502) STX() uint8 { return 0 }
func (c *CPU6502) STY() uint8 { return 0 }
func (c *CPU6502) TAX() uint8 { return 0 }
func (c *CPU6502) TAY() uint8 { return 0 }
func (c *CPU6502) TSX() uint8 { return 0 }
func (c *CPU6502) TXA() uint8 { return 0 }
func (c *CPU6502) TXS() uint8 { return 0 }
func (c *CPU6502) TYA() uint8 { return 0 }
