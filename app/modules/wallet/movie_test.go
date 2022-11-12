package wallet_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"goke/modules/wallet"
)

var _ = Describe("Movie", func() {
	var m1, m2 interface{
		GetTitle() string
		GetLenght() time.Duration
	}
	t1 := "m1"
	t2 := "m2"
	l1 := time.Duration(30)
	l2 := time.Duration(60)
	BeforeEach(func() {
		m1 = wallet.NewMovie(t1, l1)
		m2 = wallet.NewMovie(t2, l2)
	})

	Describe("Creating movies", func() {
		Context("with correct titles", func() {
		  It("title should be exact as when created", func() {
			Expect(m1.GetTitle()).To(Equal(t1))
			Expect(m2.GetTitle()).To(Equal(t2))
		  })
		})
	
		Context("with correct length", func() {
		  It("length should be exac as when created", func() {
			Expect(m1.GetLenght()).To(Equal(l1))
			Expect(m2.GetLenght()).To(Equal(l2))
		  })
		})
	  })
})
