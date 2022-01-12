package um32cpu

import "fmt"

const (
	registerAmount    = 8
	programArrayIndex = 0

	//These are the shifts needed to get the relevant operator number
	registerShiftA = 6
	registerShiftB = 3
	//registerShiftC = 0 unused as it is 0 just mask

	orthoShiftA = 25

	//Bit masks
	registerMask   = 0x0000_0007 //Shift and then mask
	operatorMask   = 0xF000_0000
	orthoValueMask = 0x01FF_FFFF
)

type (
	Platter          = uint32
	OperatorFunction = func(*Platter, *Platter, *Platter, *ArrayCollection) error
)

type CPU struct {
	//Array offset for program (PC)
	finger    Platter
	running   bool
	registers [registerAmount]Platter

	arrays ArrayCollection
}

type Operation struct {
	operator OperatorFunction
	a        byte
	b        byte
	c        byte

	isOrtho bool
	v       Platter
}

func InitializeCPU(program []Platter) (cpu CPU) {
	cpu = CPU{}
	cpu.arrays = CreateArrayCollection(program)
	cpu.arrays.setArray(programArrayIndex, program)
	cpu.running = true

	return cpu
}

func (cpu *CPU) Cycle() error {
	operatorCode := cpu.arrays.getOperator(cpu.finger)
	cpu.finger++

	op, err := decode(operatorCode, cpu)
	if err == nil {
		if op.isOrtho {
			return Ortho(&cpu.registers[op.a], op.v)
		} else {
			return op.operator(&cpu.registers[op.a], &cpu.registers[op.b], &cpu.registers[op.c], &cpu.arrays)
		}
	}
	return err
}

func decode(opCode Platter, cpu *CPU) (Operation, error) {
	//Mask out all non-opcode bits
	operatorNumber := opCode & operatorMask
	op := Operation{}

	//No operator above ortho
	if operatorNumber > 0xD000_0000 {
		return op, fmt.Errorf("FAIL in decode: Platter %b is an invalid operation", opCode)
	} else if operatorNumber == 0xD000_0000 {
		op.isOrtho = true
		op.v = opCode & orthoValueMask
		op.a = byte((opCode >> orthoShiftA) & registerMask)
	} else {
		op.a = byte((opCode >> registerShiftA) & registerMask)
		op.b = byte((opCode >> registerShiftB) & registerMask)
		op.c = byte(opCode & registerMask)
		switch operatorNumber {
		case 0x0000_0000:
			op.operator = Move
		case 0x1000_0000:
			op.operator = Index
		case 0x2000_0000:
			op.operator = Amend
		case 0x3000_0000:
			op.operator = Add
		case 0x4000_0000:
			op.operator = Multiply
		case 0x5000_0000:
			op.operator = Division
		case 0x6000_0000:
			op.operator = NAnd
		case 0x7000_0000:
			op.operator = Halt
		case 0x8000_0000:
			op.operator = Allocate
		case 0x9000_0000:
			op.operator = Abandon
		case 0xA000_0000:
			op.operator = Output
		case 0xB000_0000:
			op.operator = Input
		case 0xC000_0000:
			cpu.finger = cpu.registers[op.c]
			op.operator = Load
		}
	}

	return op, nil
}
