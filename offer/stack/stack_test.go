package stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mucean/algorithms-go/offer/stack"
)

var _ = Describe("Stack", func() {
	Context("21-MinStack", func() {
		It("test normal", func() {
			minStack := Constructor()
			minStack.Push(5)
			Expect(minStack.Top()).To(Equal(5))
			Expect(minStack.Min()).To(Equal(5))

			minStack.Push(4)
			Expect(minStack.Top()).To(Equal(4))
			Expect(minStack.Min()).To(Equal(4))

			minStack.Push(6)
			Expect(minStack.Top()).To(Equal(6))
			Expect(minStack.Min()).To(Equal(4))

			minStack.Push(2)
			Expect(minStack.Top()).To(Equal(2))
			Expect(minStack.Min()).To(Equal(2))

			minStack.Push(3)
			Expect(minStack.Top()).To(Equal(3))
			Expect(minStack.Min()).To(Equal(2))

			minStack.Pop()
			Expect(minStack.Top()).To(Equal(2))
			Expect(minStack.Min()).To(Equal(2))

			minStack.Pop()
			Expect(minStack.Top()).To(Equal(6))
			Expect(minStack.Min()).To(Equal(4))

			minStack.Pop()
			Expect(minStack.Top()).To(Equal(4))
			Expect(minStack.Min()).To(Equal(4))

			minStack.Pop()
			Expect(minStack.Top()).To(Equal(5))
			Expect(minStack.Min()).To(Equal(5))

			minStack.Pop()
			Expect(minStack.Top()).To(Equal(0))
			Expect(minStack.Min()).To(Equal(0))

			minStack.Pop()
			Expect(minStack.Top()).To(Equal(0))
			Expect(minStack.Min()).To(Equal(0))
		})

		It("test same element", func() {
			minStack := Constructor()

			minStack.Push(5)
			minStack.Push(2)
			minStack.Push(3)
			minStack.Push(2)
			minStack.Push(4)

			minStack.Pop()
			Expect(minStack.Min()).To(Equal(2))

			minStack.Pop()
			Expect(minStack.Min()).To(Equal(2))

			minStack.Pop()
			Expect(minStack.Min()).To(Equal(2))

			minStack.Pop()
			Expect(minStack.Min()).To(Equal(5))

			minStack.Pop()
			Expect(minStack.Min()).To(Equal(0))
		})
	})

	Context("22-IsPopOrder", func() {
		DescribeTable(
			"IsPopOrder",
			func(pushed, popped []int, res bool) {
				Expect(IsPopOrder(pushed, popped)).To(Equal(res))
			},
			Entry("nil, nil", []int(nil), []int(nil), true),
			Entry("empty, empty", []int{}, []int{}, true),
			Entry("length not equal", []int{1}, []int{1, 2}, false),
			Entry("is", []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true),
			Entry("is", []int{1, 0}, []int{1, 0}, true),
			Entry("not is", []int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false),
		)
	})
})
