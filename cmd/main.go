package main

import (
	"github.com/FuzzyStatic/trimr/internal/cmd"
)

const (
	progName = "trimr"
)

// Compile time variables
var (
	Version   string
	BuildTime string
	BuildHost string
)

func main() {
	t, err := cmd.NewTrimr(
		progName,
		Version,
		BuildTime,
		BuildHost,
	)
	if err != nil {
		panic(err)
	}
	_ = t.Execute()
}
