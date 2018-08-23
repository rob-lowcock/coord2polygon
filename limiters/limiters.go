package limiters

import (
	"github.com/rob-lowcock/coord2polygon/grid"
)

// TopLimit returns whether a cell should be used when scanning from the top
func TopLimit(colIndex, rowIndex int, g grid.Grid) bool {
	if rowIndex <= g.XRes/2 {
		if colIndex <= g.YRes/2 {
			return rowIndex <= colIndex
		}

		return (g.XRes - colIndex) > rowIndex
	}

	return false
}

// RightLimit returns whether a cell should be used when scanning from the right
func RightLimit(colIndex, rowIndex int, g grid.Grid) bool {
	if rowIndex < g.XRes/2 {
		if colIndex < g.YRes/2 {
			return false
		}

		return (g.XRes - (colIndex + 1)) <= rowIndex

	}

	return rowIndex <= colIndex
}

// BottomLimit returns whether a cell should be used when scanning from the bottom
func BottomLimit(colIndex, rowIndex int, g grid.Grid) bool {
	if rowIndex >= g.XRes/2 {
		if colIndex >= g.YRes/2 {
			return rowIndex >= colIndex
		}

		return (g.XRes - (colIndex + 1)) <= rowIndex
	}

	return false
}

// Left returns whether a cell should be used when scanning from the left
func LeftLimit(colIndex, rowIndex int, g grid.Grid) bool {
	if rowIndex > g.XRes/2 {
		if colIndex > g.YRes/2 {
			return false
		}

		return (g.XRes - (colIndex + 1)) >= rowIndex

	}

	return rowIndex >= colIndex
}
