package grid

import (
	"github.com/kr/pretty"
)

// Longitude: Perpendicular to equator (X axis)
// Latitude: Parallel to equator (Y axis)

// Cell contains the details of a cell on the grid
type Cell struct {
	X, Y float64
	Fill bool
}

// Grid contains a slice of cells, and meta information about the grid
type Grid struct {
	Cells        []Cell
	SizeX, SizeY float64
	XRes, YRes   int
}

// Coord is a set of Latitude (X) and Longitude (Y) coordinates.
type Coord struct {
	Long float64 `csv:"longitude"`
	Lat  float64 `csv:"latitude"`
}

// GetCoords returns the outer coordinates for a populated grid
func (g Grid) GetCoords() []Coord {
	var out []Coord

	// Traverse columns left to right
	for i := 0; i < g.XRes; i++ {
		// Scan down the columns
		for j := i; j < len(g.Cells)/2; j += g.XRes {
			if g.Cells[j].Fill == true {
				out = append(out, Coord{g.Cells[j].X, g.Cells[j].Y})
				break
			}
		}
	}

	// Traverse the rows top to bottom
	for i := 0; i < g.YRes; i++ {
		// Scan backwards through the rows
		limit := (i * g.XRes) + (g.XRes / 2)

		for j := ((i + 1) * g.XRes) - 1; j >= limit; j-- {
			if g.Cells[j].Fill == true {
				coord := Coord{g.Cells[j].X, g.Cells[j].Y}
				if !inSlice(out, coord) {
					out = append(out, coord)
				}
				break
			}
		}
	}

	// Traverse the columns right to left
	for i := len(g.Cells) - 1; i >= len(g.Cells)-g.XRes; i-- {
		// Scan the columns bottom to top
		for j := i; j > i/2; j -= g.XRes {
			if g.Cells[j].Fill == true {
				coord := Coord{g.Cells[j].X, g.Cells[j].Y}
				if !inSlice(out, coord) {
					out = append(out, coord)
				}
				break
			}
		}
	}

	// Finally, traverse the rows bottom to top
	for i := len(g.Cells) - g.XRes; i >= 0; i -= g.XRes {
		// Scan through the rows
		for j := i; j < i+(g.XRes/2); j++ {
			if g.Cells[j].Fill == true {
				coord := Coord{g.Cells[j].X, g.Cells[j].Y}
				if !inSlice(out, coord) {
					out = append(out, coord)
				}
				break
			}
		}
	}

	return out
}

// GenerateGrid generates the grid to populate
func GenerateGrid(input []*Coord, xRes int, yRes int) Grid {
	leftTop, rightBottom := calculateGridLimits(input)

	var out []Cell

	sizeX := (rightBottom.Long - leftTop.Long) / float64(xRes)
	sizeY := (rightBottom.Lat - leftTop.Lat) / float64(yRes)

	pretty.Println("Sizes:")
	pretty.Println(sizeX)
	pretty.Println(sizeY)

	for j := 0; j < yRes; j++ {
		for i := 0; i < xRes; i++ {
			out = append(out, Cell{
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
