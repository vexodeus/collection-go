package grid

import "fmt"

type Grid struct {
	Fields [][]*Field
}
type Field struct {
	Row    int
	Column int
	Value  any
}

func NewGrid(m int, n int) *Grid {
	grid := &Grid{}
	for i := 0; i < m; i++ {
		grid.Fields = append(grid.Fields, make([]*Field, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid.Fields[i][j].Row = i
			grid.Fields[i][j].Row = j
		}
	}
	return grid
}
func (g *Grid) Size() (int, int) {
	rows := len(g.Fields)
	if rows == 0 {
		return rows, -1
	}
	return rows, len(g.Fields[0])
}
func (g *Grid) Valid(i int, j int) bool {
	m, n := g.Size()
	if i > m-1 || j > n-1 || i < 0 || j < 0 {
		return false
	}
	return true
}
func (g *Grid) Print() {
	for _, row := range g.Fields {
		fmt.Printf("\n")
		for _, field := range row {
			fmt.Printf("%v ", field.Value)
		}
	}
}
