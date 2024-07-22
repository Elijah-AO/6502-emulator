import numpy as np
# TODO: Check np zeros
from cpu import cpu

class bus:
    def __init__(self):
        self.cpu = cpu()
        self.ram = [0x00] * 64 * 1024
        
    def write(self, addr, data):
        if addr >= 0x0000 and addr <= 0xFFFF:
            self.ram[addr] = data
        
    def read(self, addr, read_only=False):
        if addr >= 0x0000 and addr <= 0xFFFF:
            return self.ram[addr]
        return 0x00
    
