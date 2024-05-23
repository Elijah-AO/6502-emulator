package cpu

import (
	"nes-emulator/pkg/memory"
)

type CPU6502 struct {
	bus memory.Bus
}

func NewCPU6502() CPU6502 {
	return CPU6502{}
}

func (b memory.Bus) ConnectBus(cpu *CPU6502) {
	cpu.bus = b
}