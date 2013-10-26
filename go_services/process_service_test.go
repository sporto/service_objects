package go_services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/service_objects/go_services"
)

type FakeFindService struct {
}

func (c *FakeFindService) Run() []int {
	return []int{1, 2, 3}
}

var _ = Describe("ProcessService", func() {
	var (
		serv *go_services.ProcessService
	)

	BeforeEach(func() {
		serv = &go_services.ProcessService{}
	})

	It("uses the defaults", func() {
		res := serv.Run(2)
		Expect(res).To(Equal([]int{2, 4}))
	})

	It("can take a fake", func () {
		// change dependecies
		serv.FindService = new(FakeFindService)

		res := serv.Run(2)
		Expect(res).To(Equal([]int{2, 4, 6}))
	})
})