package memory

import "nes-emulator/pkg/cpu"

type Bus struct {
	cpu cpu.CPU6502
	ram [64 * 1024]uint8
	// TODO: Clear RAM contents
}

func NewBus(cpu cpu.CPU6502) Bus {
	bus := Bus{
		cpu: cpu,
	}
	// Initialize RAM to zero
	for i := range bus.ram {
		bus.ram[i] = 0
	}
	return bus
}
func (b *Bus) Write(addr uint16, data uint8) {
	if addr >= 0x0000 && addr <= 0xFFFF {
		b.ram[addr] = data
	}
}

func (b *Bus) Read(addr uint16, readOnly bool) uint8 {
	if addr >= 0x0000 && addr <= 0xFFFF {
		return b.ram[addr]
	}
	return 0x00
}

func ReadOnly() bool {
	return false
}
