package grid

import . "github.com/vexodeus/collection-go/structures/generic/queue"

type ReturnFunc func(field *Field, config *BFSConfig) bool

type ProcessFunc func(field *Field, config *BFSConfig)

type BFSConfig struct {
	Directions map[string]Direction
	Return     ReturnFunc
	Process    ProcessFunc
	Data       *BFSData
}
type BFSData struct {
	Queue      *Queue[*Field]
	Visited    map[*Field]bool
	ResultGrid *Grid
}

func (g *Grid) NewBFSConfig() *BFSConfig {
	return &BFSConfig{
		Directions: make(map[string]Direction),
		Return:     func(field *Field, config *BFSConfig) bool { return true },
		Process:    func(field *Field, config *BFSConfig) {},
		Data: &BFSData{
			Queue:      NewQueue[*Field](),
			Visited:    make(map[*Field]bool),
			ResultGrid: NewGrid(g.Size()),
		},
	}
}
func (bfs *BFSConfig) SetDirections(directions map[string]Direction) {
	bfs.Directions = directions
}
func (bfs *BFSConfig) SetReturn(f ReturnFunc) {
	bfs.Return = f
}
func (bfs *BFSConfig) SetProcess(f ProcessFunc) {
	bfs.Process = f
}
func (bfs *BFSConfig) Visited(f *Field) bool {
	_, ok := bfs.Data.Visited[f]
	return ok
}
func (g *Grid) BFS(i int, j int, config *BFSConfig) {
	var current, adjacent *Field
	if !g.Valid(i, j) {
		return
	}
	current = g.Fields[i][j]
	if config.Return(current, config) || config.Visited(current) {
		return
	}
	config.Process(current, config)
	config.Data.Visited[g.Fields[i][j]] = true
	config.Data.Queue.Enqueue(current)
	for config.Data.Queue.Len() > 0 {
		current = config.Data.Queue.Dequeue()
		config.Process(current, config)
		config.Data.Visited[g.Fields[i][j]] = true
		for _, dir := range config.Directions {
			if g.Valid(current.Row+dir.DX, current.Column+dir.DY) {
				adjacent = g.Fields[current.Row+dir.DX][current.Column+dir.DY]
				if config.Return(adjacent, config) || config.Visited(adjacent) {
					continue
				}
				config.Data.Queue.Enqueue(adjacent)
			}
		}
	}
	config.Data.Queue.Dequeue()
}
