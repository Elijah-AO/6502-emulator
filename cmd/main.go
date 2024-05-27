package main

import (
	"time"
	//"time"
	"fmt"
	"nes-emulator/pkg/cpu"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type visualisation struct {
	cpu *cpu.CPU6502
	bus cpu.Bus
}

func NewVisualisation(cpu *cpu.CPU6502, bus cpu.Bus) *visualisation {
	return &visualisation{cpu: cpu, bus: bus}
}

func (v *visualisation) displayMemory(txt *text.Text, startAddr uint16, numRows, numColumns int) {
	for row := 0; row < numRows; row++ {
		address := startAddr + uint16(row*numColumns)
		fmt.Fprintf(txt, "$%04X: ", address)
		for col := 0; col < numColumns; col++ {
			data := v.bus.Read(address+uint16(col), true)
			fmt.Fprintf(txt, "%02X ", data)
		}
		fmt.Fprintln(txt)
	}
}

func (v *visualisation) displayCPU(txt *text.Text) {
	flags := [8]string{"C", "Z", "I", "D", "B", "-", "V", "N"}
	flagArray := [8]bool{v.cpu.GetCFlag(), v.cpu.GetZFlag(), v.cpu.GetIFlag(), v.cpu.GetDFlag(), v.cpu.GetBFlag(), v.cpu.GetUFlag(), v.cpu.GetVFlag(), v.cpu.GetNFlag()}
	fmt.Fprintf(txt, "Status: ")
		for i := 0; i < 8; i++ {
			if flagArray[i] {
				txt.Color = colornames.Lightgreen
			} else {
				txt.Color = colornames.Red 
			}
			fmt.Fprintf(txt, " %s", flags[i])
		}
	txt.Color = colornames.White
	time.Sleep(50 * time.Millisecond)
	a, x, y, stkp, pc, status, fetched, addrAbs, addrRel, opcode, cycles := v.cpu.GetState()
	fmt.Fprintf(txt, "\nA: $%02X [%d] \nX: $%02X [%d] \nY: $%02X [%d] \nStack P: $%04X \nPC: $%04X \n\nStatus: $%04X \nFetched: $%04X \nAddrAbs: $%04X \nAddrRel: $%04X \nOpcode: $%X \nCycles: %d \n", a, a, x, x, y, y, stkp, pc, status, fetched, addrAbs, addrRel, opcode, cycles)
}

func run() {
	bus := cpu.NewDefaultBus()
	cpu := cpu.NewCPU6502()
	cpu.ConnectBus(bus)
	vis := NewVisualisation(cpu, bus)
	cpu.Reset()

	cfg := pixelgl.WindowConfig{
		Title:  "NES Emulator",
		Bounds: pixel.R(0, 0, 1280, 1000),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(10, 970), atlas)
	txt.Color = colornames.White
	txtCPU := text.New(pixel.V(800, 970), atlas)

	for !win.Closed() {
		win.Clear(colornames.Darkblue)
        txtCPU.Clear()
		txt.Clear()		
		switch {
		case win.Pressed(pixelgl.KeyI):
			cpu.SetIFlag(!cpu.GetIFlag())
		case win.Pressed(pixelgl.KeyC):
			cpu.SetCFlag(!cpu.GetCFlag())
		case win.Pressed(pixelgl.KeyZ):
			cpu.SetZFlag(!cpu.GetZFlag())
		case win.Pressed(pixelgl.KeyD):
			cpu.SetDFlag(!cpu.GetDFlag())
		case win.Pressed(pixelgl.KeyB):
			cpu.SetBFlag(!cpu.GetBFlag())
		case win.Pressed(pixelgl.KeyU):
			cpu.SetUFlag(!cpu.GetUFlag())
		case win.Pressed(pixelgl.KeyV):
			cpu.SetVFlag(!cpu.GetVFlag())
		case win.Pressed(pixelgl.KeyN):
			cpu.SetNFlag(!cpu.GetNFlag())
		case win.Pressed(pixelgl.KeySpace):
			cpu.Clock()
		}
		
		vis.displayMemory(txt, 0x0000, 15, 16)
		fmt.Fprintf(txt, "\n")	
		vis.displayMemory(txt, 0x8000, 15, 16)
		vis.displayCPU(txtCPU)
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))
		txtCPU.Draw(win, pixel.IM.Scaled(txtCPU.Orig, 2))
		
        win.Update()
    }
}

func main() {
    pixelgl.Run(run)
}
/*
TODO:
Understand PC 8000 at reset
Test some instructions
add help to the bottom of the screen





*/