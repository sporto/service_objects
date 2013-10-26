package go_services

type FindServiceInt interface {
	Run() []int
}

type FindService struct {
}

func (c *FindService) Run() []int {
	return []int{1,2}
}
