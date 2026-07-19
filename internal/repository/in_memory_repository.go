package repository

import "ylambda/internal/model"

type InMemoryRepository struct {
	repository Repository
	data       map[string]model.Function
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (repository *InMemoryRepository) Save(functionName string, function model.Function) (model.Function, error) {
	_, exists := repository.data[functionName]

	if exists {
		return function, error.New("Function with function name")
	}
	repository.data[functionName] = function
	return function, nil
}
