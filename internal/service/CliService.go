package service

import (
	"ylambda/internal/executor"
	"ylambda/internal/model"
	"ylambda/internal/repository"
)

type CliService struct {
	repository repository.Repository
	executor   executor.Executor
}

func NewCliService(repository repository.Repository, executor executor.Executor) *CliService {
	return &CliService{
		repository: repository,
		executor:   executor,
	}
}

func (cliService *CliService) Execute(functionJsonDef *model.FunctionJsonDef) (bool, error) {
	function := cliService.repository.ToFunction(functionJsonDef)

	_, err := cliService.repository.Save(function)

	if err != nil {
		return false, err
	}

	err = cliService.executor.Execute(function)

	if err != nil {
		return false, err
	}
	return true, nil
}
