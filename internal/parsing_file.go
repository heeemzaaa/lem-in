package helpers

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// this function parses the file given to extract and validate data for an ant farm simulation
func (F *Farm) ReadFile(file *os.File) error {
	seen := make(map[string]bool, 0)
	var err error

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)
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
	F.Rooms = nil

	F.ExtarctBadRooms()

	return nil
}
