package cpu

func (c *CPU6502) GetA() uint8 {
	return c.a
}

func (c *CPU6502) GetX() uint8 {
	return c.x
}

func (c *CPU6502) GetY() uint8 {
	return c.y
}

func (c *CPU6502) GetStkp() uint8 {
	return c.stkp
}

func (c *CPU6502) GetPc() uint16 {
	return c.pc
}

func (c *CPU6502) GetStatus() uint8 {
	return c.status
}

func (c *CPU6502) GetCFlag() bool {
	return c.GetFlag(C)
}

func (c *CPU6502) GetZFlag() bool {
	return c.GetFlag(Z)
}

func (c *CPU6502) GetIFlag() bool {
	return c.GetFlag(I)
}

func (c *CPU6502) GetDFlag() bool {
	return c.GetFlag(D)
}

func (c *CPU6502) GetBFlag() bool {
	return c.GetFlag(B)
}

func (c *CPU6502) GetUFlag() bool {
	return c.GetFlag(U)
}

func (c *CPU6502) GetVFlag() bool {
	return c.GetFlag(V)
}

func (c *CPU6502) GetNFlag() bool {
	return c.GetFlag(N)
}

//func (c *CPU6502) GetBus() Bus {
//	return c.bus
//}

func (c *CPU6502) GetFetched() uint8 {
	return c.fetched
}

func (c *CPU6502) GetAddrAbs() uint16 {
	return c.addrAbs
}

func (c *CPU6502) GetAddrRel() uint16 {
	return c.addrRel
}

func (c *CPU6502) GetOpcode() uint8 {
	return c.opcode
}

func (c *CPU6502) GetCycles() uint8 {
	return c.cycles
}

func (c *CPU6502) GetState() (uint8, uint8, uint8, uint8, uint16, uint8, uint8, uint16, uint16, uint8, uint8) {
	return c.a, c.x, c.y, c.stkp, c.pc, c.status, c.fetched, c.addrAbs, c.addrRel, c.opcode, c.cycles
}

func (c *CPU6502) GetLookup() [256]Instruction {
	return c.lookup
}

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

func (c *CPU6502) SetIFlag(value bool) {
	c.SetFlag(I, value)
}

func (c *CPU6502) SetDFlag(value bool) {
	c.SetFlag(D, value)
}

func (c *CPU6502) SetBFlag(value bool) {
	c.SetFlag(B, value)
}

func (c *CPU6502) SetUFlag(value bool) {
	c.SetFlag(U, value)
}

func (c *CPU6502) SetVFlag(value bool) {
	c.SetFlag(V, value)
}

func (c *CPU6502) SetNFlag(value bool) {
	c.SetFlag(N, value)
}

func (c *CPU6502) SetCFlag(value bool) {
	c.SetFlag(C, value)
}

func (c *CPU6502) SetZFlag(value bool) {
	c.SetFlag(Z, value)
}

func (c *CPU6502) SetA(value uint8) {
	c.a = value
}

func (c *CPU6502) SetX(value uint8) {
	c.x = value
}

func (c *CPU6502) SetY(value uint8) {
	c.y = value
}

func (c *CPU6502) SetStkp(value uint8) {
	c.stkp = value
}

func (c *CPU6502) SetPc(value uint16) {
	c.pc = value
}

func (c *CPU6502) SetStatus(value uint8) {
	c.status = value
}

func (c *CPU6502) SetFetched(value uint8) {
	c.fetched = value
}

func (c *CPU6502) SetAddrAbs(value uint16) {
	c.addrAbs = value
}

func (c *CPU6502) SetAddrRel(value uint16) {
	c.addrRel = value
}

func (c *CPU6502) SetOpcode(value uint8) {
	c.opcode = value
}

func (c *CPU6502) SetCycles(value uint8) {
	c.cycles = value
}
