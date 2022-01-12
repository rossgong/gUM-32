package um32cpu

import (
	"fmt"
)

//See notes.txt for personal notes documenting the spec

//All return values are if the operation was successful

//Operator #0. Conditional Move
//IF C!=0:B->A
func Move(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	if *c > 0 {
		*a = *b
	}
	return nil
}

//Operator #1. Array Index
//B[C]->A
//Fails if Array B doesn't exist or if index is out of bounds
func Index(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	array, exists := collection.set[*b]

	if exists {
		if *c >= Platter(len(array)) {
			return fmt.Errorf("FAIL on operator \"Array Index\": Index %v is out of bounds", *c)
		} else {
			*a = array[*c]
			return nil
		}
	} else {
		return fmt.Errorf("FAIL on operator \"Array Index\": Array %v doesn't exist", *b)
	}
}

//Operator #2. Array Amendendment
//C -> A[B]
//Fails if Array A doesn't exist or B is out of bounds
func Amend(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	array, exists := collection.set[*a]

	if exists {
		if *b >= Platter(len(array)) {
			return fmt.Errorf("FAIL on operator \"Array Amendendment\": Index %v is out of bounds", *b)
		} else {
			array[*b] = *c
			return nil
		}
	} else {
		return fmt.Errorf("FAIL on operator \"Array Amendendment\": Array %v doesn't exist", *a)
	}
}

//Operator #3. Addition
//Spec doesn't specify whether to do anything with carry so ignore overflow
func Add(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	*a = *b + *c
	return nil
}

//Operator #4. Multiplication
//Spec doesn't specify whether to do anything with carry so ignore overflow
func Multiply(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	*a = *b * *c
	return nil
}

//Operator #5. Division
//B/C->A if B/C exists
//Fail if division by zero
func Division(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	if *c == 0 {
		return fmt.Errorf("FAIL on operator \"Division\": Divide by zero")
	} else {
		*a = *b / *c
		return nil
	}
}

//Operator #6. NAND
//B NAND C -> A
func NAnd(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	*a = ^(*b & *c)
	return nil
}

//Operator #7. Halt
//Stops machine
func Halt(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	return fmt.Errorf("HALT")
}

//Operator #8. Allocation
//Creates a new array with capacity c initalized with 0's
//Sets B with a unique identifier
func Allocate(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	*b = collection.newArray(*c)
	return nil
}

//Operator #9. Abandonment
//Deallocate array C
func Abandon(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	if *c == 0 {
		return fmt.Errorf("FAIL on operator \"Abandonment\": Cannot abandon program array (0)")
	}
	_, exists := collection.set[*c]

	if exists {
		delete(collection.set, *c)
		return nil
	} else {
		return fmt.Errorf("FAIL on operator \"Abandonment\": Array %v doesn't exist", *c)
	}
}

//Operator #10. Output
//display character in register C
func Output(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	if *c > 255 {
		//TODO: Fail
	}
	//TODO: Implement ouputs
	return nil
}

//Operator #11. Input
//Takes input from console into C. If it is end of input, C is filled with 1s
func Input(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	//TODO: Implement inputs
	return nil
}

//Operator #12. Load Program
//Load array B into program array (0) and set offset to C
func Load(a *Platter, b *Platter, c *Platter, collection *ArrayCollection) error {
	_, exists := collection.set[*b]

	if exists {
		collection.LoadProgramArray(*b)
		return nil
	} else {
		return fmt.Errorf("FAIL on operator \"Load Program\": Array %v doesn't exist", *b)
	}
}

//Operator #13. Orthography
//Set a to be value provided by opcode
func Ortho(a *Platter, v uint32) {
	*a = v
}
