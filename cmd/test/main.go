package main

import (
	"os"

	"github.com/pkg/profile"
	"gongaware.org/gUM32/pkg/um32cpu"
	"gongaware.org/gUM32/pkg/um32parser"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	file, err := os.Open(os.Args[1])
	if err == nil {
		program, err := um32parser.Parse(file)
		if err != nil {
			panic(err)
		}
		cpu := um32cpu.InitializeCPU(program)
		for i := 0; true; i++ {
			err = cpu.Cycle()
			if err != nil {
				panic(err)
			}
		}
	} else {
		panic(err)
	}
}
