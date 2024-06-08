package main

import (
	"fmt"
	"nes-emulator/pkg/cpu"
	//"time"

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

	a, x, y, stkp, pc, status, fetched, addrAbs, addrRel, opcode, cycles := v.cpu.GetState()
	fmt.Fprintf(txt, "\nA: $%02X [%d] \nX: $%02X [%d] \nY: $%02X [%d] \nStack P: $%04X \nPC: $%04X \n\nStatus: $%04X \nFetched: $%04X \nAddrAbs: $%04X \nAddrRel: $%04X \nOpcode: $%X \nCycles: %d \n", a, a, x, x, y, y, stkp, pc, status, fetched, addrAbs, addrRel, opcode, cycles)
}

func (v *visualisation) displayCode(txt *text.Text, lines, start int) {
	disassembly := v.cpu.Disassemble(0x0000, 0xFFFF)
	curr := disassembly[uint16(start)]
	for i := 1; i < lines; i++ {
		if curr != "" {
			fmt.Fprintln(txt, curr)
		}
		curr = disassembly[uint16(start)+uint16(i)]
	}
}

func run() {
	bus := cpu.NewDefaultBus()
	cpu := cpu.NewCPU6502()
	cpu.ConnectBus(bus)
	vis := NewVisualisation(cpu, bus)

	cpu.Write(0xFFFC, 0x00)
	cpu.Write(0xFFFD, 0x80)
	instructions := []string{"A2", "0A", "8E", "00", "00", "A2", "03", "8E", "01", "00", "AC", "00", "00", "A9", "00", "18", "6D", "01", "00", "88", "D0", "FA", "8D", "02", "00", "EA", "EA", "EA"}
	cpu.LoadProgram(instructions, 0x8000)
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
	txtCode := text.New(pixel.V(800, 556), atlas)

	for !win.Closed() {
		win.Clear(colornames.Darkblue)
		txtCPU.Clear()
		txt.Clear()
		txtCode.Clear()
		switch {
		case win.Pressed(pixelgl.KeyI):
			cpu.SetIFlag(!cpu.GetIFlag())
		case win.Pressed(pixelgl.KeyN):
			cpu.SetNFlag(!cpu.GetNFlag())
		case win.Pressed(pixelgl.KeySpace):
			cpu.Clock()
		case win.Pressed(pixelgl.KeyR):
			cpu.Reset()
		}
		//time.Sleep(100 * time.Millisecond)

		vis.displayMemory(txt, 0x0000, 15, 16)
		fmt.Fprintf(txt, "\n")
		vis.displayMemory(txt, 0x8000, 15, 16)
		vis.displayCPU(txtCPU)
		vis.displayCode(txtCode, 50, int(vis.cpu.GetPc()))
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))
		txtCPU.Draw(win, pixel.IM.Scaled(txtCPU.Orig, 2))
		txtCode.Draw(win, pixel.IM.Scaled(txtCode.Orig, 2))

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
