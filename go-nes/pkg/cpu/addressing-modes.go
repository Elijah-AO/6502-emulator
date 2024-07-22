package cpu

// Addressing Modes

// Implied
func (c *CPU6502) IMP() uint8 {
	c.fetched = c.a
	return 0
}

// Immediate
func (c *CPU6502) IMM() uint8 {
	c.addrAbs = c.pc
	c.pc++
	return 0
}

// Zero Page
func (c *CPU6502) ZP0() uint8 {
	c.addrAbs = uint16(c.Read(uint16(c.pc)))
	c.pc++
	c.addrAbs &= 0x00FF
	return 0
}

// Zero Page with X Offset
func (c *CPU6502) ZPX() uint8 {
	c.addrAbs = uint16(c.Read(uint16(c.pc) + uint16(c.x)))
	c.pc++
	c.addrAbs &= 0x00FF
	return 0
}

// Zero Page with Y Offset
func (c *CPU6502) ZPY() uint8 {
	c.addrAbs = uint16(c.Read(uint16(c.pc) + uint16(c.y)))
	c.pc++
	c.addrAbs &= 0x00FF
	return 0
}

// Relative
func (c *CPU6502) REL() uint8 {
	c.addrRel = uint16(c.Read(uint16(c.pc)))
	c.pc++
	if (c.addrRel & 0x80) != 0 {
		c.addrRel |= 0xFF00
	}
	return 0
}

// Absolute
func (c *CPU6502) ABS() uint8 {
	lo := uint16(c.Read(uint16(c.pc)))
	c.pc++
	hi := uint16(c.Read(uint16(c.pc)))
	c.pc++
	c.addrAbs = (hi << 8) | lo
	return 0
}

// Absolute with X Offset
func (c *CPU6502) ABX() uint8 {
	lo := uint16(c.Read(uint16(c.pc)))
	c.pc++
	hi := uint16(c.Read(uint16(c.pc)))
	c.pc++
	c.addrAbs = (hi << 8) | lo
	c.addrAbs += uint16(c.x)
	if (c.addrAbs & 0xFF00) != (hi << 8) {
		return 1
	}
	return 0
}

// Absolute with Y Offset
func (c *CPU6502) ABY() uint8 {
	lo := uint16(c.Read(uint16(c.pc)))
	c.pc++
	hi := uint16(c.Read(uint16(c.pc)))
	c.pc++
	c.addrAbs = (hi << 8) | lo
	c.addrAbs += uint16(c.y)
	if (c.addrAbs & 0xFF00) != (hi << 8) {
		return 1
	}
	return 0
}

// Indirect
func (c *CPU6502) IND() uint8 {
	ptrLo := uint16(c.Read(uint16(c.pc)))
	c.pc++
	ptrHi := uint16(c.Read(uint16(c.pc)))
	c.pc++

	ptr := (ptrHi << 8) | ptrLo

	if ptrLo == 0x00FF {
		c.addrAbs = (uint16(c.Read(ptr&0xFF00)) << 8) | uint16(c.Read(ptr+0))
	} else {
		c.addrAbs = (uint16(c.Read(ptr+1) << 8)) | uint16(c.Read(ptr+0))
	}
	return 0
}

// Indirect Zero Page with X Offset
func (c *CPU6502) IZX() uint8 {

	t := uint16(c.Read(uint16(c.pc)))
	c.pc++

	lo := uint16(c.Read((t + uint16(c.x)) & 0x00FF))
	hi := uint16(c.Read((t + uint16(c.x) + 1) & 0x00FF))

	c.addrAbs = (hi << 8) | lo

	return 0
}

// Indirect Zero Page with Y Offset
func (c *CPU6502) IZY() uint8 {

	t := uint16(c.Read(uint16(c.pc)))
	c.pc++

	lo := uint16(c.Read(t & 0x00FF))
	hi := uint16(c.Read((t + 1) & 0x00FF))

	c.addrAbs = (hi << 8) | lo
	c.addrAbs += uint16(c.y)

	if (c.addrAbs & 0xFF00) != (hi << 8) {
		return 1
	}

	return 0
}

