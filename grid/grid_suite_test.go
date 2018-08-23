package grid_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGrid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grid Suite")
}
