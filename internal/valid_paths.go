package helpers

// This function finds all valid paths from the start room to the end room,
// categorizing them into valid and bad paths based on the presence of bad rooms.
func (F *Farm) FindPaths() {
	type QueueItem struct {
		room string
		path []string
	}
	startNeighbors := F.Links[F.StartRoom]
	for _, room := range startNeighbors {
		queue := []QueueItem{{room: room, path: []string{room}}}
		visited := make(map[string]bool)
		visited[F.StartRoom] = true

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if current.room == F.EndRoom {
				path := []string{F.StartRoom}
				if !Contains(current.path, F.BadRooms) {
					if len(F.ValidPaths) == F.Ants {
						break
					}
					path = append(path, current.path...)
					F.ValidPaths = append(F.ValidPaths, path)
				} else {
					path = append(path, current.path...)
					F.Badpaths = append(F.Badpaths, path)
					continue
				}
			}

			visited[current.room] = true
			for _, neighbor := range F.Links[current.room] {
				if !visited[neighbor] {
					newPath := append([]string{}, current.path...)
					newPath = append(newPath, neighbor)
					queue = append(queue, QueueItem{room: neighbor, path: newPath})
				}
			}
		}
	}
}


// this function filters bad paths, 
// adding them to valid paths if they don't overlap with existing valid paths,
// while removing redundant bad paths.
func (F *Farm) Filter() {
	for i, badpath := range F.Badpaths {
		count := 0
		for _, goodpath := range F.ValidPaths {
			if !Contains(badpath[1:len(badpath)-1], goodpath[1:len(goodpath)-1]) {
				count++
			} else {
				if i == 0 {
					continue
				}
				F.Badpaths[i] = nil
			}
		}

		if count == len(F.ValidPaths) {
			F.ValidPaths = append(F.ValidPaths, badpath)
		}

	}
}
