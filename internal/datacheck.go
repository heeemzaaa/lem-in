package lem

import (
	"os"
	"strconv"
	"strings"
)

// this function checks if the arguments settings are correct and returns nil if not
func ValidArgs(args []string) *os.File {
	if len(args) != 2 {
		return nil
	}
	data, err := os.Open(args[1])
	if err != nil {
		return nil
	}
	return data
}

// this function checks if the line given is a link or not
func Checklink(link string) bool {
	linkParts := strings.Split(link, "-")
	if len(linkParts) != 2 {
		return false
	}
	if checklinks[link] || checklinks[linkParts[1]+"-"+linkParts[0]] {
		return false
	}
	if !seen[linkParts[1]] || !seen[linkParts[0]] {
		return false
	}
	checklinks[link] = true
	checklinks[linkParts[1]+"-"+linkParts[0]] = true
	return true
}

// this function checks if the line given is a room or not
func Checkroom(room string) bool {
	roomParts := strings.Fields(room)
	if len(roomParts) != 3 {
		return false
	}
	_, err1 := strconv.Atoi(roomParts[1])
	if _, err2 := strconv.Atoi(roomParts[2]); err1 != nil || err2 != nil {
		return false
	}
	roomname := roomParts[0]
	if roomname[0] == 'L' {
		return false
	}
	if seen[roomname] {
		return false
	}
	seen[roomname] = true
	return true
}
