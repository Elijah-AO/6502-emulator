package cpu

type CPU6502 struct {
	bus     Bus
	lookup  [256]Instruction
	a       uint8
	x       uint8
	y       uint8
	stkp    uint8
	pc      uint16
	//status  Flags6502
	status  uint8
	fetched uint8
	addrAbs uint16
	addrRel uint16
	opcode  uint8
	cycles  uint8
}

type Instruction struct {
	Name        string
	AddrName string
	Operate     func() uint8
	AddrMode func() uint8
	Cycles      uint8
}

type Flags6502 uint8

const (
	// Status register flags
	C Flags6502 = 1 << iota // Carry
	Z                       // Zero
	I                       // Interrupt Disable
	D                       // Decimal Mode (unused in NES)
	B                       // Break
	U                       // Unused
	V                       // Overflow
	N                       // Negative
)



func NewCPU6502() *CPU6502 {
	cpu := &CPU6502{
		a:       0x00,
		x:       0x00,
		y:       0x00,
		stkp:    0x00,
		pc:      0x00,
		status:  0x00,
		fetched: 0x00,
		addrAbs: 0x00,
		addrRel: 0x00,
		opcode:  0x00,
		cycles:  0,
	}
	cpu.lookup = [256]Instruction{
									{"BRK", "IMM", (*cpu).BRK, (*cpu).IMM, 7}, {"ORA", "IZX", (*cpu).ORA, (*cpu).IZX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"???", "IMP", (*cpu).NOP, (*cpu).IMP, 3}, {"ORA", "ZP0", (*cpu).ORA, (*cpu).ZP0, 3}, {"ASL", "ZP0", (*cpu).ASL, (*cpu).ZP0, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 5}, {"PHP", "IMP", (*cpu).PHP, (*cpu).IMP, 3}, {"ORA", "IMM", (*cpu).ORA, (*cpu).IMM, 2}, {"ASL", "IMP", (*cpu).ASL, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"ORA", "ABS", (*cpu).ORA, (*cpu).ABS, 4}, {"ASL", "ABS", (*cpu).ASL, (*cpu).ABS, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6},
									{"BPL", "REL", (*cpu).BPL, (*cpu).REL, 0}, {"ORA", "IZY", (*cpu).ORA, (*cpu).IZY, 5}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 8}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"ORA", "ZPX", (*cpu).ORA, (*cpu).ZPX, 4}, {"ASL", "ZPX", (*cpu).ASL, (*cpu).ZPX, 6}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 6}, {"CLC", "IMP", (*cpu).CLC, (*cpu).IMP, 2}, {"ORA", "ABY", (*cpu).ORA, (*cpu).ABY, 4}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 2}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7}, {"NOP", "IMP", (*cpu).NOP, (*cpu).IMP, 4}, {"ORA", "ABX", (*cpu).ORA, (*cpu).ABX, 4}, {"ASL", "ABX", (*cpu).ASL, (*cpu).ABX, 7}, {"???", "IMP", (*cpu).XXX, (*cpu).IMP, 7},
