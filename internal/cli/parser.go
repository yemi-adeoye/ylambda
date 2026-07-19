package cli

import (
	"flag"
	"fmt"
)

func ParseArgs() {
	var functionName string
	var functionVersion string
	var functionDescription string

	flag.StringVar(&functionName, "f", "", "Function name")
	flag.StringVar(&functionVersion, "v", "", "Function version")
	flag.StringVar(&functionDescription, "d", "", "Function description")

	flag.Parse()

	fmt.Println("Function Name:", functionName)
	fmt.Println("Function Version:", functionVersion)
	fmt.Println("Function Description:", functionDescription)

	otherArgs := flag.Args()

	for _, arg := range otherArgs {
		fmt.Println(arg)
	}
}
