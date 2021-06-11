package mmu

/*
 * The GameBoy has a 16-bit address bus
 * that is used to address ROM, RAM and I/O.
 *
 *
 * Memory map:
 * 0x0000 - 0x3FFF: ROM bank 00
 * 0x4000 - 0x7FFF: ROM bank 01
 * 0x8000 - 0x9FFF: VRAM
 * 0xA000 - 0xBFFF: RAM
 * 0xC000 - 0xCFFF: WRAM (Work RAM)
 * 0xD000 - 0xDFFF: WRAM
 * 0xE000 - 0xFDFF: ECHO RAM (Mirror of 0xC000 - 0xDDFF)
 * 0xFE00 - 0xFE9F: OAM (Sprite attribute table)
 * 0xFEA0 - 0xFEFF: Not usable
 * 0xFF00 - 0xFF7F: I/O registers
 * 0xFF80 - 0xFFFE: HRAM (High RAM)
 * 0xFFFF - 0xFFFF: IE (Interrupt Enable register)
 *
 */

// Read 8-bit byte from a given address
func RB(address uint16) (value uint8) {
	return 0xff
}

// Read 16-bit word from a given address
func RW(address uint16) (value uint16) {
	return 0xffff
}

// Write 8-bit byte to a given address
func WB(address uint16, value uint8) {}

// Write 16-bit word to a given address
func WW(address uint16, value uint16)
