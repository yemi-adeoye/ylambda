package repository

import "ylambda/internal/model"

type Repository interface {
	Save(functionName string, function model.Function) (model.Function, error)
	get(functionName string) model.Function
}
