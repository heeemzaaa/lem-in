package helpers

type Farm struct {
	Rooms, LinksCheck        map[string]any
	Links                    map[string][]string
	StartNeighbots, BadRooms []string
	Badpaths, ValidPaths     [][]string
	StartRoom, EndRoom       string
	Ants                     int
}
