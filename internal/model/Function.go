package model

import (
	"errors"
	"fmt"
)

type Function struct {
	Name        string
	Version     string
	Description string
	Runtime     Runtime
}

const (
	GoRuntime     Runtime = "go"
	PythonRuntime Runtime = "python"
	NodeRuntime   Runtime = "node"
)

type Runtime string

func NewFunction(name string, version string, runtime Runtime) (*Function, error) {
	if name == "" {
		return nil, errors.New("invalid function name")
	}

	if version == "" {
		return nil, errors.New("invalid function version")
	}

	if runtime == "" {
		return nil, errors.New("invalid function runtime")
	}

	return &Function{
		Name:    name,
		Runtime: runtime,
		Version: version,
	}, nil
}

func (function *Function) String() string {
	return fmt.Sprintf("Function: \nname: %s, \nversion: %s, \nruntime: %s",
		function.Name,
		function.Version,
		function.Runtime,
	)
}
