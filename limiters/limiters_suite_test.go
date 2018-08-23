package limiters_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLimiters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Limiters Suite")
}
