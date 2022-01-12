package um32cpu

const (
	registerAmount = 8
)

type (
	Platter = uint32
)

type CPU struct {
	finger    uint
	registers [registerAmount]Platter

	arrays ArrayCollection
}
