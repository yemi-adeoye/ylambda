package model

import "fmt"

type FunctionJsonDef struct {
	Name        string
	Version     string
	Runtime     Runtime
	Description string
	Text        string
}

func (functionJsonDef FunctionJsonDef) String() string {
	return fmt.Sprintf("Function Definition: \nname: %s, \nversion: %s, \nruntime: %s, \ndescription: %s, \ntext: %s",
		functionJsonDef.Name,
		functionJsonDef.Version,
		functionJsonDef.Runtime,
		functionJsonDef.Description,
		functionJsonDef.Text)
}
