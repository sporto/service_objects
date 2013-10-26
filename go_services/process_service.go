package go_services

import (
// "github.com/sporto/service_objects/services"
	// "github.com/sporto/service_objects/go_services"
)

type ProcessService struct {
	FindService   FindServiceInt
	input         int
  found         []int
}

func (c *ProcessService) Run(input int) []int {
	c.input = input
	c.find()
	res := []int{}
	for _,element := range c.found {
		// element is the element from someSlice for where we are
		res = append(res, element * input)
	}
	return res
}

func (c *ProcessService) find() {
	if c.FindService == nil {
		c.FindService = new(FindService)
	}
	c.found = c.FindService.Run()
}
