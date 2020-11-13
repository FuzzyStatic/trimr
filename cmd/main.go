package main

import (
	"github.com/FuzzyStatic/trimr/internal/cmd"
)

const (
	progName = "trimr"
)

// Compile time variables
var Version string

func main() {
	t, err := cmd.NewTrimr(progName, Version)
	if err != nil {
		panic(err)
	}
	_ = t.Execute()
}
