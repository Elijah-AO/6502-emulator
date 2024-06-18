package ppu

import (
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/memory"
)

type PPU2C02 struct {
	cpuBus memory.Bus
	ppuBus memory.Bus
	cart   cartridge.Cartridge
	tblName [2][1024]uint8
	//tblPattern [2][4096]uint8 
	tblPalette [32]uint8
}

func NewPPU2C02(cpuBus memory.Bus, ppuBus memory.Bus, cart cartridge.Cartridge) *PPU2C02 {
	return &PPU2C02{
		cpuBus: cpuBus,
		ppuBus: ppuBus,
		cart:   cartridge.Cartridge{},
	}
}

func (p *PPU2C02) cpuRead(addr uint16) uint8 {
	var data uint8 = 0x00
	switch addr {
	case 0x0000:
		// Control
	case 0x0001:
		// Mask
	case 0x0002:
		// Status
	case 0x0003:
		// OAM Address
	case 0x0004:
		// OAM Data
	case 0x0005:
		// Scroll
	case 0x0006:
		// PPU Address
	case 0x0007:
		// PPU Data
	}
	return data
}

func (p *PPU2C02) cpuWrite(addr uint16, data uint8) {
	switch addr {
	case 0x0000:
		// Control
	case 0x0001:
		// Mask
	case 0x0002:
		// Status
	case 0x0003:
		// OAM Address
	case 0x0004:
		// OAM Data
	case 0x0005:
		// Scroll
	case 0x0006:
		// PPU Address
	case 0x0007:
		// PPU Data
	}
}

func (p *PPU2C02) ppuRead(addr uint16) uint8 {
	var data uint8 = 0x00
	addr &= 0x3FFF
	return data
}

func (p *PPU2C02) ppuWrite(addr uint16, data uint8) {
	addr &= 0x3FFF
}
