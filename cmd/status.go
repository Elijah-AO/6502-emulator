package main

import (
	"fmt"
	"nes-emulator/pkg/cpu"
)

func main1() {
	bus := cpu.NewDefaultBus()
	cpu := cpu.NewCPU6502()
	cpu.ConnectBus(bus)
	cpu.Write(0x0000, 0xF5)
	cpu.Write(0x0000, 0x00)
	fmt.Printf("Read from 0x0000: %d\n", cpu.Read(0x0000))
	fmt.Println("Instructions loaded: ", cpu.ReturnLookup())
}

// TODO: Implement pixel visualisation
