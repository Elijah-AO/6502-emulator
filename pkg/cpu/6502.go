package cpu

type CPU6502 struct {
	bus     Bus
	lookup  [3]Instruction
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

func NewCPU6502() *CPU6502 {
	cpu := &CPU6502{
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
	cpu.lookup = [3]Instruction{{"BRK", "IMM", (*cpu).BRK, (*cpu).IMM, false, 7}, 
								{"RTI", "IMP", (*cpu).RTI, (*cpu).IMP, true, 6},
								{"LDA", "IMM", (*cpu).LDA, (*cpu).IMM, false, 2}}
	return cpu
}

func (c *CPU6502) GetState() (uint8, uint8, uint8, uint8, uint16, uint8, uint8, uint16, uint16, uint8, uint8) {
	return c.a, c.x, c.y, c.stkp, c.pc, c.status, c.fetched, c.addrAbs, c.addrRel, c.opcode, c.cycles
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
	return (c.status & uint8(flag)) > 0
}

func (c *CPU6502) SetFlag(flag Flags6502, value bool) {
	if value {
		c.status |= uint8(flag)
	} else {
		c.status &= ^uint8(flag)
	}
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
// Registers

// Addressing modes

func (c *CPU6502) Clock() {
	if c.cycles == 0 {
		c.opcode = c.Read(c.pc)
		c.pc++
		c.cycles = c.lookup[c.opcode].Cycles
		additionalCycle1 := c.lookup[c.opcode].AddressMode()
		additionalCycle2 := c.lookup[c.opcode].Operate()

		c.cycles += (additionalCycle1 & additionalCycle2)
	}
	c.cycles--
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
		c.Write(0x0100+uint16(c.stkp), uint8((c.pc >> 8) & 0x00FF))
		c.stkp--
		c.Write(0x0100+uint16(c.stkp), uint8(c.pc & 0x00FF))
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
	c.Write(0x0100+uint16(c.stkp), uint8((c.pc >> 8) & 0x00FF))
	c.stkp--
	c.Write(0x0100+uint16(c.stkp), uint8(c.pc & 0x00FF))
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

func (c *CPU6502) Fetch() uint8 {
	if !c.lookup[c.opcode].IMPFlag {
		c.fetched = c.Read(c.addrAbs)
	}
	return c.fetched
}

type Instruction struct {
	Name        string
	AddressName string
	Operate     func() uint8
	AddressMode func() uint8
	IMPFlag    bool
	Cycles      uint8
}

func (c *CPU6502) ReturnLookup() [3]Instruction {
	return c.lookup
}

