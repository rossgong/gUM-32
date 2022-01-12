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

func (collection *ArrayCollection) LoadProgramArray(arrayIndex Platter) {
	newProgram := make([]Platter, len(collection.set[arrayIndex]))
	copy(newProgram, collection.set[arrayIndex])
	collection.set[0] = newProgram
}

func (collection *ArrayCollection) setArray(index Platter, array []Platter) {
	collection.set[index] = array
}

func CreateArrayCollection(program []Platter) (collection ArrayCollection) {
	collection = *new(ArrayCollection)
	return
}
