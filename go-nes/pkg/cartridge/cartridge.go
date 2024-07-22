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


func NewCartridge(cpuBus memory.Bus, ppuBus memory.Bus, filename string) *Cartridge {
	file, err := os.Open(filename)
	defer file.Close()
	
	if err != nil {
		fmt.Println("Error: ", err)
	}
	
	header := sHeader{}
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	if header.Mapper1&0x04 > 0 {
		_, err = file.Seek(512, 1)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil
		} 
	}

	mapperID := ((header.Mapper2 >> 4) << 4) | (header.Mapper1 >> 4)

	var fileType uint8 = 1
	var prgRom []uint8

	switch fileType {
	case 0:
		// Do nothing
	case 1:
		nPRGBanks := int(header.PRG_ROM_chunks)
		prgRom := make([]uint8, nPRGBanks*16*1024)
		_, err = file.Read(prgRom)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil
		}

		nCHRBanks := int(header.CHR_ROM_chunks)
		chrRom := make([]uint8, nCHRBanks*8*1024)
		_, err = file.Read(chrRom)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil
		}
		

	case 2:
		// Do nothing
	}


	return &Cartridge{
		cpuBus: cpuBus,
		ppuBus: ppuBus,
		mapperID: mapperID,

		prgRom: prgRom,
	}
}

func (c *Cartridge) cpuRead(addr uint16) bool {
	return false

}

func (c *Cartridge) cpuWrite(addr uint16, data uint8) bool {
	return false
}

func (c *Cartridge) ppuRead(addr uint16) bool {
	return false
}

func (c *Cartridge) ppuWrite(addr uint16, data uint8) bool {
	return false
}

func (c *Cartridge) insertCartridge() {

}
