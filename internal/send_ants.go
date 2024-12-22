package helpers

import (
	"fmt"
	"strconv"
)

// this function initializes a map that associates each index of a 2D slice 
// with the length of its corresponding inner slice.
func InitMap(s [][]string) map[int]int {
	mapp := make(map[int]int)
	for v := range s {
		mapp[v] = len(s[v])
	}
	return mapp
}


// this function returns the key of the smallest value in a map of integers.
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

// this function assigns ants to valid paths in a balanced manner
// based on path lengths and manages their traffic accordingly.
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


// this function simulates and manages the movement of ants along valid paths,
// ensuring no room conflicts and tracking progress until all ants reach the end room.
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
