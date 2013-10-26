package go_services_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/sporto/service_objects/go_services"
)

var _ = Describe("ProcessService", func() {
  var (
    serv *go_services.ProcessService
  )

  BeforeEach(func() {
    serv = &go_services.ProcessService{}
  })

  It("uses the defaults", func() {
    Expect(serv.Run(2)).To(Equal([]int{2, 4}))
  })
})