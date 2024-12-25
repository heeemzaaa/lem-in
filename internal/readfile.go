package lem

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// this function reads the file and parse it to have all the data needed
func ReadFile(file *os.File) string {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	nline := 0
	linked := false
	for scanner.Scan() {
		nline++
		line := strings.TrimSpace(scanner.Text())
		if i == 0 {
			num, err := strconv.Atoi(line)
			if err != nil {
				return "ERROR: invalid number of ants (LINE:" + strconv.Itoa(nline) + ")"
			}
			if num == 0 {
				return "ERROR: number of ants cant be 0 (LINE:" + strconv.Itoa(nline) + ")"
			}
			Ants = num
			i = 1
			continue
		}
		if strings.HasPrefix(line, "#") || line == "" {
			if line == "##start" {
				i = 2
			} else if line == "##end" {
				i = 3
			}
			continue
		}
		if i == 2 {
			if !Checkroom(line) || Start != "" {
				return "ERROR: invalid start rooom (LINE:" + strconv.Itoa(nline) + ")"
			}
			Start = room(line)
			i = 1
		} else if i == 3 {
			if !Checkroom(line) || End != "" {
				return "ERROR: invalid end rooom (LINE:" + strconv.Itoa(nline) + ")"
			}
			End = room(line)
			i = 1
		} else if !Checkroom(line) {
			if !Checklink(line) {
				return "ERROR: invalid data format (LINE:" + strconv.Itoa(nline) + ")"
			}
			linked = true
			link := strings.Split(line, "-")
			room1 := link[0]
			room2 := link[1]
			Ways[room1] = append(Ways[room1], room2)
			Ways[room2] = append(Ways[room2], room1)
		} else if linked {
			return "ERROR: invalid data format , link inside rooms (LINE:" + strconv.Itoa(nline) + ")"
		}
	}
	checklinks = nil
	if Start == "" {
		return "ERROR: invalid data format , no start room found"
	}
	if End == "" {
		return "ERROR: invalid data format , no end room found"
	}
	return ""
}

// this function splits the line of the room and take the first index as the name of the room
func room(s string) string {
	return strings.Fields(s)[0]
}
