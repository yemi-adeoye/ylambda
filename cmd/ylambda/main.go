package main

import (
	"fmt"
	"ylambda/internal/cli"
	"ylambda/internal/executor"
	"ylambda/internal/reader"
	"ylambda/internal/repository"
	"ylambda/internal/service"
)

func main() {

	// get the file name
	args := cli.ParseArgs()

	def, err := reader.Read(reader.FileName(args.FileName))

	// create service
	// create in memory repository
	repo := repository.NewInMemoryRepository()
	exec := executor.NewDefaultExecutor()
	svc := service.NewCliService(repo, exec)

	res, err := svc.Execute(def)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)

}
