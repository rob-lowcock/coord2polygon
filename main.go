package main

import "fmt"
import (
	"github.com/kr/pretty"
	"math"
)

type Coord struct {
	Lat, Long float64
}

type GridCell struct {
	X, Y float64
	Fill bool
}

type Grid struct {
	Cells        []GridCell
	SizeX, SizeY float64
	XRes, YRes         int
}

func main() {
	input := []Coord{
		{4, 1 },
		{3, 2},
		{4, 2},
		{2, 3 },
		{4, 3 },
		{5, 4 },
		{2, 4 },
		{1, 5 },
		{6, 5 },
		{2, 6 },
		{5, 6 },
		{1, 9 },
		{5, 9 },
		{1, 11 },
		{3, 11 },
		{6, 11 },
		{2, 12 },
		{5, 12 },
		{3, 13 },
		{6, 13 },
		{4, 13},
		{4, 15 },
		{4, 17 },
	}

	xRes := 6
	yRes := 17

	leftTop, rightBottom := calculateGridLimits(input)
	grid := generateGrid(leftTop, rightBottom, xRes, yRes)
	grid = populateGrid(grid, input)

	pretty.Println(leftTop)
	pretty.Println(rightBottom)

	pretty.Println("Cells: ----")
	pretty.Println(grid)

	// 1. Generate grid from coordinates
	// 2. Populate grid from coordinates
	// 3. Calculate extremities

	fmt.Println("Hello world")
}
func populateGrid(grid Grid, coords []Coord) Grid {
	out := grid.Cells


	for _, value := range coords {
		for k, cell := range grid.Cells {
			if isCell(value, cell, grid.SizeX, grid.SizeY, (k + 1 % grid.XRes) == 0, (k + 1) >= len(grid.Cells) - grid.YRes) {
				out[k].Fill = true
			}
		}
	}

	grid.Cells = out
	return grid
}

func isCell(coord Coord, cell GridCell, sizeX float64, sizeY float64, endRow, endCol bool) bool {
	if endCol && endRow {
		return coord.Long >= cell.Y && coord.Lat >= cell.X
	}

	if endRow {
		return coord.Lat >= cell.X &&
			coord.Long >= cell.Y &&
			coord.Long < (cell.Y + sizeY)
	}

	if endCol {
		return coord.Lat >= cell.X &&
			coord.Lat < (cell.X + sizeX) &&
			coord.Long >= cell.Y
	}

	return coord.Lat >= cell.X &&
		coord.Lat < (cell.X + sizeX) &&
		coord.Long >= cell.Y &&
		coord.Long < (cell.Y + sizeY)
}

func generateGrid(leftTop Coord, rightBottom Coord, xRes int, yRes int) Grid {
	var out []GridCell

	sizeX := math.Ceil(rightBottom.Lat - leftTop.Lat) / float64(xRes)
	sizeY := math.Ceil(rightBottom.Long - leftTop.Long) / float64(yRes)

	for j := leftTop.Long; j < float64(yRes); j += sizeY {
		for i := leftTop.Lat; i < float64(xRes); i += sizeX {
			out = append(out, GridCell{
				X: i,
				Y: j,
			})
		}
	}

	return Grid{
		Cells: out,
		SizeX: sizeX,
		SizeY: sizeY,
		XRes: xRes,
		YRes: yRes,
	}
}

func calculateGridLimits(coords []Coord) (Coord, Coord) {
	leftTop := coords[0]
	rightBottom := coords[0]
	for _, value := range coords {
		if value.Lat < leftTop.Lat {
			leftTop.Lat = value.Lat
		}

		if value.Long < leftTop.Long {
			leftTop.Long = value.Long
		}

		if value.Lat > rightBottom.Lat {
			rightBottom.Lat = value.Lat
		}

		if value.Long > rightBottom.Long {
			rightBottom.Long = value.Long
		}
	}

	return leftTop, rightBottom
}
