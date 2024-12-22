package helpers

import "sort"

// this function identifies and records the bad rooms in the farm
// by analyzing rooms with the most links, excluding the start and end rooms.
func (F *Farm) ExtarctBadRooms() {
	BadRooms := make([][]string, 0)
	for room, links := range F.Links {
		badroom := []string{room}
		badroom = append(badroom, links...)
		BadRooms = append(BadRooms, badroom)
	}

	sort.Slice(BadRooms, func(i, j int) bool {
		return len(BadRooms[i]) < len(BadRooms[j])
	})
	maxsize := len(BadRooms[len(BadRooms)-1]) - 1
	for i, room := range BadRooms {
		if len(room) == maxsize+1 && room[0] != F.EndRoom && room[0] != F.StartRoom {
			F.BadRooms = append(F.BadRooms, room[0])
		}
		BadRooms[i] = nil
	}
}


