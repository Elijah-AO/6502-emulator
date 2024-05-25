package cpu

type Bus interface {
	Read(addr uint16, readOnly bool) uint8
	Write(addr uint16, data uint8)
}

type DefaultBus struct {
	cpu CPU6502
	ram [64 * 1024]uint8

	// TODO: Clear RAM contents
}

func NewDefaultBus() *DefaultBus {
	b := &DefaultBus{}
	for i := range b.ram {
		b.ram[i] = 0x00
	}
	b.cpu.ConnectBus(b)
	return b
}

func (b *DefaultBus) Write(addr uint16, data uint8) {
	if addr >= 0x0000 && addr <= 0xFFFF {
		b.ram[addr] = data
	}
}

func (b *DefaultBus) Read(addr uint16, readOnly bool) uint8 {
	if addr >= 0x0000 && addr <= 0xFFFF {
		return b.ram[addr]
	}
	return 0x00
}
