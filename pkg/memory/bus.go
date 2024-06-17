package memory

type Bus interface {
	Read(addr uint16, readOnly bool) uint8
	Write(addr uint16, data uint8)
}

type DefaultBus struct {
	//cpu CPU6502
	ram []uint8
}

func NewDefaultBus(ramLength int) *DefaultBus {
	b := &DefaultBus{ram: make([]uint8, ramLength)}
	for i := range b.ram {
		b.ram[i] = 0x00
	}
	return b
}


func (b *DefaultBus) Write(addr uint16, data uint8) {
	if addr >= 0x0000 && addr <= 0xFFFF {
		b.ram[addr & 0x07FF] = data
	}
}

func (b *DefaultBus) Read(addr uint16, readOnly bool) uint8 {
	var data uint8 = 0x00
	if addr >= 0x0000 && addr <= 0xFFFF {
		data = b.ram[addr & 0x07FF]
	}
	return data
}
