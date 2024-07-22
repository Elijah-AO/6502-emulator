from enum import Enum

class Flags(Enum):
        C = (1 << 0)  
        Z = (1 << 1)  
        I = (1 << 2)  
        D = (1 << 3)  
        B = (1 << 4)  
        U = (1 << 5)  
        V = (1 << 6)  
        N = (1 << 7) 
        
class Instruction():
        def __init__(self, name, addr_name, operate, addr_mode, cycles):
            self.name = name
            self.addr_name = addr_name
            self.operate = operate
            self.addr_mode = addr_mode
            self.cycles = cycles
            
        def execute(self):
            return 0
        
class cpu:
    def __init__(self):
        self.a = 0x00
        self.x = 0x00
        self.y = 0x00 
        
        self.sp = 0x00
        self.pc = 0x0000
        
        self.status = 0x00
        self.addr_abs = 0x0000
        self.addr_rel = 0x00
        self.fetched = 0x00
        self.opcode = 0x00
        self.cycles = 0
        self.lookup = [
            Instruction("BRK", "IMP",self.brk, self.imp, 7), Instruction("ORA", "IZX", self.ora, self.izx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZP0", self.nop, self.zp0, 3), Instruction("ORA", "ZP0", self.ora, self.zp0, 3), Instruction("ASL", "ZP0", self.asl, self.zp0, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("PHP", "IMP", self.php, self.imp, 3), Instruction("ORA", "IMM", self.ora, self.imm, 2), Instruction("ASL", "IMP", self.asl, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("NOP", "ABS", self.nop, self.abs, 4), Instruction("ORA", "ABS", self.ora, self.abs, 4), Instruction("ASL", "ABS", self.asl, self.abs, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6),
            Instruction("BPL", "REL", self.bpl, self.rel, 2), Instruction("ORA", "IZY", self.ora, self.izy, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZPX", self.nop, self.zpx, 4), Instruction("ORA", "ZPX", self.ora, self.zpx, 4), Instruction("ASL", "ZPX", self.asl, self.zpx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("CLC", "IMP", self.clc, self.imp, 2), Instruction("ORA", "ABY", self.ora, self.aby, 4), Instruction("NOP", "IMP", self.nop, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 7), Instruction("NOP", "ABX", self.nop, self.abx, 4), Instruction("ORA", "ABX", self.ora, self.abx, 4), Instruction("ASL", "ABX", self.asl, self.abx, 7), Instruction("XXX", "IMP", self.xxx, self.imp, 7),
            Instruction("JSR", "ABS", self.jsr, self.abs, 6), Instruction("AND", "IZX", self.and_, self.izx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("BIT", "ZP0", self.bit, self.zp0, 3), Instruction("AND", "ZP0", self.and_, self.zp0, 3), Instruction("ROL", "ZP0", self.rol, self.zp0, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("PLP", "IMP", self.plp, self.imp, 4), Instruction("AND", "IMM", self.and_, self.imm, 2), Instruction("ROL", "IMP", self.rol, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("BIT", "ABS", self.bit, self.abs, 4), Instruction("AND", "ABS", self.and_, self.abs, 4), Instruction("ROL", "ABS", self.rol, self.abs, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6),
            Instruction("BMI", "REL", self.bmi, self.rel, 2), Instruction("AND", "IZY", self.and_, self.izy, 3), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZPX", self.nop, self.zpx, 4), Instruction("AND", "ZPX", self.and_, self.zpx, 4), Instruction("ROL", "ZPX", self.rol, self.zpx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("SEC", "IMP", self.sec, self.imp, 2), Instruction("AND", "ABY", self.and_, self.aby, 4), Instruction("NOP", "IMP", self.nop, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 7), Instruction("NOP", "ABX", self.nop, self.abx, 4), Instruction("AND", "ABX", self.and_, self.abx, 4), Instruction("ROL", "ABX", self.rol, self.abx, 7), Instruction("XXX", "IMP", self.xxx, self.imp, 7),
            Instruction("RTI", "IMP", self.rti, self.imp, 6), Instruction("EOR", "IZX", self.eor, self.izx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZP0", self.nop, self.zp0, 3), Instruction("EOR", "ZP0", self.eor, self.zp0, 3), Instruction("LSR", "ZP0", self.lsr, self.zp0, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("PHA", "IMP", self.pha, self.imp, 3), Instruction("EOR", "IMM", self.eor, self.imm, 2), Instruction("LSR", "IMP", self.lsr, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("JMP", "ABS", self.jmp, self.abs, 3), Instruction("EOR", "ABS", self.eor, self.abs, 4), Instruction("LSR", "ABS", self.lsr, self.abs, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6),
            Instruction("BVC", "REL", self.bvc, self.rel, 2), Instruction("EOR", "IZY", self.eor, self.izy, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZPX", self.nop, self.zpx, 4), Instruction("EOR", "ZPX", self.eor, self.zpx, 4), Instruction("LSR", "ZPX", self.lsr, self.zpx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("CLI", "IMP", self.cli, self.imp, 2), Instruction("EOR", "ABY", self.eor, self.aby, 4), Instruction("NOP", "IMP", self.nop, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 7), Instruction("NOP", "ABX", self.nop, self.abx, 4), Instruction("EOR", "ABX", self.eor, self.abx, 4), Instruction("LSR", "ABX", self.lsr, self.abx, 7), Instruction("XXX", "IMP", self.xxx, self.imp, 7),
            Instruction("RTS", "IMP", self.rts, self.imp, 6), Instruction("ADC", "IZX", self.adc, self.izx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 2),  Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZP0", self.nop, self.zp0, 3), Instruction("ADC", "ZP0", self.adc, self.zp0, 3), Instruction("ROR", "ZP0", self.ror, self.zp0, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("PLA", "IMP", self.pla, self.imp, 4), Instruction("ADC", "IMM", self.adc, self.imm, 2), Instruction("ROR", "IMP", self.ror, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("JMP", "IND", self.jmp, self.ind, 5), Instruction("ADC", "ABS", self.adc, self.abs, 4), Instruction("ROR", "ABS", self.ror, self.abs, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6),
            Instruction("BVS", "REL", self.bvs, self.rel, 2), Instruction("ADC", "IZY", self.adc, self.izy, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZPX", self.nop, self.zpx, 4), Instruction("ADC", "ZPX", self.adc, self.zpx, 4), Instruction("ROR", "ZPX", self.ror, self.zpx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("SEI", "IMP", self.sei, self.imp, 2), Instruction("ADC", "ABY", self.adc, self.aby, 4), Instruction("NOP", "IMP", self.nop, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 7), Instruction("NOP", "ABX", self.nop, self.abx, 4), Instruction("ADC", "ABX", self.adc, self.abx, 4), Instruction("ROR", "ABX", self.ror, self.abx, 7), Instruction("XXX", "IMP", self.xxx, self.imp, 7),
            Instruction("NOP", "IMM", self.nop, self.imm, 2), Instruction("STA", "IZX", self.sta, self.izx, 6), Instruction("NOP", "IMM", self.nop, self.imm, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("STY", "ZP0", self.sty, self.zp0, 3), Instruction("STA", "ZP0", self.sta, self.zp0, 3), Instruction("STX", "ZP0", self.stx, self.zp0, 3), Instruction("XXX", "IMP", self.xxx, self.imp, 3), Instruction("DEY", "IMP", self.dey, self.imp, 2), Instruction("NOP", "IMM", self.nop, self.imm, 2), Instruction("TXA", "IMP", self.txa, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("STY", "ABS", self.sty, self.abs, 4), Instruction("STA", "ABS", self.sta, self.abs, 4), Instruction("STX", "ABS", self.stx, self.abs, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 4), 
            Instruction("BCC", "REL", self.bcc, self.rel, 2), Instruction("STA", "IZY", self.sta, self.izy, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("STY", "ZPX", self.sty, self.zpx, 4), Instruction("STA", "ZPX", self.sta, self.zpx, 4), Instruction("STX", "ZPY", self.stx, self.zpy, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 4), Instruction("TYA", "IMP", self.tya, self.imp, 2), Instruction("STA", "ABY", self.sta, self.aby, 5), Instruction("TXS", "IMP", self.txs, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("STA", "ABX", self.sta, self.abx, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5),
            Instruction("LDY", "IMM", self.ldy, self.imm, 2), Instruction("LDA", "IZX", self.lda, self.izx, 6), Instruction("LDX", "IMM", self.ldx, self.imm, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("LDY", "ZP0", self.ldy, self.zp0, 3), Instruction("LDA", "ZP0", self.lda, self.zp0, 3), Instruction("LDX", "ZP0", self.ldx, self.zp0, 3), Instruction("XXX", "IMP", self.xxx, self.imp, 3), Instruction("TAY", "IMP", self.tay, self.imp, 2), Instruction("LDA", "IMM", self.lda, self.imm, 2), Instruction("TAX", "IMP", self.tax, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("LDY", "ABS", self.ldy, self.abs, 4), Instruction("LDA", "ABS", self.lda, self.abs, 4), Instruction("LDX", "ABS", self.ldx, self.abs, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 4),
            Instruction("BCS", "REL", self.bcs, self.rel, 2), Instruction("LDA", "IZY", self.lda, self.izy, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("LDY", "ZPX", self.ldy, self.zpx, 4), Instruction("LDA", "ZPX", self.lda, self.zpx, 4), Instruction("LDX", "ZPY", self.ldx, self.zpy, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 4), Instruction("CLV", "IMP", self.clv, self.imp, 2), Instruction("LDA", "ABY", self.lda, self.aby, 4), Instruction("TSX", "IMP", self.tsx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 4), Instruction("LDY", "ABX", self.ldy, self.abx, 4), Instruction("LDA", "ABX", self.lda, self.abx, 4), Instruction("LDX", "ABY", self.ldx, self.aby, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 4),
            Instruction("CPY", "IMM", self.cpy, self.imm, 2), Instruction("CMP", "IZX", self.cmp, self.izx, 6), Instruction("NOP", "IMM", self.nop, self.imm, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("CPY", "ZP0", self.cpy, self.zp0, 3), Instruction("CMP", "ZP0", self.cmp, self.zp0, 3), Instruction("DEC", "ZP0", self.dec, self.zp0, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("INY", "IMP", self.iny, self.imp, 2), Instruction("CMP", "IMM", self.cmp, self.imm, 2), Instruction("DEX", "IMP", self.dex, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("CPY", "ABS", self.cpy, self.abs, 4), Instruction("CMP", "ABS", self.cmp, self.abs, 4), Instruction("DEC", "ABS", self.dec, self.abs, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6),
            Instruction("BNE", "REL", self.bne, self.rel, 2), Instruction("CMP", "IZY", self.cmp, self.izy, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZPX", self.nop, self.zpx, 4), Instruction("CMP", "ZPX", self.cmp, self.zpx, 4), Instruction("DEC", "ZPX", self.dec, self.zpx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("CLD", "IMP", self.cld, self.imp, 2), Instruction("CMP", "ABY", self.cmp, self.aby, 4), Instruction("NOP", "IMP", self.nop, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 7), Instruction("NOP", "ABX", self.nop, self.abx, 4), Instruction("CMP", "ABX", self.cmp, self.abx, 4), Instruction("DEC", "ABX", self.dec, self.abx, 7), Instruction("XXX", "IMP", self.xxx, self.imp, 7),
            Instruction("CPX", "IMM", self.cpx, self.imm, 2), Instruction("SBC", "IZX", self.sbc, self.izx, 6), Instruction("NOP", "IMM", self.nop, self.imm, 2),  Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("CPX", "ZP0", self.cpx, self.zp0, 3), Instruction("SBC", "ZP0", self.sbc, self.zp0, 3), Instruction("INC", "ZP0", self.inc, self.zp0, 5), Instruction("XXX", "IMP", self.xxx, self.imp, 5), Instruction("INX", "IMP", self.inx, self.imp, 2), Instruction("SBC", "IMM", self.sbc, self.imm, 2), Instruction("NOP", "IMP", self.nop, self.imp, 2), Instruction("SBC", "IMP", self.sbc, self.imp, 2), Instruction("CPX", "ABS", self.cpx, self.abs, 4), Instruction("SBC", "ABS", self.sbc, self.abs, 4), Instruction("INC", "ABS", self.inc, self.abs, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6),
            Instruction("BEQ", "REL", self.beq, self.rel, 2), Instruction("SBC", "IZY", self.sbc, self.izy, 4), Instruction("XXX", "IMP", self.xxx, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 8), Instruction("NOP", "ZPX", self.nop, self.zpx, 4), Instruction("SBC", "ZPX", self.sbc, self.zpx, 4), Instruction("INC", "ZPX", self.inc, self.zpx, 6), Instruction("XXX", "IMP", self.xxx, self.imp, 6), Instruction("SED", "IMP", self.sed, self.imp, 2), Instruction("SBC", "ABY", self.sbc, self.aby, 4), Instruction("NOP", "IMP", self.nop, self.imp, 2), Instruction("XXX", "IMP", self.xxx, self.imp, 7), Instruction("NOP", "ABX", self.nop, self.abx, 4), Instruction("SBC", "ABX", self.sbc, self.abx, 4), Instruction("INC", "ABX", self.inc, self.abx, 7), Instruction("XXX", "IMP", self.xxx, self.imp, 7)
            
            
        ]
    

    def read(self, addr):
        return self.bus.read(addr, False)
    
    def write(self, addr, data):
        self.bus.write(addr, data)
        
    def connect_bus(self, bus):
        self.bus = bus
        
    def get_flag(self, flag):
        return (self.status & flag.value) != 0

    def set_flag(self, flag, value):
        if value:
            self.status |= flag.value
        else:
            self.status &= ~flag.value
    
    def clock(self):
        if self.cycles == 0:
            self.opcode = self.read(self.pc)
            self.pc += 1
            
            self.cycles = self.lookup[self.opcode].cycles
            add_cycles_1 = self.lookup[self.opcode].addr_mode()
            add_cycles_2 = self.lookup[self.opcode].operate()
            self.cycles += (add_cycles_1 & add_cycles_2)
            
        self.cycles -= 1
    
    def reset(self):
        self.a = 0
        self.x = 0
        self.y = 0
        self.sp = 0xFD
        self.status = 0x00 | Flags.U.value
        self.addr_abs = 0xFFFC
        lo = self.read(self.addr_abs + 0)
        hi = self.read(self.addr_abs + 1)
        self.pc = (hi << 8) | lo
        self.addr_abs = 0x0000
        self.addr_rel = 0x0000
        self.fetched = 0x00
        self.cycles = 8
    
    def irq(self):
        if self.get_flag(Flags.I) == 0:
            self.write(0x0100 + self.sp, (self.pc >> 8) & 0x00FF)
            self.sp -= 1
            self.write(0x0100 + self.sp, self.pc & 0x00FF)
            self.sp -= 1
            
            self.set_flag(Flags.B, 0)
            self.set_flag(Flags.U, 1)
            self.set_flag(Flags.I, 1)
            self.write(0x0100 + self.sp, self.status)
            self.sp -= 1
            
            self.addr_abs = 0xFFFE
            lo = self.read(self.addr_abs + 0)
            hi = self.read(self.addr_abs + 1)
            self.pc = (hi << 8) | lo
            self.cycles = 7
    
    def nmi(self):
        self.write(0x0100 + self.sp, (self.pc >> 8) & 0x00FF)
        self.sp -= 1
        self.write(0x0100 + self.sp, self.pc & 0x00FF)
        self.sp -= 1
        
        self.set_flag(Flags.B, 0)
        self.set_flag(Flags.U, 1)
        self.set_flag(Flags.I, 1)
        self.write(0x0100 + self.sp, self.status)
        self.sp -= 1
        
        self.addr_abs = 0xFFFA
        lo = self.read(self.addr_abs + 0)
        hi = self.read(self.addr_abs + 1)
        self.pc = (hi << 8) | lo
        self.cycles = 8
    
    def fetch(self):
        if self.lookup[self.opcode].addr_name != "IMP":
            self.fetched = self.read(self.addr_abs)
        return self.fetched

    def complete(self):
        self.cycles = 0
        
    # Addressing Modes
    def imp(self):
        self.fetched = self.a
        return 0

    def imm(self):
        self.addr_abs = self.pc
        self.pc += 1
        return 0

    def zp0(self):
        self.addr_abs = self.read(self.pc)
        self.pc += 1
        self.addr_abs &= 0x00FF
        return 0

    def zpx(self):
        self.addr_abs = (self.read(self.pc) + self.x)
        self.pc += 1
        self.addr_abs &= 0x00FF
        return 0

    def zpy(self):
        self.addr_abs = (self.read(self.pc) + self.y)
        self.pc += 1
        self.addr_abs &= 0x00FF
        return 0

    def rel(self):
        self.addr_rel = self.read(self.pc)
        self.pc += 1
        if (self.addr_rel & 0x80) != 0:
            self.addr_rel |= 0xFF00
        return 0

    def abs(self):
        lo = self.read(self.pc)
        self.pc += 1
        hi = self.read(self.pc)
        self.pc += 1
        self.addr_abs = (hi << 8) | lo
        return 0

    def abx(self):
        lo = self.read(self.pc)
        self.pc += 1
        hi = self.read(self.pc)
        self.pc += 1
        self.addr_abs = (hi << 8) | lo
        self.addr_abs += self.x
        if (self.addr_abs & 0xFF00) != (hi << 8):
            return 1
        return 0

    def aby(self):
        lo = self.read(self.pc)
        self.pc += 1
        hi = self.read(self.pc)
        self.pc += 1
        self.addr_abs = (hi << 8) | lo
        self.addr_abs += self.y
        if (self.addr_abs & 0xFF00) != (hi << 8):
            return 1
        return 0

    def ind(self):
        ptr_lo = self.read(self.pc)
        self.pc += 1
        ptr_hi = self.read(self.pc)
        self.pc += 1
        ptr = (ptr_hi << 8) | ptr_lo
        
        if ptr_lo == 0x00FF:
            self.addr_abs = (self.read(ptr & 0xFF00) << 8) | self.read(ptr)
        else:
            self.addr_abs = (self.read(ptr + 1) << 8) | self.read(ptr)
        return 0

    def izx(self):
        t = self.read(self.pc)
        self.pc += 1
        lo = self.read((t + self.x) & 0x00FF)
        hi = self.read((t + self.x + 1) & 0x00FF)
        self.addr_abs = (hi << 8) | lo
        return 0

    def izy(self):
        t = self.read(self.pc)
        self.pc += 1
        lo = self.read(t & 0x00FF)
        hi = self.read((t + 1) & 0x00FF)
        self.addr_abs = (hi << 8) | lo
        self.addr_abs += self.y
        if (self.addr_abs & 0xFF00) != (hi << 8):
            return 1
        return 0
        
        
    def compare(self, reg):
        self.fetch()
        match reg:
            case 'a':
                temp = self.a - self.fetched
                self.set_flag(Flags.C, self.a >= self.fetched)
                
            case 'x':
                temp = self.x - self.fetched
                self.set_flag(Flags.C, self.x >= self.fetched)
                
            case 'y':
                temp = self.y - self.fetched
                self.set_flag(Flags.C, self.y >= self.fetched)
                
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0x0000)
        self.set_flag(Flags.N, temp & 0x0080)
    
        
    def branch(self):
        self.cycles += 1
        self.addr_abs = (self.pc + self.addr_rel) & 0xFFFF
        
        if (self.addr_abs & 0xFF00) != (self.pc & 0xFF00):
            self.cycles += 1
        self.pc = self.addr_abs
        
    # TODO: Implement illegal opcodes
    # Instructions

    def adc(self):
        self.fetch()
        temp = self.a + self.fetched + self.get_flag(Flags.C)
        self.set_flag(Flags.C, temp > 255)
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0)
        self.set_flag(Flags.V, (~self.a ^ self.fetched) & (self.a ^ temp) & 0x0080)
        self.set_flag(Flags.N, temp & 0x80)
        self.a = temp & 0x00FF
        return 1

    def and_(self):
        self.fetch()
        self.a = self.a & self.fetched
        self.set_flag(Flags.Z, self.a == 0x00)
        self.set_flag(Flags.N, self.a & 0x80)
        return 1

    def asl(self):
        self.fetch()
        temp = self.fetched << 1
        self.set_flag(Flags.C, (temp & 0xFF00) > 0)
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0x00)
        self.set_flag(Flags.N, temp & 0x80)
        if self.lookup[self.opcode].addr_name == "IMP":
            self.a = temp & 0x00FF 
        else:
            self.write(self.addr_abs, temp & 0x00FF)
        return 0

    def bcc(self):
        if self.get_flag(Flags.C) == 0:
            self.branch()
        return 0

    def bcs(self):
        if self.get_flag(Flags.C) == 1:
            self.branch()
        return 0

    def beq(self):
        if self.get_flag(Flags.Z) == 1:
            self.branch()
        return 0

    def bit(self):
        self.fetch()
        temp = self.a & self.fetched
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0x00)
        self.set_flag(Flags.N, self.fetched & (1 << 7))
        self.set_flag(Flags.V, self.fetched & (1 << 6))
        return 0

    def bmi(self):
        if self.get_flag(Flags.N) == 1:
            self.branch()
        return 0

    def bne(self):
        if self.get_flag(Flags.Z) == 0:
            self.branch()
        return 0

    def bpl(self):
        if self.get_flag(Flags.N) == 0:
            self.branch()
        return 0

    def brk(self):
        self.pc += 1
        self.set_flag(Flags.I, 1)
        self.write(0x0100 + self.sp, (self.pc >> 8) & 0x00FF)
        self.sp -= 1
        self.write(0x0100 + self.sp, self.pc & 0x00FF)
        self.sp -= 1
        self.set_flag(Flags.B, 1)
        self.write(0x0100 + self.sp, self.status)
        self.sp -= 1
        self.set_flag(Flags.B, 0)
        self.pc = self.read(0xFFFE) | (self.read(0xFFFF) << 8)
        return 0

    def bvc(self):
        if self.get_flag(Flags.V) == 0:
            self.branch()
        return 0

    def bvs(self):
        if self.get_flag(Flags.V) == 1:
            self.branch()
        return 0

    def clc(self):
        self.set_flag(Flags.C, 0)
        return 0

    def cld(self):
        self.set_flag(Flags.D, 0)
        return 0

    def cli(self):
        self.set_flag(Flags.I, 0)
        return 0

    def clv(self):
        self.set_flag(Flags.V, 0)
        return 0

    def cmp(self):
        self.compare("a")
        return 1

    def cpx(self):
        self.compare("x")
        return 0

    def cpy(self):
        self.compare("y")
        return 0

    def dec(self):
        self.fetch()
        temp = self.fetched - 1
        self.write(self.addr_abs, temp & 0x00FF)
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0x0000)
        self.set_flag(Flags.N, temp & 0x80)
        return 0

    def dex(self):
        self.x -= 1
        self.set_flag(Flags.Z, self.x == 0x00)
        self.set_flag(Flags.N, self.x & 0x80)
        return 0

    def dey(self):
        self.y -= 1
        self.set_flag(Flags.Z, self.y == 0x00)
        self.set_flag(Flags.N, self.y & 0x80)
        return 0

    def eor(self):
        self.fetch()
        self.a = self.a ^ self.fetched
        self.set_flag(Flags.Z, self.a == 0x00)
        self.set_flag(Flags.N, self.a & 0x80)
        return 1

    def inc(self):
        self.fetch()
        temp = self.fetched + 1
        self.write(self.addr_abs, temp & 0x00FF)
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0x0000)
        self.set_flag(Flags.N, temp & 0x80)
        return 0

    def inx(self):
        self.x += 1
        self.set_flag(Flags.Z, self.x == 0x00)
        self.set_flag(Flags.N, self.x & 0x80)
        return 0

    def iny(self):
        self.y += 1
        self.set_flag(Flags.Z, self.y == 0x00)
        self.set_flag(Flags.N, self.y & 0x80)
        return 0

    def jmp(self):
        self.pc = self.addr_abs
        return 0

    def jsr(self):
        self.pc -= 1
        self.write(0x0100 + self.sp, (self.pc >> 8) & 0x00FF)
        self.sp -= 1
        self.write(0x0100 + self.sp, self.pc & 0x00FF)
        self.sp -= 1
        self.pc = self.addr_abs
        return 0

    def lda(self):
        self.fetch()
        self.a = self.fetched
        self.set_flag(Flags.Z, self.a == 0x00)
        self.set_flag(Flags.N, self.a & 0x80)
        return 1

    def ldx(self):
        self.fetch()
        self.x = self.fetched
        self.set_flag(Flags.Z, self.x == 0x00)
        self.set_flag(Flags.N, self.x & 0x80)
        return 1

    def ldy(self):
        self.fetch()
        self.y = self.fetched
        self.set_flag(Flags.Z, self.y == 0x00)
        self.set_flag(Flags.N, self.y & 0x80)
        return 1

    def lsr(self):
        self.fetch()
        self.set_flag(Flags.C, self.fetched & 0x0001)
        temp = self.fetched >> 1
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0x0000)
        self.set_flag(Flags.N, temp & 0x80)
        if self.lookup[self.opcode].addr_name == "IMP":
            self.a = temp & 0x00FF
        else:
            self.write(self.addr_abs, temp & 0x00FF)
        return 0

    #TODO: Implement NOP fully
    def nop(self):
        match self.opcode:
            case 0x1C, 0x3C, 0x5C, 0x7C, 0xDC, 0xFC:
                return 1
        return 0

    def ora(self):
        self.fetch()
        self.a = self.a | self.fetched
        self.set_flag(Flags.Z, self.a == 0x00)
        self.set_flag(Flags.N, self.a & 0x80)
        return 1

    def pha(self):
        self.write(0x0100 + self.sp, self.a)
        self.sp -= 1
        return 0

    def php(self):
        self.write(0x0100 + self.sp, self.status | Flags.B | Flags.U)
        self.set_flag(Flags.B, 0)
        self.set_flag(Flags.U, 0)
        self.sp -= 1
        return 0

    def pla(self):
        self.sp += 1
        self.a = self.read(0x0100 + self.sp)
        self.set_flag(Flags.Z, self.a == 0x00)
        self.set_flag(Flags.N, self.a & 0x80)
        return 0

    def plp(self):
        self.sp += 1
        self.status = self.read(0x0100 + self.sp)
        self.set_flag(Flags.U, 1)
        return 0

    def rol(self):
        self.fetch()
        temp = (self.fetched << 1) | self.get_flag(Flags.C)
        self.set_flag(Flags.C, temp & 0xFF00)
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0x0000)
        self.set_flag(Flags.N, temp & 0x80)
        if self.lookup[self.opcode].addr_name == "IMP":
            self.a = temp & 0x00FF
        else:
            self.write(self.addr_abs, temp & 0x00FF)
        return 0

    def ror(self):
        self.fetch()
        temp = (self.get_flag(Flags.C) << 7) | (self.fetched >> 1)
        self.set_flag(Flags.C, self.fetched & 0x0001)
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0x00)
        self.set_flag(Flags.N, temp & 0x80)
        if self.lookup[self.opcode].addr_name == "IMP":
            self.a = temp & 0x00FF
        else:
            self.write(self.addr_abs, temp & 0x00FF)
        return 0

    def rti(self):
        self.sp += 1
        self.status = self.read(0x0100 + self.sp)
        self.status &= ~Flags.B
        self.status &= ~Flags.U
        self.sp += 1
        self.pc = self.read(0x0100 + self.sp)
        self.sp += 1
        self.pc |= self.read(0x0100 + self.sp) << 8    
        return 0

    def rts(self):
        self.sp += 1
        self.pc = self.read(0x0100 + self.sp)
        self.sp += 1
        self.pc |= self.read(0x0100 + self.sp) << 8
        self.pc += 1
        return 0

    def sbc(self):
        self.fetch()
        value = self.fetched ^ 0x00FF
        temp = self.a + value + self.get_flag(Flags.C)
        self.set_flag(Flags.C, temp & 0xFF00)
        self.set_flag(Flags.Z, (temp & 0x00FF) == 0)
        self.set_flag(Flags.V, (temp ^ self.a) & (temp ^ value) & 0x0080)
        self.set_flag(Flags.N, temp & 0x80)
        self.a = temp & 0x00FF
        return 1

    def sec(self):
        self.set_flag(Flags.C, 1)
        return 0

    def sed(self):
        self.set_flag(Flags.D, 1)
        return 0

    def sei(self):
        self.set_flag(Flags.I, 1)
        return 0

    def sta(self):
        self.write(self.addr_abs, self.a)
        return 0

    def stx(self):
        self.write(self.addr_abs, self.x)
        return 0

    def sty(self):
        self.write(self.addr_abs, self.y)
        return 0

    def tax(self):
        self.x = self.a
        self.set_flag(Flags.Z, self.x == 0x00)
        self.set_flag(Flags.N, self.x & 0x80)
        return 0

    def tay(self):
        self.y = self.a
        self.set_flag(Flags.Z, self.y == 0x00)
        self.set_flag(Flags.N, self.y & 0x80)
        return 0

    def tsx(self):
        self.x = self.sp
        self.set_flag(Flags.Z, self.x == 0x00)
        self.set_flag(Flags.N, self.x & 0x80)
        return 0

    def txa(self):
        self.a = self.x
        self.set_flag(Flags.Z, self.a == 0x00)
        self.set_flag(Flags.N, self.a & 0x80)
        return 0

    def txs(self):
        self.sp = self.x
        return 0

    def tya(self):
        self.a = self.y
        self.set_flag(Flags.Z, self.a == 0x00)
        self.set_flag(Flags.N, self.a & 0x80)
        return 0

    def xxx(self):
        return 0

    def disassemble(self, start, stop):
        addr = start
        map_lines = {}

        while addr <= stop:
            line_addr = addr
            opcode = self.bus.read(addr, True)
            instruction = self.lookup[opcode]
            inst = instruction.name
            name = instruction.addr_name
            line = f"${addr:04X}: {inst:3} "
            addr += 1
            line += f"{opcode:02X} "

            match name:
                case "IMP":
                    line += "{IMP}"
                case "IMM":
                    value = self.bus.read(addr, True)
                    addr += 1
                    line += f"${value:02X} {{IMM}}"
                case "ZP0":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = 0x00
                    line += f"${lo:02X} {{ZP0}}"
                case "ZPX":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = 0x00
                    line += f"${lo:02X} X {{ZPX}}"
                case "ZPY":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = 0x00
                    line += f"${lo:02X} Y {{ZPY}}"
                case "IZX":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = 0x00
                    line += f"(${lo:02X}, X) {{IZX}}"
                case "IZY":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = 0x00
                    line += f"(${lo:02X}), Y {{IZY}}"
                case "ABS":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = self.bus.read(addr, True)
                    addr += 1
                    line += f"${hi<<8|lo:04X} {{ABS}}"
                case "ABX":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = self.bus.read(addr, True)
                    addr += 1
                    line += f"${hi<<8|lo:04X} X {{ABX}}"
                case "ABY":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = self.bus.read(addr, True)
                    addr += 1
                    line += f"${hi<<8|lo:04X} Y {{ABY}}"
                case "IND":
                    lo = self.bus.read(addr, True)
                    addr += 1
                    hi = self.bus.read(addr, True)
                    addr += 1
                    line += f"(${hi<<8|lo:04X}) {{IND}}"
                case "REL":
                    value = self.bus.read(addr, True)
                    addr += 1
                    line += f"${value:02X} [$ {addr+value:04X}] {{REL}}"
                case _:
                    line += "???"

            map_lines[line_addr] = line

        return map_lines
        
            
    def load_program(self, instructions, offset):
        for i, instruction in enumerate(instructions):
            converted = int(instruction, 16)  # Convert hex string to integer
            self.write(offset + i, converted)  # Write to memory

