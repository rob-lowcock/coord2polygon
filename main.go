package main

import (
	"github.com/kr/pretty"
	"math"
)

type Coord struct {
	Long, Lat float64
}

// Longitude: Perpendicular to equator (X axis)
// Latitude: Parallel to equator (Y axis)

type GridCell struct {
	X, Y float64
	Fill bool
}

type Grid struct {
	Cells        []GridCell
	SizeX, SizeY float64
	XRes, YRes         int
}

func (grid Grid) GetCoords() []Coord {
	var out []Coord

	// Traverse columns left to right
	for i := 0; i < grid.XRes; i += 1{
		// Scan down the columns
		for j := i; j < len(grid.Cells)/2; j += grid.XRes  {
			if grid.Cells[j].Fill == true {
				out = append(out, Coord{grid.Cells[j].X, grid.Cells[j].Y})
				break
			}
		}
	}

	// Traverse the rows top to bottom
	for i := 0; i < grid.YRes; i += 1 {
		// Scan backwards through the rows
		limit := (i * grid.XRes) + (grid.XRes / 2)

		for j := ((i + 1) * grid.XRes) - 1; j >= limit; j-- {
			if grid.Cells[j].Fill == true {
				coord := Coord{grid.Cells[j].X, grid.Cells[j].Y}
				if !inSlice(out, coord) {
					out = append(out, coord)
				}
				break
			}
		}
	}

	// Traverse the columns right to left
	for i := len(grid.Cells) - 1; i >= len(grid.Cells) - grid.XRes; i-- {
		// Scan the columns bottom to top
		for j := i; j > i/2; j -= grid.XRes {
			if grid.Cells[j].Fill == true {
				coord := Coord{grid.Cells[j].X, grid.Cells[j].Y}
				if !inSlice(out, coord) {
					out = append(out, coord)
				}
				break
			}
		}
	}
	
	// Finally, traverse the rows bottom to top
	for i := len(grid.Cells) - grid.XRes; i >= 0; i -= grid.XRes {
		// Scan through the rows
		for j := i; j < i + (grid.XRes/2); j++ {
			if grid.Cells[j].Fill == true {
				coord := Coord{grid.Cells[j].X, grid.Cells[j].Y}
				if !inSlice(out, coord) {
					out = append(out, coord)
				}
				break
			}
		}
	}

	return out
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

	coords := grid.GetCoords()

	pretty.Println(coords)

	// 1. Generate grid from coordinates
	// 2. Populate grid from coordinates
	// 3. Calculate extremities
}
func populateGrid(grid Grid, coords []Coord) Grid {
	out := grid.Cells


	for _, value := range coords {
		for k, cell := range grid.Cells {
			if isCell(value, cell, grid.SizeX, grid.SizeY, (k + 1) % grid.XRes == 0, (k + 1) >= len(grid.Cells) - grid.YRes) {
				out[k].Fill = true
			}
		}
	}

	grid.Cells = out
	return grid
}

func isCell(coord Coord, cell GridCell, sizeX float64, sizeY float64, endRow, endCol bool) bool {
	// If this cell is at the end of the column and the end of the row(i.e. it's the last cell)
	if endCol && endRow {
		return coord.Lat >= cell.Y && coord.Long >= cell.X
	}

	// If this cell is at the end of a row
	if endRow {
		return coord.Long >= cell.X &&
			coord.Lat >= cell.Y &&
			coord.Lat < (cell.Y + sizeY)
	}

	// If this cell is at the end of a column
	if endCol {
		return coord.Long >= cell.X &&
			coord.Long < (cell.X + sizeX) &&
			coord.Lat >= cell.Y
	}

	return coord.Long >= cell.X &&
		coord.Long < (cell.X + sizeX) &&
		coord.Lat >= cell.Y &&
		coord.Lat < (cell.Y + sizeY)
}

func generateGrid(leftTop Coord, rightBottom Coord, xRes int, yRes int) Grid {
	var out []GridCell

	sizeX := math.Ceil(rightBottom.Long- leftTop.Long) / float64(xRes)
	sizeY := math.Ceil(rightBottom.Lat- leftTop.Lat) / float64(yRes)

	for j := leftTop.Lat; j < float64(yRes); j += sizeY {
		for i := leftTop.Long; i < float64(xRes); i += sizeX {
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
		if value.Long < leftTop.Long {
			leftTop.Long = value.Long
		}

		if value.Lat < leftTop.Lat {
			leftTop.Lat = value.Lat
		}

		if value.Long > rightBottom.Long {
			rightBottom.Long = value.Long
		}

		if value.Lat > rightBottom.Lat {
			rightBottom.Lat = value.Lat
		}
	}

	return leftTop, rightBottom
}

func inSlice(haystack []Coord, needle Coord) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}

	return false
}