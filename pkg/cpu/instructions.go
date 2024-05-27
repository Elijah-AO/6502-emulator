package cpu

import "fmt"
// Instructions

// ADC - Add Memory to Accumulator with Carry
func (c *CPU6502) ADC() uint8 { return 0 }

// AND - "AND" Memory with Accumulator
func (c *CPU6502) AND() uint8 {
	c.fetched = c.Read(c.addrAbs)
	c.a = c.a & c.fetched
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlag(N, c.a&0x80 != 0)
	return 1
}

// ASL - Shift left One Bit (Memory or Accumulator)
func (c *CPU6502) ASL() uint8 { return 0 }

// BCC - Branch on Carry Clear
func (c *CPU6502) BCC() uint8 { return 0 }

// BCS - Branch on Carry Set
func (c *CPU6502) BCS() uint8 { return 0 }

// BEQ - Branch on Result Zero
func (c *CPU6502) BEQ() uint8 { return 0 }

// BIT - Test Bits in Memory with Accumulator
func (c *CPU6502) BIT() uint8 { return 0 }

// BMI - Branch on Result Minus
func (c *CPU6502) BMI() uint8 { return 0 }

// BNE - Branch on Result not Zero
func (c *CPU6502) BNE() uint8 { return 0 }

// BPL - Branch on Result Plus
func (c *CPU6502) BPL() uint8 { return 0 }

// BRK - Force Break
func (c *CPU6502) BRK() uint8 { 
	c.pc++
	c.SetFlag(I, true)

	c.Write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
	c.stkp--
	c.Write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
	c.stkp--
	c.SetFlag(B, true)
	c.Write(0x0100+uint16(c.stkp), c.status)
	c.stkp--
	c.SetFlag(B, false)
	c.pc = uint16(c.Read(0xFFFE)) | uint16(c.Read(0xFFFF))<<8
	return 0
}

// BVC - Branch on Overflow Clear
func (c *CPU6502) BVC() uint8 { return 0 }

// BVS - Branch on Overflow Set
func (c *CPU6502) BVS() uint8 { return 0 }

// CLC - Clear Carry Flag
func (c *CPU6502) CLC() uint8 { return 0 }

// CLD - Clear Decimal Mode
func (c *CPU6502) CLD() uint8 { return 0 }

// CLI - Clear Interrupt Disable Bit
func (c *CPU6502) CLI() uint8 { return 0 }

// CLV - Clear Overflow Flag
func (c *CPU6502) CLV() uint8 { return 0 }

// CMP - Compare Memory and Accumulator
func (c *CPU6502) CMP() uint8 { return 0 }

// CPX - Compare Memory and Index X
func (c *CPU6502) CPX() uint8 { return 0 }

// CPY - Compare Memory and Index Y
func (c *CPU6502) CPY() uint8 { return 0 }

// DEC - Decrement Memory by One
func (c *CPU6502) DEC() uint8 { return 0 }

// DEX - Decrement Index X by One
func (c *CPU6502) DEX() uint8 { return 0 }

// DEY - Decrement Index Y by One
func (c *CPU6502) DEY() uint8 { return 0 }

// EOR - "Exclusive-or" Memory with Accumulator
func (c *CPU6502) EOR() uint8 { return 0 }

// INC - Increment Memory by One
func (c *CPU6502) INC() uint8 { return 0 }

// INX - Increment Index X by One
func (c *CPU6502) INX() uint8 { return 0 }

// INY - Increment Index Y by One
func (c *CPU6502) INY() uint8 { return 0 }

// JMP - Jump to New Location
func (c *CPU6502) JMP() uint8 { return 0 }

// JSR - Jump to New Location Saving Return Address
func (c *CPU6502) JSR() uint8 { return 0 }

// LDA - Load Accumulator with Memory
func (c *CPU6502) LDA() uint8 { 
	c.Fetch()	
	c.a = c.fetched
	c.SetFlag(Z, c.a == 0x00)
	c.SetFlag(N, c.a&0x80 != 0)
	return 1
}

// LDX - Load Index X with Memory
func (c *CPU6502) LDX() uint8 { return 0 }

