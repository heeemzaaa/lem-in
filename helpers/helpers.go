// helpers/helpers.go
package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Farm struct {
	 LinksCheck        map[string]any
	Links                    map[string][]string
	StartNeighbots, BadRooms []string
	Badpaths, ValidPaths     [][]string
	StartRoom, EndRoom       string
	Ants                     int
}

// extract all the bad rooms
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
	// lets extract the room that has more links
	maxsize := len(BadRooms[len(BadRooms)-1]) - 1
	for i, room := range BadRooms {
		if len(room) == maxsize+1 && room[0] != F.EndRoom && room[0] != F.StartRoom {
			F.BadRooms = append(F.BadRooms, room[0])
		}
		// free up the room
		BadRooms[i] = nil
	}
}
// find the paths
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
			// Dequeue the front item
			current := queue[0]
			queue = queue[1:]

			// If we reach the end room, save the path
			if current.room == F.EndRoom {
				path := []string{F.StartRoom}
				if !Contains(current.path, F.BadRooms) {
					// visited_for_Bad	rooms[current.path[0]] = true
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

			// Mark the current room as visited
			visited[current.room] = true
			// Explore neighbors
			for _, neighbor := range F.Links[current.room] {
				if !visited[neighbor] {
					newPath := append([]string{}, current.path...) // Copy the current path
					newPath = append(newPath, neighbor)            // Add the neighbor to the path
					queue = append(queue, QueueItem{room: neighbor, path: newPath})
				}
			}
		}
	}
}

// check if the pats are vcalid
func Contains(s []string, e []string) bool {
	for _, a := range s {
		for _, b := range e {
			if a == b {
				return true
			}
		}
	}
	return false
}

// reading the file and getting the data
func (F *Farm) ReadFile(file *os.File) error {
	// open the file
	seen := make(map[string]bool, 0)
	var err error

	// read the file by using the buffio pkg
	// that can give us convenient way to read input from a file
	// line by line using the  function scan()
	// befor looping lets inisialise our maps

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)
		// lets check if the first line is the valid number off ants
		if i == 0 {
			F.Ants, err = strconv.Atoi(line)
			if err != nil {
				return err
			}
			if F.Ants <= 0 {
				return errors.New("ERROR: invalid data format, minvalid ants number")
			}
			i++
			continue
		}
		if line == "" || (strings.HasPrefix(line, "#") && line != "##start" && line != "##end") {
			continue
		}

		if i == 2 {
			check := strings.Fields(line)
			F.StartRoom = check[0]

			i = 1

		}

		if i == 3 {
			check := strings.Split(line, " ")
			F.EndRoom = check[0]

			i = 1

		}

		if line == "##start" {
			if F.StartRoom != "" {
				return errors.New("ERROR: invalid data format, multiple start room")
			}
			i = 2
			continue
		}
		if line == "##end" {
			if F.EndRoom != "" {
				return errors.New("ERROR: invalid data format, multiple end room")
			}
			i = 3
			continue
		}

		check := strings.Fields(line)
		if len(check) == 3 {
			if strings.HasPrefix(check[0], "L") {
				return errors.New("ERROR: invalid data format, Invalid room ")
			}
			if i == -7777777777777777777 {
				return errors.New("ERROR: invalid data format, link inside rooms ")
			}
			_, exist := seen[check[0]]
			if !exist {
				seen[check[0]] = true
			} else {
				return errors.New("ERROR: invalid data format,found Duplicated rooms")
			}

		} else if len(check) == 1 {
			link := strings.Split(line, "-")

			i = -7777777777777777777
			if len(link) != 2 {
				return errors.New("ERROR: invalid data format, no valid link found")
			}
			_, exist := seen[link[0]]
			if !exist {
				return errors.New("ERROR: invalid data format, no valid link found")
			}
			_, exist1 := seen[link[1]]
			if !exist1 {
				return errors.New("ERROR: invalid data format, no valid link found")
			}
			if link[0] == link[1] {
				return errors.New("ERROR: invalid data format, no valid link found")
			}
			_, exist2 := F.LinksCheck[link[0]+"-"+link[1]]
			_, exist3 := F.LinksCheck[link[1]+"-"+link[0]]
			if exist2 || exist3 {
				return errors.New("ERROR: invalid data format, duplicated link found")
			}
			F.LinksCheck[link[0]+"-"+link[1]] = nil
			F.LinksCheck[link[1]+"-"+link[0]] = nil
			F.Links[link[0]] = append(F.Links[link[0]], link[1])
			F.Links[link[1]] = append(F.Links[link[1]], link[0])

		} else {
			return errors.New("ERROR: invalid data format")
		}

	}
	if F.EndRoom == "" || F.StartRoom == "" {
		return errors.New("ERROR: invalid data format, no start or end room")
	}
	F.LinksCheck = nil
	

	F.ExtarctBadRooms()

	return nil
}

// filter the bad rooms and return the valid paths with the bad room
func (F *Farm) Filter() {
	for i, badpath := range F.Badpaths {
		count := 0
		for _, goodpath := range F.ValidPaths {
			if !Contains(badpath[1:len(badpath)-1], goodpath[1:len(goodpath)-1]) {
				count++
			} else {
				// delete this path
				// dont remove all bad paths save the first one is shortest one
				// for 1 ants
				if i == 0 {
					if F.Ants == 1 {
						F.ValidPaths = append(F.ValidPaths, F.Badpaths[0])
					}
				}
				F.Badpaths[i] = nil
			}
		}

		if count == len(F.ValidPaths) {
			F.ValidPaths = append(F.ValidPaths, badpath)
		}

	}
}

func InitMap(s [][]string) map[int]int {
	mapp := make(map[int]int)
	for v := range s {
		mapp[v] = len(s[v])
	}
	return mapp
}

func Getmin(v map[int]int) int {
	i := v[0]
	x := 0
	for t, vv := range v {
		if vv < i {
			i = vv
			x = t
		}
	}
	return x
}

// sending the ants to the farm
func (F *Farm) Sendants() {
	antGroups := make([][]string, len(F.ValidPaths))
	antId := 1
	mapp := InitMap(F.ValidPaths)

	for antId <= F.Ants {
		minPath := Getmin(mapp)
		antGroups[minPath] = append(antGroups[minPath], "L"+strconv.Itoa(antId))
		antId++
		mapp[minPath]++
	}
	F.control_trafic(antGroups)
}

func (F *Farm) control_trafic(antGroups [][]string) {
	trafic := make(map[string]int)
	Emptyroom := make(map[string]bool)
	finished := 0
	for finished != F.Ants {
		for i := 0; i < len(F.ValidPaths); i++ {
			Emptyroom[F.EndRoom] = false
			for currentStep := 0; currentStep < len(antGroups[i]); currentStep++ {
				ant := antGroups[i][currentStep]
				nextroom := F.ValidPaths[i][trafic[ant]+1]
				if !Emptyroom[nextroom] {
					fmt.Printf("%v-%v ", ant, nextroom)
					Emptyroom[nextroom] = true
					Emptyroom[F.ValidPaths[i][trafic[ant]]] = false
					if nextroom == F.EndRoom {
						finished++
						delete(trafic, ant)
						antGroups[i] = append(antGroups[i][:currentStep], antGroups[i][currentStep+1:]...)
						currentStep--
						Emptyroom[F.EndRoom] = true
						continue
					}
					trafic[ant]++
				}
			}
		}
		fmt.Println()
	}
}
