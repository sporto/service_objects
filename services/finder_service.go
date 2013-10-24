package services

type FinderServiceInt interface {
	Run() string
}

type FinderService struct {
}

func (c *FinderService) Run() string {
	return "ants"
}
