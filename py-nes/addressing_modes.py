# Addressing modes for the 6502 CPU
from cpu import cpu

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

cpu.imp = imp
cpu.imm = imm
cpu.zp0 = zp0
cpu.zpx = zpx
cpu.zpy = zpy
cpu.rel = rel
cpu.abs = abs
cpu.abx = abx
cpu.aby = aby
cpu.ind = ind
cpu.izx = izx
cpu.izy = izy