/*									{ "JSR", &a::JSR, &a::ABS, 6 },            { "AND", &a::AND, &a::IZX, 6 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "BIT", &a::BIT, &a::ZP0, 3 },            { "AND", &a::AND, &a::ZP0, 3 },            { "ROL", &a::ROL, &a::ZP0, 5 },            { "???", &a::XXX, &a::IMP, 5 },            { "PLP", &a::PLP, &a::IMP, 4 },            { "AND", &a::AND, &a::IMM, 2 },            { "ROL", &a::ROL, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 2 },            { "BIT", &a::BIT, &a::ABS, 4 },            { "AND", &a::AND, &a::ABS, 4 },            { "ROL", &a::ROL, &a::ABS, 6 },            { "???", &a::XXX, &a::IMP, 6 },
									
									{ "BMI", &a::BMI, &a::REL, 2 },            { "AND", &a::AND, &a::IZY, 5 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "???", &a::NOP, &a::IMP, 4 },            { "AND", &a::AND, &a::ZPX, 4 },            { "ROL", &a::ROL, &a::ZPX, 6 },            { "???", &a::XXX, &a::IMP, 6 },            { "SEC", &a::SEC, &a::IMP, 2 },            { "AND", &a::AND, &a::ABY, 4 },            { "???", &a::NOP, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 7 },            { "???", &a::NOP, &a::IMP, 4 },            { "AND", &a::AND, &a::ABX, 4 },            { "ROL", &a::ROL, &a::ABX, 7 },            { "???", &a::XXX, &a::IMP, 7 },          

									{ "RTI", &a::RTI, &a::IMP, 6 },            { "EOR", &a::EOR, &a::IZX, 6 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "???", &a::NOP, &a::IMP, 3 },            { "EOR", &a::EOR, &a::ZP0, 3 },            { "LSR", &a::LSR, &a::ZP0, 5 },            { "???", &a::XXX, &a::IMP, 5 },            { "PHA", &a::PHA, &a::IMP, 3 },            { "EOR", &a::EOR, &a::IMM, 2 },            { "LSR", &a::LSR, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 2 },            { "JMP", &a::JMP, &a::ABS, 3 },            { "EOR", &a::EOR, &a::ABS, 4 },            { "LSR", &a::LSR, &a::ABS, 6 },            { "???", &a::XXX, &a::IMP, 6 },          

									{ "BVC", &a::BVC, &a::REL, 2 },            { "EOR", &a::EOR, &a::IZY, 5 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "???", &a::NOP, &a::IMP, 4 },            { "EOR", &a::EOR, &a::ZPX, 4 },            { "LSR", &a::LSR, &a::ZPX, 6 },            { "???", &a::XXX, &a::IMP, 6 },            { "CLI", &a::CLI, &a::IMP, 2 },            { "EOR", &a::EOR, &a::ABY, 4 },            { "???", &a::NOP, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 7 },            { "???", &a::NOP, &a::IMP, 4 },            { "EOR", &a::EOR, &a::ABX, 4 },            { "LSR", &a::LSR, &a::ABX, 7 },            { "???", &a::XXX, &a::IMP, 7 },          

									{ "RTS", &a::RTS, &a::IMP, 6 },            { "ADC", &a::ADC, &a::IZX, 6 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "???", &a::NOP, &a::IMP, 3 },            { "ADC", &a::ADC, &a::ZP0, 3 },            { "ROR", &a::ROR, &a::ZP0, 5 },            { "???", &a::XXX, &a::IMP, 5 },            { "PLA", &a::PLA, &a::IMP, 4 },            { "ADC", &a::ADC, &a::IMM, 2 },            { "ROR", &a::ROR, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 2 },            { "JMP", &a::JMP, &a::IND, 5 },            { "ADC", &a::ADC, &a::ABS, 4 },            { "ROR", &a::ROR, &a::ABS, 6 },            { "???", &a::XXX, &a::IMP, 6 },          

									{ "BVS", &a::BVS, &a::REL, 2 },            { "ADC", &a::ADC, &a::IZY, 5 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "???", &a::NOP, &a::IMP, 4 },            { "ADC", &a::ADC, &a::ZPX, 4 },            { "ROR", &a::ROR, &a::ZPX, 6 },            { "???", &a::XXX, &a::IMP, 6 },            { "SEI", &a::SEI, &a::IMP, 2 },            { "ADC", &a::ADC, &a::ABY, 4 },            { "???", &a::NOP, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 7 },            { "???", &a::NOP, &a::IMP, 4 },            { "ADC", &a::ADC, &a::ABX, 4 },            { "ROR", &a::ROR, &a::ABX, 7 },            { "???", &a::XXX, &a::IMP, 7 },          

									{ "???", &a::NOP, &a::IMP, 2 },            { "STA", &a::STA, &a::IZX, 6 },            { "???", &a::NOP, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 6 },            { "STY", &a::STY, &a::ZP0, 3 },            { "STA", &a::STA, &a::ZP0, 3 },            { "STX", &a::STX, &a::ZP0, 3 },            { "???", &a::XXX, &a::IMP, 3 },            { "DEY", &a::DEY, &a::IMP, 2 },            { "???", &a::NOP, &a::IMP, 2 },            { "TXA", &a::TXA, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 2 },            { "STY", &a::STY, &a::ABS, 4 },            { "STA", &a::STA, &a::ABS, 4 },            { "STX", &a::STX, &a::ABS, 4 },            { "???", &a::XXX, &a::IMP, 4 },          

									{ "BCC", &a::BCC, &a::REL, 2 },            { "STA", &a::STA, &a::IZY, 6 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 6 },            { "STY", &a::STY, &a::ZPX, 4 },            { "STA", &a::STA, &a::ZPX, 4 },            { "STX", &a::STX, &a::ZPY, 4 },            { "???", &a::XXX, &a::IMP, 4 },            { "TYA", &a::TYA, &a::IMP, 2 },            { "STA", &a::STA, &a::ABY, 5 },            { "TXS", &a::TXS, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 5 },            { "???", &a::NOP, &a::IMP, 5 },            { "STA", &a::STA, &a::ABX, 5 },            { "???", &a::XXX, &a::IMP, 5 },            { "???", &a::XXX, &a::IMP, 5 },          

									{ "LDY", &a::LDY, &a::IMM, 2 },            { "LDA", &a::LDA, &a::IZX, 6 },            { "LDX", &a::LDX, &a::IMM, 2 },            { "???", &a::XXX, &a::IMP, 6 },            { "LDY", &a::LDY, &a::ZP0, 3 },            { "LDA", &a::LDA, &a::ZP0, 3 },            { "LDX", &a::LDX, &a::ZP0, 3 },            { "???", &a::XXX, &a::IMP, 3 },            { "TAY", &a::TAY, &a::IMP, 2 },            { "LDA", &a::LDA, &a::IMM, 2 },            { "TAX", &a::TAX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 2 },            { "LDY", &a::LDY, &a::ABS, 4 },            { "LDA", &a::LDA, &a::ABS, 4 },            { "LDX", &a::LDX, &a::ABS, 4 },            { "???", &a::XXX, &a::IMP, 4 },          

									{ "BCS", &a::BCS, &a::REL, 2 },            { "LDA", &a::LDA, &a::IZY, 5 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 5 },            { "LDY", &a::LDY, &a::ZPX, 4 },            { "LDA", &a::LDA, &a::ZPX, 4 },            { "LDX", &a::LDX, &a::ZPY, 4 },            { "???", &a::XXX, &a::IMP, 4 },            { "CLV", &a::CLV, &a::IMP, 2 },            { "LDA", &a::LDA, &a::ABY, 4 },            { "TSX", &a::TSX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 4 },            { "LDY", &a::LDY, &a::ABX, 4 },            { "LDA", &a::LDA, &a::ABX, 4 },            { "LDX", &a::LDX, &a::ABY, 4 },            { "???", &a::XXX, &a::IMP, 4 },          

									{ "CPY", &a::CPY, &a::IMM, 2 },            { "CMP", &a::CMP, &a::IZX, 6 },            { "???", &a::NOP, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "CPY", &a::CPY, &a::ZP0, 3 },            { "CMP", &a::CMP, &a::ZP0, 3 },            { "DEC", &a::DEC, &a::ZP0, 5 },            { "???", &a::XXX, &a::IMP, 5 },            { "INY", &a::INY, &a::IMP, 2 },            { "CMP", &a::CMP, &a::IMM, 2 },            { "DEX", &a::DEX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 2 },            { "CPY", &a::CPY, &a::ABS, 4 },            { "CMP", &a::CMP, &a::ABS, 4 },            { "DEC", &a::DEC, &a::ABS, 6 },            { "???", &a::XXX, &a::IMP, 6 },          

									{ "BNE", &a::BNE, &a::REL, 2 },            { "CMP", &a::CMP, &a::IZY, 5 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "???", &a::NOP, &a::IMP, 4 },            { "CMP", &a::CMP, &a::ZPX, 4 },            { "DEC", &a::DEC, &a::ZPX, 6 },            { "???", &a::XXX, &a::IMP, 6 },            { "CLD", &a::CLD, &a::IMP, 2 },            { "CMP", &a::CMP, &a::ABY, 4 },            { "NOP", &a::NOP, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 7 },            { "???", &a::NOP, &a::IMP, 4 },            { "CMP", &a::CMP, &a::ABX, 4 },            { "DEC", &a::DEC, &a::ABX, 7 },            { "???", &a::XXX, &a::IMP, 7 },          

									{ "CPX", &a::CPX, &a::IMM, 2 },            { "SBC", &a::SBC, &a::IZX, 6 },            { "???", &a::NOP, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "CPX", &a::CPX, &a::ZP0, 3 },            { "SBC", &a::SBC, &a::ZP0, 3 },            { "INC", &a::INC, &a::ZP0, 5 },            { "???", &a::XXX, &a::IMP, 5 },            { "INX", &a::INX, &a::IMP, 2 },            { "SBC", &a::SBC, &a::IMM, 2 },            { "NOP", &a::NOP, &a::IMP, 2 },            { "???", &a::SBC, &a::IMP, 2 },            { "CPX", &a::CPX, &a::ABS, 4 },            { "SBC", &a::SBC, &a::ABS, 4 },            { "INC", &a::INC, &a::ABS, 6 },            { "???", &a::XXX, &a::IMP, 6 },          

									{ "BEQ", &a::BEQ, &a::REL, 2 },            { "SBC", &a::SBC, &a::IZY, 5 },            { "???", &a::XXX, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 8 },            { "???", &a::NOP, &a::IMP, 4 },            { "SBC", &a::SBC, &a::ZPX, 4 },            { "INC", &a::INC, &a::ZPX, 6 },            { "???", &a::XXX, &a::IMP, 6 },            { "SED", &a::SED, &a::IMP, 2 },            { "SBC", &a::SBC, &a::ABY, 4 },            { "NOP", &a::NOP, &a::IMP, 2 },            { "???", &a::XXX, &a::IMP, 7 },            { "???", &a::NOP, &a::IMP, 4 },            { "SBC", &a::SBC, &a::ABX, 4 },            { "INC", &a::INC, &a::ABX, 7 },            { "???", &a::XXX, &a::IMP, 7 },
*/
								}

	return cpu
}
/*
{
	{ "BRK", &a::BRK, &a::IMM, 7 },{ "ORA", &a::ORA, &a::IZX, 6 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 3 },{ "ORA", &a::ORA, &a::ZP0, 3 },{ "ASL", &a::ASL, &a::ZP0, 5 },{ "???", &a::XXX, &a::IMP, 5 },{ "PHP", &a::PHP, &a::IMP, 3 },{ "ORA", &a::ORA, &a::IMM, 2 },{ "ASL", &a::ASL, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::NOP, &a::IMP, 4 },{ "ORA", &a::ORA, &a::ABS, 4 },{ "ASL", &a::ASL, &a::ABS, 6 },{ "???", &a::XXX, &a::IMP, 6 },
	{ "BPL", &a::BPL, &a::REL, 2 },{ "ORA", &a::ORA, &a::IZY, 5 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 4 },{ "ORA", &a::ORA, &a::ZPX, 4 },{ "ASL", &a::ASL, &a::ZPX, 6 },{ "???", &a::XXX, &a::IMP, 6 },{ "CLC", &a::CLC, &a::IMP, 2 },{ "ORA", &a::ORA, &a::ABY, 4 },{ "???", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 7 },{ "???", &a::NOP, &a::IMP, 4 },{ "ORA", &a::ORA, &a::ABX, 4 },{ "ASL", &a::ASL, &a::ABX, 7 },{ "???", &a::XXX, &a::IMP, 7 },
	{ "JSR", &a::JSR, &a::ABS, 6 },{ "AND", &a::AND, &a::IZX, 6 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "BIT", &a::BIT, &a::ZP0, 3 },{ "AND", &a::AND, &a::ZP0, 3 },{ "ROL", &a::ROL, &a::ZP0, 5 },{ "???", &a::XXX, &a::IMP, 5 },{ "PLP", &a::PLP, &a::IMP, 4 },{ "AND", &a::AND, &a::IMM, 2 },{ "ROL", &a::ROL, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 2 },{ "BIT", &a::BIT, &a::ABS, 4 },{ "AND", &a::AND, &a::ABS, 4 },{ "ROL", &a::ROL, &a::ABS, 6 },{ "???", &a::XXX, &a::IMP, 6 },
	{ "BMI", &a::BMI, &a::REL, 2 },{ "AND", &a::AND, &a::IZY, 5 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 4 },{ "AND", &a::AND, &a::ZPX, 4 },{ "ROL", &a::ROL, &a::ZPX, 6 },{ "???", &a::XXX, &a::IMP, 6 },{ "SEC", &a::SEC, &a::IMP, 2 },{ "AND", &a::AND, &a::ABY, 4 },{ "???", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 7 },{ "???", &a::NOP, &a::IMP, 4 },{ "AND", &a::AND, &a::ABX, 4 },{ "ROL", &a::ROL, &a::ABX, 7 },{ "???", &a::XXX, &a::IMP, 7 },
	{ "RTI", &a::RTI, &a::IMP, 6 },{ "EOR", &a::EOR, &a::IZX, 6 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 3 },{ "EOR", &a::EOR, &a::ZP0, 3 },{ "LSR", &a::LSR, &a::ZP0, 5 },{ "???", &a::XXX, &a::IMP, 5 },{ "PHA", &a::PHA, &a::IMP, 3 },{ "EOR", &a::EOR, &a::IMM, 2 },{ "LSR", &a::LSR, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 2 },{ "JMP", &a::JMP, &a::ABS, 3 },{ "EOR", &a::EOR, &a::ABS, 4 },{ "LSR", &a::LSR, &a::ABS, 6 },{ "???", &a::XXX, &a::IMP, 6 },
	{ "BVC", &a::BVC, &a::REL, 2 },{ "EOR", &a::EOR, &a::IZY, 5 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 4 },{ "EOR", &a::EOR, &a::ZPX, 4 },{ "LSR", &a::LSR, &a::ZPX, 6 },{ "???", &a::XXX, &a::IMP, 6 },{ "CLI", &a::CLI, &a::IMP, 2 },{ "EOR", &a::EOR, &a::ABY, 4 },{ "???", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 7 },{ "???", &a::NOP, &a::IMP, 4 },{ "EOR", &a::EOR, &a::ABX, 4 },{ "LSR", &a::LSR, &a::ABX, 7 },{ "???", &a::XXX, &a::IMP, 7 },
	{ "RTS", &a::RTS, &a::IMP, 6 },{ "ADC", &a::ADC, &a::IZX, 6 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 3 },{ "ADC", &a::ADC, &a::ZP0, 3 },{ "ROR", &a::ROR, &a::ZP0, 5 },{ "???", &a::XXX, &a::IMP, 5 },{ "PLA", &a::PLA, &a::IMP, 4 },{ "ADC", &a::ADC, &a::IMM, 2 },{ "ROR", &a::ROR, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 2 },{ "JMP", &a::JMP, &a::IND, 5 },{ "ADC", &a::ADC, &a::ABS, 4 },{ "ROR", &a::ROR, &a::ABS, 6 },{ "???", &a::XXX, &a::IMP, 6 },
	{ "BVS", &a::BVS, &a::REL, 2 },{ "ADC", &a::ADC, &a::IZY, 5 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 4 },{ "ADC", &a::ADC, &a::ZPX, 4 },{ "ROR", &a::ROR, &a::ZPX, 6 },{ "???", &a::XXX, &a::IMP, 6 },{ "SEI", &a::SEI, &a::IMP, 2 },{ "ADC", &a::ADC, &a::ABY, 4 },{ "???", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 7 },{ "???", &a::NOP, &a::IMP, 4 },{ "ADC", &a::ADC, &a::ABX, 4 },{ "ROR", &a::ROR, &a::ABX, 7 },{ "???", &a::XXX, &a::IMP, 7 },
	{ "???", &a::NOP, &a::IMP, 2 },{ "STA", &a::STA, &a::IZX, 6 },{ "???", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 6 },{ "STY", &a::STY, &a::ZP0, 3 },{ "STA", &a::STA, &a::ZP0, 3 },{ "STX", &a::STX, &a::ZP0, 3 },{ "???", &a::XXX, &a::IMP, 3 },{ "DEY", &a::DEY, &a::IMP, 2 },{ "???", &a::NOP, &a::IMP, 2 },{ "TXA", &a::TXA, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 2 },{ "STY", &a::STY, &a::ABS, 4 },{ "STA", &a::STA, &a::ABS, 4 },{ "STX", &a::STX, &a::ABS, 4 },{ "???", &a::XXX, &a::IMP, 4 },
	{ "BCC", &a::BCC, &a::REL, 2 },{ "STA", &a::STA, &a::IZY, 6 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 6 },{ "STY", &a::STY, &a::ZPX, 4 },{ "STA", &a::STA, &a::ZPX, 4 },{ "STX", &a::STX, &a::ZPY, 4 },{ "???", &a::XXX, &a::IMP, 4 },{ "TYA", &a::TYA, &a::IMP, 2 },{ "STA", &a::STA, &a::ABY, 5 },{ "TXS", &a::TXS, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 5 },{ "???", &a::NOP, &a::IMP, 5 },{ "STA", &a::STA, &a::ABX, 5 },{ "???", &a::XXX, &a::IMP, 5 },{ "???", &a::XXX, &a::IMP, 5 },
	{ "LDY", &a::LDY, &a::IMM, 2 },{ "LDA", &a::LDA, &a::IZX, 6 },{ "LDX", &a::LDX, &a::IMM, 2 },{ "???", &a::XXX, &a::IMP, 6 },{ "LDY", &a::LDY, &a::ZP0, 3 },{ "LDA", &a::LDA, &a::ZP0, 3 },{ "LDX", &a::LDX, &a::ZP0, 3 },{ "???", &a::XXX, &a::IMP, 3 },{ "TAY", &a::TAY, &a::IMP, 2 },{ "LDA", &a::LDA, &a::IMM, 2 },{ "TAX", &a::TAX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 2 },{ "LDY", &a::LDY, &a::ABS, 4 },{ "LDA", &a::LDA, &a::ABS, 4 },{ "LDX", &a::LDX, &a::ABS, 4 },{ "???", &a::XXX, &a::IMP, 4 },
	{ "BCS", &a::BCS, &a::REL, 2 },{ "LDA", &a::LDA, &a::IZY, 5 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 5 },{ "LDY", &a::LDY, &a::ZPX, 4 },{ "LDA", &a::LDA, &a::ZPX, 4 },{ "LDX", &a::LDX, &a::ZPY, 4 },{ "???", &a::XXX, &a::IMP, 4 },{ "CLV", &a::CLV, &a::IMP, 2 },{ "LDA", &a::LDA, &a::ABY, 4 },{ "TSX", &a::TSX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 4 },{ "LDY", &a::LDY, &a::ABX, 4 },{ "LDA", &a::LDA, &a::ABX, 4 },{ "LDX", &a::LDX, &a::ABY, 4 },{ "???", &a::XXX, &a::IMP, 4 },
	{ "CPY", &a::CPY, &a::IMM, 2 },{ "CMP", &a::CMP, &a::IZX, 6 },{ "???", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "CPY", &a::CPY, &a::ZP0, 3 },{ "CMP", &a::CMP, &a::ZP0, 3 },{ "DEC", &a::DEC, &a::ZP0, 5 },{ "???", &a::XXX, &a::IMP, 5 },{ "INY", &a::INY, &a::IMP, 2 },{ "CMP", &a::CMP, &a::IMM, 2 },{ "DEX", &a::DEX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 2 },{ "CPY", &a::CPY, &a::ABS, 4 },{ "CMP", &a::CMP, &a::ABS, 4 },{ "DEC", &a::DEC, &a::ABS, 6 },{ "???", &a::XXX, &a::IMP, 6 },
	{ "BNE", &a::BNE, &a::REL, 2 },{ "CMP", &a::CMP, &a::IZY, 5 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 4 },{ "CMP", &a::CMP, &a::ZPX, 4 },{ "DEC", &a::DEC, &a::ZPX, 6 },{ "???", &a::XXX, &a::IMP, 6 },{ "CLD", &a::CLD, &a::IMP, 2 },{ "CMP", &a::CMP, &a::ABY, 4 },{ "NOP", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 7 },{ "???", &a::NOP, &a::IMP, 4 },{ "CMP", &a::CMP, &a::ABX, 4 },{ "DEC", &a::DEC, &a::ABX, 7 },{ "???", &a::XXX, &a::IMP, 7 },
	{ "CPX", &a::CPX, &a::IMM, 2 },{ "SBC", &a::SBC, &a::IZX, 6 },{ "???", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "CPX", &a::CPX, &a::ZP0, 3 },{ "SBC", &a::SBC, &a::ZP0, 3 },{ "INC", &a::INC, &a::ZP0, 5 },{ "???", &a::XXX, &a::IMP, 5 },{ "INX", &a::INX, &a::IMP, 2 },{ "SBC", &a::SBC, &a::IMM, 2 },{ "NOP", &a::NOP, &a::IMP, 2 },{ "???", &a::SBC, &a::IMP, 2 },{ "CPX", &a::CPX, &a::ABS, 4 },{ "SBC", &a::SBC, &a::ABS, 4 },{ "INC", &a::INC, &a::ABS, 6 },{ "???", &a::XXX, &a::IMP, 6 },
	{ "BEQ", &a::BEQ, &a::REL, 2 },{ "SBC", &a::SBC, &a::IZY, 5 },{ "???", &a::XXX, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 8 },{ "???", &a::NOP, &a::IMP, 4 },{ "SBC", &a::SBC, &a::ZPX, 4 },{ "INC", &a::INC, &a::ZPX, 6 },{ "???", &a::XXX, &a::IMP, 6 },{ "SED", &a::SED, &a::IMP, 2 },{ "SBC", &a::SBC, &a::ABY, 4 },{ "NOP", &a::NOP, &a::IMP, 2 },{ "???", &a::XXX, &a::IMP, 7 },{ "???", &a::NOP, &a::IMP, 4 },{ "SBC", &a::SBC, &a::ABX, 4 },{ "INC", &a::INC, &a::ABX, 7 },{ "???", &a::XXX, &a::IMP, 7 },
}
*/
func (c *CPU6502) ConnectBus(bus Bus) {
	c.bus = bus
}

