import pygame
import numpy as np
from time import sleep

from bus import bus
from cpu import cpu, Flags

def display_memory(bus, start, num_rows, num_cols):
    memory_lines = []
    for i in range(num_rows):
        addr = start + i * num_cols
        line = f"${addr:04X} "
        for j in range(num_cols):
            data = bus.read(addr + j, True)
            line += f"{data:02X} "
        memory_lines.append(line)
    return memory_lines

def display_cpu(cpu):
    cpu_lines = [
        f"C: {cpu.get_flag(Flags.C):1d} Z: {cpu.get_flag(Flags.Z):1d}  I: {cpu.get_flag(Flags.I):1d}  D: {cpu.get_flag(Flags.D):1d}  B: {cpu.get_flag(Flags.B):1d}  U: {cpu.get_flag(Flags.U):1d}  V: {cpu.get_flag(Flags.V):1d}  N: {cpu.get_flag(Flags.N):1d}",
        "",
        f"A: ${cpu.a:02X}  [{cpu.a}]",
        f"X: ${cpu.x:02X}  [{cpu.x}]",
        f"Y: ${cpu.y:02X}  [{cpu.y}]\n",
        f"SP: ${cpu.sp:04X}\n",
        f"PC: ${cpu.pc:04X}",
        "",
        f"Status: ${cpu.status:04X}",
        f"Fetched: ${cpu.fetched:04X}",
        f"Addr abs: ${cpu.addr_abs:04X}",
        f"Addr rel: ${cpu.addr_rel:04X}",
        f"Opcode: ${cpu.opcode :02X}",
        f"Cycles: {cpu.cycles}",
        
        
    ]
    
    return cpu_lines

def display_instructions(cpu, start, lines):
    instruction_lines = []
    disassembly = cpu.disassemble(0x0000, 0xFFFF) 
    curr = start
    for i in range(lines):
        if curr in disassembly:
            instruction_lines.append(disassembly[curr])
        curr += 1
    return instruction_lines

def main():
    b = bus()
    c = cpu()
    c.connect_bus(b)
    pygame.init()

    width, height = 1400, 850
    screen = pygame.display.set_mode((width, height), pygame.RESIZABLE, vsync=1)
    pygame.display.set_caption("NES Emulator")

    font = pygame.font.Font("py-nes/Courier Prime.ttf", 24) 

    c.write(0xFFFC, 0x00)
    c.write(0xFFFD, 0x80)
    instructions = ["A2", "0A", "8E", "00", "00", "A2", "03", "8E", "01", "00", "AC", "00", "00", "A9", "00", "18", "6D", "01", "00", "88", "D0", "FA", "8D", "02", "00", "EA", "EA", "EA"]
    c.load_program(instructions, 0x8000)
    c.reset()
    running = True
    while running:
        screen.fill((0, 0, 255))
        
        if pygame.key.get_pressed()[pygame.K_i]:
            c.irq()
        elif pygame.key.get_pressed()[pygame.K_n]:
            c.nmi()
        elif pygame.key.get_pressed()[pygame.K_SPACE]:
            c.clock()
        elif pygame.key.get_pressed()[pygame.K_r]:
            c.reset()
        elif pygame.key.get_pressed()[pygame.K_z]:
            c.set_flag(Flags.Z, c.get_flag(Flags.Z) ^ 1)
        sleep(0.1)
        
        memory_lines = display_memory(b, 0x0000, 15, 16)
        y_offset = 0
        for line in memory_lines:
            memory_text = font.render(line, True, (255, 255, 255))
            screen.blit(memory_text, (10, 10 + y_offset))
            y_offset += 25
        memory_lines = display_memory(b, 0x8000, 15, 16)
        y_offset += 25
        for line in memory_lines:
            memory_text = font.render(line, True, (255, 255, 255))
            screen.blit(memory_text, (10, 10 + y_offset))
            y_offset += 25
            
        cpu_lines = display_cpu(c)
        y_offset = 0
        for line in cpu_lines:
            cpu_text = font.render(line, True, (255, 255, 255))
            screen.blit(cpu_text, (800, 10 + y_offset))
            y_offset += 25
            
        instruction_lines = display_instructions(c, c.pc, 50)
        y_offset = 400
        for line in instruction_lines:
            instruction_text = font.render(line, True, (255, 255, 255))
            screen.blit(instruction_text, (800, 10 + y_offset))
            y_offset += 25

        
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False

        pygame.display.update()

    pygame.quit()

if __name__ == "__main__":
    main()
