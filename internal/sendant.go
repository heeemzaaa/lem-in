package lem

import (
	"fmt"
	"strconv"
)

// this function initialize a map that have the index of the slice in the key and the length in the value
func InitMap(s [][]string) map[int]int {
	mapp := make(map[int]int)
	for v := range s {
		mapp[v] = len(s[v])
	}
	return mapp
}

// this function takes the smallest slice in the solution set
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
func Sendants() {
	antGroups := make([][]string, len(solutions))
	antId := 1
	mapp := InitMap(solutions)

	for antId <= Ants {
		minPath := Getmin(mapp)
		antGroups[minPath] = append(antGroups[minPath], "L"+strconv.Itoa(antId))
		antId++
		mapp[minPath]++
	}
	control_trafic(antGroups)
}

// this function controls the trafic of the ants , ensuring no room conflicts and tracking progress until all ants reaches the end room
func control_trafic(antGroups [][]string) {
	trafic := make(map[string]int)
	Emptyroom := make(map[string]bool)
	finished := 0
	for finished != Ants {
		for i := 0; i < len(solutions); i++ {
			Emptyroom[End] = false
			for currentStep := 0; currentStep < len(antGroups[i]); currentStep++ {
				ant := antGroups[i][currentStep]
				nextroom := solutions[i][trafic[ant]+1]
				if !Emptyroom[nextroom] {
					fmt.Printf("%v-%v ", ant, nextroom)
					Emptyroom[nextroom] = true
					Emptyroom[solutions[i][trafic[ant]]] = false
					if nextroom == End {
						finished++
						delete(trafic, ant)
						antGroups[i] = append(antGroups[i][:currentStep], antGroups[i][currentStep+1:]...)
						currentStep--
						Emptyroom[End] = true
						continue
					}
					trafic[ant]++
				}
			}
		}
		fmt.Println()
	}
}