// LDY - Load Index Y with Memory
func (c *CPU6502) LDY() uint8 { return 0 }

// LSR - Shift One Bit Right (Memory or Accumulator)
func (c *CPU6502) LSR() uint8 { return 0 }

// NOP - No Operation
func (c *CPU6502) NOP() uint8 { return 0 }

// ORA - "OR" Memory with Accumulator
func (c *CPU6502) ORA() uint8 { return 0 }

// PHA - Push Accumulator on Stack
func (c *CPU6502) PHA() uint8 { return 0 }

// PHP - Push Processor Status on Stack
func (c *CPU6502) PHP() uint8 { return 0 }

// PLA - Pull Accumulator from Stack
func (c *CPU6502) PLA() uint8 { return 0 }

// PLP - Pull Processor Status from Stack
func (c *CPU6502) PLP() uint8 { return 0 }

// ROL - Rotate One Bit Left (Memory or Accumulator)
func (c *CPU6502) ROL() uint8 { return 0 }

// ROR - Rotate One Bit Right (Memory or Accumulator)
func (c *CPU6502) ROR() uint8 { return 0 }

// RTI - Return from Interrupt
func (c *CPU6502) RTI() uint8 { 
	c.stkp++
	c.status = c.Read(0x0100 + uint16(c.stkp))
	c.status &= ^uint8(B)
	c.status &= ^uint8(U)
	c.stkp++
	c.pc = uint16(c.Read(0x0100 + uint16(c.stkp)))
	c.stkp++
	c.pc |= uint16(c.Read(0x0100 + uint16(c.stkp))) << 8
	c.stkp++
	return 0
}

// RTS - Return from Subroutine
func (c *CPU6502) RTS() uint8 { return 0 }

// SBC - Subtract Memory from Accumulator with Borrow
func (c *CPU6502) SBC() uint8 { return 0 }

// SEC - Set Carry Flag
func (c *CPU6502) SEC() uint8 { return 0 }

// SED - Set Decimal Mode
func (c *CPU6502) SED() uint8 { return 0 }

// SEI - Set Interrupt Disable Status
func (c *CPU6502) SEI() uint8 { return 0 }

// STA - Store Accumulator in Memory
func (c *CPU6502) STA() uint8 { return 0 }

// STX - Store Index X in Memory
func (c *CPU6502) STX() uint8 { return 0 }

// STY - Store Index Y in Memory
func (c *CPU6502) STY() uint8 { return 0 }

// TAX - Transfer Accumulator to Index X
func (c *CPU6502) TAX() uint8 { return 0 }

// TAY - Transfer Accumulator to Index Y
func (c *CPU6502) TAY() uint8 { return 0 }

// TSX - Transfer Stack Pointer to Index X
func (c *CPU6502) TSX() uint8 { return 0 }

// TXA - Transfer Index X to Accumulator
func (c *CPU6502) TXA() uint8 { return 0 }

// TXS - Transfer Index X to Stack Pointer
func (c *CPU6502) TXS() uint8 { return 0 }

// TYA - Transfer Index Y to Accumulator
func (c *CPU6502) TYA() uint8 { return 0 }

// XXX - Undefined Operation
func (c *CPU6502) XXX() uint8 { return 0 }


// Disassembler

func (c *CPU6502) Disassemble(start, stop uint16) map[uint16]string {
	var addr uint16 = start
	var value uint8 = 0x00
	//var lo uint8 = 0x00
	//var hi uint8 = 0x00
	var lineAddr uint16 = 0
	var mapLines map[uint16]string = make(map[uint16]string)

	for addr <= stop {
		lineAddr = addr
		line := "$" + fmt.Sprintf("%04X: ", addr)
		opcode := c.bus.Read(addr, true)
		addr++
		line += fmt.Sprintf("%02X ", opcode)
		addrName := c.lookup[opcode].AddressName
		switch addrName {
		case "IMP":
			line += "IMP"
		case "IMM":
			value = c.bus.Read(addr, true)
			addr++
			line += fmt.Sprintf("IMM $%02X", value)
		
		// TODO: Add more address modes
		default:
			line += "UNK"
		}
		mapLines[lineAddr] = line
	}
	return mapLines
}