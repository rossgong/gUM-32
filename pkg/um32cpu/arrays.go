package um32cpu

type ArrayCollection struct {
	set map[Platter][]Platter

	//Have a seperate Program array to prevent a map lookup every cycle
	programArray []Platter

	nextSlot Platter
}

func (collection *ArrayCollection) newArray(capacity Platter) (index Platter) {
	for {
		_, exists := collection.set[collection.nextSlot]
		if exists {
			collection.nextSlot++
		} else {
			break
		}
		//Check to see if all slots are filled? (that's ALOT of arrays)
	}

	collection.set[collection.nextSlot] = make([]Platter, capacity)
	collection.nextSlot++
	return collection.nextSlot - 1
}

func (collection *ArrayCollection) LoadProgramArray(arrayIndex Platter) {
	futureProgram := collection.set[arrayIndex]
	numChanged := copy(collection.programArray, futureProgram)

	//If the array was fully copied make sure to get rid of the rest
	//If the numChanged is less than the source array, append the rest
	//This is to prevent maving to make new backing arrays every time
	//This is a very large speedup
	if numChanged == len(futureProgram) {
		collection.programArray = collection.programArray[:numChanged]
	} else {
		collection.programArray = append(collection.programArray, futureProgram[numChanged:]...)
	}
	collection.set[0] = collection.programArray
}

func (collection *ArrayCollection) setArray(index Platter, array []Platter) {
	collection.set[index] = array
}

func CreateArrayCollection(program []Platter) (collection *ArrayCollection) {
	collection = new(ArrayCollection)
	collection.set = make(map[uint32][]uint32)
	collection.nextSlot = 1
	//Put dummy array to prevent this slot from being used as 0 is the programArray
	//collection set 0 will always be a shallow copy of this
	collection.programArray = program
	collection.set[0] = collection.programArray
	return
}
