package nes

import (
	"nes-emulator/pkg/memory"
	"nes-emulator/pkg/cpu"
	"nes-emulator/pkg/ppu"
	"nes-emulator/pkg/cartridge"
)

type NES struct {
	cpu *cpu.CPU6502
	ppu *ppu.PPU2C02
	cart *cartridge.Cartridge
	cpuBus *memory.DefaultBus
	ppuBus *memory.DefaultBus

	nSystemClockCounter uint64
	nFrameCounter uint64

}

func NewNES(filename string) *NES {
	cpuBus := memory.NewDefaultBus(2*1024)
	ppuBus := memory.NewDefaultBus(8*1024)
	cart := cartridge.NewCartridge(cpuBus, ppuBus, filename)	
	cpu := cpu.NewCPU6502(cpuBus)
	ppu := ppu.NewPPU2C02(cpuBus, ppuBus, *cart)
	return &NES{cpu: cpu, ppu: ppu, cart: cart, cpuBus: cpuBus, ppuBus: ppuBus}
}

func (n *NES) InsertCartridge(cart *cartridge.Cartridge) {
	n.cart = cart
}

func (n *NES) Reset() {
	n.cpu.Reset()
	n.nSystemClockCounter = 0
	n.nFrameCounter = 0
}

func (n *NES) Clock() {
	n.cpu.Clock()
	n.nSystemClockCounter++
}
