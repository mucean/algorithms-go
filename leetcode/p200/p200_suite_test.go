package p200_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestP200(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P200 Suite")
}
