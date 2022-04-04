package integer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInteger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integer Suite")
}
