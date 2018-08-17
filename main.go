package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"

	"github.com/kr/pretty"
)

// Coord is a set of Latitude (X) and Longitude (Y) coordinates.
type Coord struct {
	Long float64 `csv:"longitude"`
	Lat  float64 `csv:"latitude"`
}

// Longitude: Perpendicular to equator (X axis)
// Latitude: Parallel to equator (Y axis)

// GridCell contains the details of a cell on the grid
type GridCell struct {
	X, Y float64
	Fill bool
}

// Grid contains a slice of cells, and meta information about the grid
type Grid struct {
	Cells        []GridCell
	SizeX, SizeY float64
	XRes, YRes   int
}

// GetCoords returns the outer coordinates for a populated grid
func (grid Grid) GetCoords() []Coord {
	var out []Coord

	// Traverse columns left to right
	for i := 0; i < grid.XRes; i++ {
		// Scan down the columns
		for j := i; j < len(grid.Cells)/2; j += grid.XRes {
			if grid.Cells[j].Fill == true {
				out = append(out, Coord{grid.Cells[j].X, grid.Cells[j].Y})
				break
			}
		}
	}

	// Traverse the rows top to bottom
	for i := 0; i < grid.YRes; i++ {
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
	for i := len(grid.Cells) - 1; i >= len(grid.Cells)-grid.XRes; i-- {
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
		for j := i; j < i+(grid.XRes/2); j++ {
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
	// input := []*Coord{
	// 	{0.4, 0.1},
	// 	{0.3, 0.2},
	// 	{0.4, 0.2},
	// 	{0.2, 0.3},
	// 	{0.4, 0.3},
	// 	{0.5, 0.4},
	// 	{0.2, 0.4},
	// 	{0.1, 0.5},
	// 	{0.6, 0.5},
	// 	{0.2, 0.6},
	// 	{0.5, 0.6},
	// 	{0.1, 0.9},
	// 	{0.5, 0.9},
	// 	{0.1, 1.1},
	// 	{0.3, 1.1},
	// 	{0.6, 1.1},
	// 	{0.2, 1.2},
	// 	{0.5, 1.2},
	// 	{0.3, 1.3},
	// 	{0.6, 1.3},
	// 	{0.4, 1.3},
	// 	{0.4, 1.5},
	// 	{0.4, 1.7},
	// }

	inputFile, err := os.OpenFile("compiled.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	var input []*Coord

	if err := gocsv.UnmarshalFile(inputFile, &input); err != nil { // Load clients from file
		panic(err)
	}

	xRes := 100
	yRes := 100

	leftTop, rightBottom := calculateGridLimits(input)
	grid := generateGrid(leftTop, rightBottom, xRes, yRes)

	grid = populateGrid(grid, input)

	debugGrid(grid, xRes)

	coords := grid.GetCoords()

	pretty.Println(coords)

	// 1. Generate grid from coordinates
	// 2. Populate grid from coordinates
	// 3. Calculate extremities
}
func debugGrid(grid Grid, xRes int) {
	counter := 0
	printer := ""
	for _, value := range grid.Cells {
		if value.Fill {
			printer = printer + "O"
		} else {
			printer = printer + "."
		}

		counter++

		if counter%xRes == 0 {
			fmt.Println(printer)
			printer = ""
		}
	}
}
func populateGrid(grid Grid, coords []*Coord) Grid {
	out := grid.Cells

	for _, value := range coords {
		for k, cell := range grid.Cells {
			if isCell(*value, cell, grid.SizeX, grid.SizeY, (k+1)%grid.XRes == 0, (k+1) >= len(grid.Cells)-grid.YRes) {
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
			coord.Lat < (cell.Y+sizeY)
	}

	// If this cell is at the end of a column
	if endCol {
		return coord.Long >= cell.X &&
			coord.Long < (cell.X+sizeX) &&
			coord.Lat >= cell.Y
	}

	return coord.Long >= cell.X &&
		coord.Long < (cell.X+sizeX) &&
		coord.Lat >= cell.Y &&
		coord.Lat < (cell.Y+sizeY)
}

func generateGrid(leftTop Coord, rightBottom Coord, xRes int, yRes int) Grid {
	var out []GridCell

	sizeX := (rightBottom.Long - leftTop.Long) / float64(xRes)
	sizeY := (rightBottom.Lat - leftTop.Lat) / float64(yRes)

	pretty.Println("Sizes:")
	pretty.Println(sizeX)
	pretty.Println(sizeY)

	for j := 0; j < yRes; j++ {
		for i := 0; i < xRes; i++ {
			out = append(out, GridCell{
				X: leftTop.Long + (float64(i) * sizeX),
				Y: leftTop.Lat + (float64(j) * sizeY),
			})
		}
	}

	return Grid{
		Cells: out,
		SizeX: sizeX,
		SizeY: sizeY,
		XRes:  xRes,
		YRes:  yRes,
	}
}

func calculateGridLimits(coords []*Coord) (Coord, Coord) {
	leftTop := Coord{
		Long: coords[0].Long,
		Lat:  coords[0].Lat,
	}
	rightBottom := Coord{
		Long: coords[0].Long,
		Lat:  coords[0].Lat,
	}
	for _, value := range coords {
		if value.Long < leftTop.Long {
			pretty.Println("New leftTop Long", value.Long)
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
