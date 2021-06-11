package cpu

// Internal CPU state

type clock struct {
	m uint16
	t uint16
}

type registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	h uint8
	l uint8
	f uint8

	pc uint16
	sp uint16

	m uint8
	t uint8
}

type state struct {
	clock     clock
	registers registers
}

var State = state{}

// Opcodes

func NOP() {}

func LD() {}

func INC() {}
func DEC() {}

func ADC() {}
func ADD() {}
func SBC() {}
func SUB() {}

func AND() {}
func XOR() {}
func OR()  {}
func CP()  {}

func POP()  {}
func PUSH() {}

func JP()   {}
func JR()   {}
func CALL() {}
func RET()  {}
func RETI() {}
func HALT() {}
func RST()  {}
func STOP() {}

func RLCA() {}
func RRCA() {}
func RLA()  {}
func RRA()  {}

func DAA() {}
func CPL() {}
func SCF() {}
func CCF() {}

func DI() {}
func EI() {}

func LDH() {}

func RLC()  {}
func RRC()  {}
func RL()   {}
func RR()   {}
func SLA()  {}
func SRA()  {}
func SWAP() {}
func SRL()  {}
func BIT()  {}
func RES()  {}
func SET()  {}
