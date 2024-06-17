package cartridge

import (
	"nes-emulator/pkg/memory"
)
type Cartridge struct{
	bus *memory.DefaultBus
}
