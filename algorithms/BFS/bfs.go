package bfs

type Grid struct {
	Slice [][]any
}

var directions = map[string][]int{
	"up":    {0, -1},
	"right": {1, 0},
	"down":  {0, 1},
	"left":  {-1, 0},
}

func (g *Grid) BFS(i int, j int) {
	if !g.Valid(i, j) {
		return
	}
	for _, dir := range directions {
		g.BFS(i+dir[0], j+dir[1])
	}
}
func (g *Grid) Valid(i int, j int) bool {
	maxX := len(g.Slice) - 1
	if i > maxX {
		return false
	}
	maxY := len(g.Slice[i])
	return j > maxY

}
