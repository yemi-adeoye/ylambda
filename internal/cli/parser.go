package cli

import (
	"flag"
)

type Args struct {
	FileName string
}

func ParseArgs() *Args {
	var fileName string

	flag.StringVar(&fileName, "f", "", "Usage: -f <fileName>")

	flag.Parse()

	return &Args{
		FileName: fileName,
	}
}
