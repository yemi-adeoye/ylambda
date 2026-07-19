package controller

type Controller interface {
	Delegate(service Service, method string)
}
