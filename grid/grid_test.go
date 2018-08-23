package grid_test

import (
	"github.com/kr/pretty"
	"github.com/rob-lowcock/coord2polygon/grid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {

	var dummyGrid grid.Grid

	BeforeEach(func() {
		dummyGrid = grid.Grid{
			Cells: []grid.Cell{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 1,
					Y: 0,
				},
				{
					X: 2,
					Y: 0,
				},
				{
					X: 3,
					Y: 0,
				},
				{
					X: 4,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
				{
					X: 1,
					Y: 1,
				},
				{
					X: 2,
					Y: 1,
				},
				{
					X: 3,
					Y: 1,
				},
				{
					X: 4,
					Y: 1,
				},
				{
					X: 0,
					Y: 2,
				},
				{
					X: 1,
					Y: 2,
				},
				{
					X: 2,
					Y: 2,
				},
				{
					X: 3,
					Y: 2,
				},
				{
					X: 4,
					Y: 2,
				},
				{
					X: 0,
					Y: 3,
				},
				{
					X: 1,
					Y: 3,
				},
				{
					X: 2,
					Y: 3,
				},
				{
					X: 3,
					Y: 3,
				},
				{
					X: 4,
					Y: 3,
				},
				{
					X: 0,
					Y: 4,
				},
				{
					X: 1,
					Y: 4,
				},
				{
					X: 2,
					Y: 4,
				},
				{
					X: 3,
					Y: 4,
				},
				{
					X: 4,
					Y: 4,
				},
			},
			XRes: 5,
			YRes: 5,
		}
	})

	It("calculates top limits correctly", func() {
		expected := []bool{
			true, true, true, true, true,
			false, true, true, true, false,
			false, false, true, false, false,
			false, false, false, false, false,
			false, false, false, false, false,
		}

		for k, v := range dummyGrid.Cells {
			Expect(dummyGrid.TopLimit(int(v.X), int(v.Y))).To(Equal(expected[k]))
		}
	})

	It("calculates right limits correctly", func() {
		expected := []bool{
			false, false, false, false, true,
			false, false, false, true, true,
			false, false, true, true, true,
			false, false, false, true, true,
			false, false, false, false, true,
		}

		for k, v := range dummyGrid.Cells {
			Expect(dummyGrid.RightLimit(int(v.X), int(v.Y))).To(Equal(expected[k]))
		}
	})

	It("calculates bottom limits correctly", func() {
		expected := []bool{
			false, false, false, false, false,
			false, false, false, false, false,
			false, false, true, false, false,
			false, true, true, true, false,
			true, true, true, true, true,
		}

		for k, v := range dummyGrid.Cells {
			Expect(dummyGrid.BottomLimit(int(v.X), int(v.Y))).To(Equal(expected[k]))
		}
	})

	It("calculates left limits correctly", func() {
		expected := []bool{
			true, false, false, false, false,
			true, true, false, false, false,
			true, true, true, false, false,
			true, true, false, false, false,
			true, false, false, false, false,
		}

		for k, v := range dummyGrid.Cells {
			Expect(dummyGrid.LeftLimit(int(v.X), int(v.Y))).To(Equal(expected[k]))
		}
	})
})

func debug(g grid.Grid, f func(x, y int, g grid.Grid) bool) {
	debugOut := "\n"

	for k, v := range g.Cells {

		if f(int(v.X), int(v.Y), g) {
			debugOut = debugOut + "O"
		} else {
			debugOut = debugOut + "."
		}

		if (k+1)%5 == 0 {
			pretty.Println(debugOut)
			debugOut = ""
		}
	}
}
