package um32cpu

type ArrayCollection struct {
	set map[Platter][]Platter

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

func (collection *ArrayCollection) getOperator(offset Platter) Platter {
	return collection.set[programArrayIndex][offset]
}

func (collection *ArrayCollection) LoadProgramArray(arrayIndex Platter) {
	futureProgram := collection.set[arrayIndex]
	numChanged := copy(collection.set[programArrayIndex], futureProgram)

	//If the array was fully copied make sure to get rid of the rest
	//If the numChanged is less than the source array, append the rest
	//This is to prevent maving to make new backing arrays every time
	//This is a very large speedup
	if numChanged == len(futureProgram) {
		collection.set[programArrayIndex] = collection.set[programArrayIndex][:numChanged]
	} else {
		collection.set[programArrayIndex] = append(collection.set[programArrayIndex], futureProgram[numChanged:]...)
	}
}

func (collection *ArrayCollection) setArray(index Platter, array []Platter) {
	collection.set[index] = array
}

func CreateArrayCollection(program []Platter) (collection ArrayCollection) {
	collection = *new(ArrayCollection)
	collection.set = make(map[uint32][]uint32)
	return
}
