package main

import "ylambda/internal/service"

func main() {

}

func doMe(service service.Service) {
	service.Execute("hello world")
}
