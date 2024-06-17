package ppu

import "nes-emulator/pkg/memory"

type PPU2C02 struct {
	cpuBus memory.Bus
	ppuBus memory.Bus
}

func NewPPU2C02(cpuBus memory.Bus, ppuBus memory.Bus) *PPU2C02 {
	return &PPU2C02{
		cpuBus: cpuBus,
		ppuBus: ppuBus,
	}
}

func (p *PPU2C02) cpuRead(addr uint16) uint8 {
	return p.cpuBus.Read(addr, false) // readOnly = false
}

func (p *PPU2C02) cpuWrite(addr uint16, data uint8) {
	p.cpuBus.Write(addr, data)
}

