package repository

import "ylambda/internal/model"

type Repository interface {
	Save(function *model.Function) (*model.Function, error)
	Get(functionName string) (*model.Function, bool)
	Exists(function *model.Function) bool
	ToFunction(def *model.FunctionJsonDef) *model.Function
}
