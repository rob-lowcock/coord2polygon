package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/rob-lowcock/coord2polygon/grid"

	"github.com/kr/pretty"
)

func topScanLimit(x, y, xRes, totalRes int) bool {
	if x < xRes/2 {
		return x < (xRes + 1*y)
	}

	return x < totalRes-(xRes/2)
}

func main() {
	inputFile, err := os.OpenFile("compiled.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	var input []*grid.Coord

	if err := gocsv.UnmarshalFile(inputFile, &input); err != nil { // Load clients from file
		panic(err)
	}

	xRes := 100
	yRes := 100

	grid := grid.GenerateGrid(input, xRes, yRes)

	grid = populateGrid(grid, input)

	debugGrid(grid, xRes)

	coords := grid.GetCoords()

	pretty.Println(coords)

	// 1. Generate grid from coordinates
	// 2. Populate grid from coordinates
	// 3. Calculate extremities
}
func debugGrid(g grid.Grid, xRes int) {
	counter := 0
	printer := ""
	for _, value := range g.Cells {
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
func populateGrid(g grid.Grid, coords []*grid.Coord) grid.Grid {
	out := g.Cells

	for _, value := range coords {
		for k, cell := range g.Cells {
			if isCell(*value, cell, g.SizeX, g.SizeY, (k+1)%g.XRes == 0, (k+1) >= len(g.Cells)-g.YRes) {
				out[k].Fill = true
			}
		}
	}

	g.Cells = out
	return g
}

func isCell(coord grid.Coord, cell grid.Cell, sizeX float64, sizeY float64, endRow, endCol bool) bool {
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
