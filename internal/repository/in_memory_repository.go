package repository

import (
	"errors"
	"ylambda/internal/model"
)

type InMemoryRepository struct {
	data map[string]*model.Function
}

func (repository *InMemoryRepository) ToFunction(def *model.FunctionJsonDef) *model.Function {
	return &model.Function{
		Name:        def.Name,
		Version:     def.Version,
		Description: def.Description,
		Runtime:     def.Runtime,
	}
}

func NewInMemoryRepository() *InMemoryRepository {

	return &InMemoryRepository{
		data: make(map[string]*model.Function),
	}
}

func (repository *InMemoryRepository) Save(function *model.Function) (*model.Function, error) {
	_, exists := repository.data[function.Name]

	if exists {
		return function, errors.New("function with function name")
	}
	repository.data[function.Name] = function
	return function, nil
}

func (repository *InMemoryRepository) Get(functionName string) (*model.Function, bool) {
	function, exists := repository.data[functionName]
	return function, exists
}

func (repository *InMemoryRepository) Exists(function *model.Function) bool {
	_, exists := repository.Get(function.Name)
	return exists
}

func (repository *InMemoryRepository) String() {
	for _, function := range repository.data {
		function.String()
	}
}
