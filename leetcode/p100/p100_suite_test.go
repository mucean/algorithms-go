package p100_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestP100(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P100 Suite")
}
