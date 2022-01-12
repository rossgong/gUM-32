package um32cpu

const (
	registerAmount    = 8
	programArrayIndex = 0
)

type (
	Platter          = uint32
	OperatorFunction = func(*Platter, *Platter, *Platter, *ArrayCollection) error
)

type CPU struct {
	//Array offset for program (PC)
	finger    uint
	running   bool
	registers [registerAmount]Platter

	arrays ArrayCollection
}

type Operation struct {
	operator OperatorFunction
	a        *Platter
	b        *Platter
	c        *Platter

	isOrtho bool
	v       Platter
}

func InitializeCPU(program []Platter) (cpu CPU) {
	cpu = CPU{}
	cpu.arrays.setArray(programArrayIndex, program)
	cpu.running = true

	return cpu
}

func (cpu *CPU) Cycle() error {
	operatorCode := cpu.arrays.getOperator(cpu.finger)
	cpu.finger++

	operation, err := decode(operatorCode)
	if err == nil {
		if operation.isOrtho {
			return Ortho(operation.a, operation.v)
		} else {
			return operation.operator(operation.a, operation.b, operation.c, &cpu.arrays)
		}
	}
	return err
}

func decode(opCode Platter) (Operation, error) {
	return Operation{}, nil
}
