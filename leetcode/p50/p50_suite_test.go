package p50_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestP50(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P50 Suite")
}