func (c *CPU6502) Read(addr uint16) uint8 {
	return c.bus.Read(addr, false) // readOnly = false
}

func (c *CPU6502) Write(addr uint16, data uint8) {
	c.bus.Write(addr, data)
}

func (c *CPU6502) Fetch() uint8 {
	if c.lookup[c.opcode].AddrName != "IMP" {
		c.fetched = c.Read(c.addrAbs)
	}
	return c.fetched
}

func (c *CPU6502) Clock() {
	if c.cycles == 0 {
		c.opcode = c.Read(c.pc)
		c.pc++
		c.cycles = c.lookup[c.opcode].Cycles
		additionalCycle1 := c.lookup[c.opcode].AddrMode()
		additionalCycle2 := c.lookup[c.opcode].Operate()

		c.cycles += (additionalCycle1 & additionalCycle2)
	}
	c.cycles--
}

func (c *CPU6502) Complete() bool {
	return c.cycles == 0
}

func (c *CPU6502) Reset() {
	c.a = 0
	c.x = 0
	c.y = 0
	c.stkp = 0xFD
	c.status = 0x00 | uint8(U)
	c.addrAbs = 0xFFFC
	lo := uint16(c.Read(c.addrAbs + 0))
	hi := uint16(c.Read(c.addrAbs + 1))
	c.pc = (hi << 8) | lo
	c.addrRel = 0x0000
	c.addrAbs = 0x0000
	c.fetched = 0x00
	c.cycles = 8
}

func (c *CPU6502) IRQ() {
	if !c.GetFlag(I) {
		c.Write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
		c.stkp--
		c.Write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
		c.stkp--

		c.SetFlag(B, false)
		c.SetFlag(U, true)
		c.SetFlag(I, true)
		c.Write(0x0100+uint16(c.stkp), c.status)
		c.stkp--

		c.addrAbs = 0xFFFE
		lo := uint16(c.Read(c.addrAbs + 0))
		hi := uint16(c.Read(c.addrAbs + 1))
		c.pc = (hi << 8) | lo
		c.cycles = 7

	}
}
func (c *CPU6502) NMI() {
	c.Write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
	c.stkp--
	c.Write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
	c.stkp--

	c.SetFlag(B, false)
	c.SetFlag(U, true)
	c.SetFlag(I, true)
	c.Write(0x0100+uint16(c.stkp), c.status)
	c.stkp--

	c.addrAbs = 0xFFFA
	lo := uint16(c.Read(c.addrAbs + 0))
	hi := uint16(c.Read(c.addrAbs + 1))
	c.pc = (hi << 8) | lo
	c.cycles = 8
}


