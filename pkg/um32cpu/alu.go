package um32cpu

//See notes.txt for personal notes documenting the spec

//Operator #0. Conditional Move
//IF C!=0:B->A
func Move(a *Platter, b *Platter, c *Platter) {
	if *c > 0 {
		*a = *b
	}
}

//Operator #1. Array Index
//B[C]->A
//Fails if Array B doesn't exist
func Index(a *Platter, b *Platter, c *Platter) {
	//TODO: Implement Arrays
}

//Operator #2. Array Amendendment
//A[B] = C
//Fails if Array B doesn't exist
func Amend(a *Platter, b *Platter, c *Platter) {
	//TODO: Implement Arrays
}

//Operator #3. Addition
//Spec doesn't specify whether to do anything with carry so ignore overflow
func Add(a *Platter, b *Platter, c *Platter) {
	*a = *b + *c
}

//Operator #4. Multiplication
//Spec doesn't specify whether to do anything with carry so ignore overflow
func Multiply(a *Platter, b *Platter, c *Platter) {
	*a = *b * *c
}

//Operator #5. Division
//B/C->A if B/C exists
//Fail if division by zero
func Division(a *Platter, b *Platter, c *Platter) {
	if *c == 0 {
		//TODO: Fail
	} else {
		*a = *b / *c
	}
}

//Operator #6. NAND
//B NAND C -> A
func NAnd(a *Platter, b *Platter, c *Platter) {
	*a = ^(*b & *c)
}

//Operator #7. Halt
//Stops machine
func Halt(a *Platter, b *Platter, c *Platter) {
	//TODO: Implement stops
}

//Operator #8. Allocation
//Creates a new array with capacity c initalized with 0's
//Seets B with a unique identifier
func Allocate(a *Platter, b *Platter, c *Platter) {
	//TODO: Implement arrays
}

//Operator #9. Abdonment
//Deallocate array C
func Abandon(a *Platter, b *Platter, c *Platter) {
	if *c == 0 {
		//TODO: Fail
	}
	//TODO: Implement arrays
}

//Operator #10. Output
//display character in register C
func Output(a *Platter, b *Platter, c *Platter) {
	if *c > 255 {
		//TODO: Fail
	}
	//TODO: Implement ouputs
}

//Operator #11. Input
//Takes input from console into C. If it is end of input, C is filled with 1s
func Input(a *Platter, b *Platter, c *Platter) {
	//TODO: Implement inputs
}

//Operator #12. Load Program
//Load array B into program array (0) and set offset to C
func Load(a *Platter, b *Platter, c *Platter) {
	//TODO: Implement arrays
}

//Operator #13. Orthography
//Set a to be value provided by opcode
func Ortho(a *Platter, v uint32) {
	*a = v
}
