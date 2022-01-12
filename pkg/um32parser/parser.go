package um32parser

import (
	"encoding/binary"
	"fmt"
	"os"

	"gongaware.org/gUM32/pkg/um32cpu"
)

func Parse(file *os.File) ([]um32cpu.Platter, error) {
	fileStats, err := file.Stat()
	if err == nil {
		fileSize := fileStats.Size()
		if fileSize%4 == 0 {
			program := make([]um32cpu.Platter, fileSize/4)
			err = binary.Read(file, binary.BigEndian, &program)

			return program, err
		} else {
			return nil, fmt.Errorf("this file appears to have a bad file size")
		}
	}
	return nil, err
}
