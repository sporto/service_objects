package services

import (
// "github.com/sporto/service_objects/services"
)

type KillerService struct {
	FinderService FinderServiceInt
	input         string
}

func (c *KillerService) Run(input string) string {
	c.input = input

	return c.find()
}

func (c *KillerService) find() string {
	if c.FinderService == nil {
		c.FinderService = new(FinderService)
	}
	found := c.FinderService.Run()
	return c.input + " found and killed " + found
}
