package differ

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Differ", func() {
	origin := "외않되"
	fixed := "왜 안돼"
	result := "[31m외않되[0m[32m왜 안돼[0m"

	Describe("Diff", func() {
		It("returns both of origin string and fixed string", func() {
			diff := Diff(origin, fixed)

			Expect(diff).To(Equal(result))
		})
	})
})
