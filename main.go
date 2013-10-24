package main

import (
	"fmt"
	"github.com/sporto/service_objects/services"
)

type FakeFinderService struct {
}

func (c *FakeFinderService) Run() string {
	return "spiders"
}

func main() {
	// create the service
	serv := new(services.KillerService)

	// run the service
	res := serv.Run("James")
	fmt.Println(res)

	// fmt.Println(new(FakeFinderService).Run())

	// change dependecies
	serv.FinderService = new(FakeFinderService)

	// run with changed dependecy
	res = serv.Run("Sam")
	fmt.Println(res)
}
