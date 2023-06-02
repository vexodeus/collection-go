package grid

type Grid struct {
	Slice [][]any
}

func NewGrid(m int, n int) *Grid {
	grid := &Grid{}
	for i := 0; i < m; i++ {
		grid.Slice = append(grid.Slice, make([]any, n))
	}
	return grid
}
func (g *Grid) Len() (int, int) {
	rows := len(g.Slice)
	if rows == 0 {
		return rows, -1
	}
	return rows, len(g.Slice[0])
}
func (g *Grid) Valid(i int, j int) bool {
	m, n := g.Len()
	if i > m-1 || j > n-1 || i < 0 || j < 0 {
		return false
	}
	return true
}
