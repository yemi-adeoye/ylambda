package executor

import (
	"fmt"
	"ylambda/internal/model"
)

type DefaultExecutor struct {
	state string
}

func NewDefaultExecutor() *DefaultExecutor {
	return &DefaultExecutor{
		state: "PENDING",
	}
}

func (e *DefaultExecutor) Execute(function *model.Function) error {
	fmt.Println("Executing function: " + function.Name)
	return nil
}
