package cartridge

import (
	"fmt"
	"nes-emulator/pkg/memory"
	"os"
	"encoding/binary"
)
type Cartridge struct{
	cpuBus memory.Bus
	ppuBus memory.Bus
	prgRom []uint8
	chrRom []uint8
	mapperID uint8
	nPRGBanks uint8
	nCHRBanks uint8
}

type sHeader struct {
	Name          [4]byte
	PRG_ROM_chunks uint8
	CHR_ROM_chunks uint8
	Mapper1        uint8
	Mapper2        uint8
	PRG_RAM_size   uint8
	TV_System1     uint8
	TV_System2     uint8
	Unused         [5]uint8
}


func readHeader(filename string) (sHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	header := sHeader{}
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return sHeader{}, fmt.Errorf("error reading header: %w", err)
	}

	return header, nil
}

func fileType(){}


func NewCartridge(cpuBus memory.Bus, ppuBus memory.Bus, filename string) *Cartridge {
	header, err := readHeader(filename)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if header.Mapper1&0x04 > 0 {
		fmt.Println("Error: ", "Trainer not supported")
	}

	mapperID := ((header.Mapper2 >> 4) << 4) | (header.Mapper1 >> 4)

	return &Cartridge{
		cpuBus: cpuBus,
		ppuBus: ppuBus,
		mapperID: mapperID,
	}
}

func (c *Cartridge) cpuRead(addr uint16) uint8 {
	return 0

}

func (c *Cartridge) cpuWrite(addr uint16, data uint8) {
}

func (c *Cartridge) ppuRead(addr uint16) uint8 {
	return 0
}

func (c *Cartridge) ppuWrite(addr uint16, data uint8) {
}

func (c *Cartridge) insertCartridge() {

}
