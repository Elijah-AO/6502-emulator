package main

import (
	"fmt"
	"nes-emulator/pkg/cpu"
	"nes-emulator/pkg/memory"
)

func main() {
	fmt.Println("Hello, NES!")
	cpu := cpu.NewCPU6502()
	bus := memory.NewBus(cpu)
	bus.Write(0xFFFC, 0x52)
	fmt.Println(bus.Read(0xFFFC, false))
}
