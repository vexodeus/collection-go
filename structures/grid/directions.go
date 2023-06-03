package grid

type Direction struct {
	DX int
	DY int
}

var Dir4 = map[string]Direction{
	"E": {1, 0},
	"S": {0, 1},
	"W": {-1, 0},
	"N": {0, -1},
}
var Dir8 = map[string]Direction{
	"E":  {1, 0},
	"SE": {1, 1},
	"S":  {0, 1},
	"SW": {-1, 1},
	"W":  {-1, 0},
	"NW": {-1, -1},
	"N":  {0, -1},
	"NE": {1, -1},
}
