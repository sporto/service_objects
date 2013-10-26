package go_services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/service_objects/go_services"
)

var _ = Describe("FindService", func() {
	var (
		serv *go_services.FindService
	)

	BeforeEach(func() {
		serv = &go_services.FindService{}
	})

	It("returns a slice", func() {
		Expect(serv.Run()).To(Equal([]int{1,2}))
	})
})