package lem

// this is the main function of our logic , it chooses the optimal set of paths for the ants
func Search() [][]string {
	visited[Start] = true
	sol1 := SearchHelper(Ants, false)
	visited = make(map[string]bool)
	visited[Start] = true
	sol2 := SearchHelper(Ants, true)
	Choose(sol1, sol2)
	Sort1()
	return solutions
}

// select the best path between the two paths given usng the average
func Choose(sol1, sol2 [][]string) {
	if len(sol1) < len(sol2) {
		solutions = sol2
		return
	}
	if len(sol1) > len(sol2) {
		solutions = sol1
		return
	}
	if len(sol1) == 0 {
		solutions = sol2
		return
	}
	avg1 := Average(sol1)
	avg2 := Average(sol2)
	if avg1 > avg2 {
		solutions = sol2
		return
	}
	solutions = sol1
}

// this function calculate the average of each slice in the slice of slices given
func Average(slice [][]string) int {
	avg := 0
	for _, v := range slice {
		avg += len(v)
	}
	return avg / len(slice)
}

// this function implements the BFS logic to find paths from the start to the end
func SearchHelper(ant int, rev bool) [][]string {
	sol := [][]string{}
	if !rev {
		for i := 0; i < len(Ways[Start]); i++ {
			Bfs(Ways[Start][i], &sol, &ant)
			Close(sol)
		}
	} else {
		for i := len(Ways[Start]) - 1; i >= 0; i-- {
			Bfs(Ways[Start][i], &sol, &ant)
			Close(sol)
		}
	}
	return sol
}

// this function checks if two paths given are the same
func Compare(s1, s2 []string) bool {
	if len(s1) != len(s2) || len(s1) == 0 || len(s2) == 0 {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}
	return true
}

// this function sorts the solutions set based on the path length and removes duplicate paths
func Sort1() {
	for i := 0; i < len(solutions)-1; i++ {
		for j := i + 1; j < len(solutions); j++ {
			if len(solutions[j]) < len(solutions[i]) {
				solutions[i], solutions[j] = solutions[j], solutions[i]
			}
			if Compare(solutions[j][1:len(solutions[i])-1], solutions[i][1:len(solutions[i])-1]) {
				solutions = append(solutions[:i], solutions[i+1:]...)
			}
		}
	}
}

// this function perfomes the BFS algorithm to find paths from a given starting place to the end
func Bfs(s string, solution1 *[][]string, ant *int) bool {
	solution := *solution1
	parent := make(map[string]string)
	parent[s] = Start
	if s == End {

		solutions = append(solutions, Findway(parent))
		return false
	}

	visited[Start] = true
	queu := []string{s}
	visited[s] = true

	for i := 0; i < len(queu); i++ {
		visiting := queu[i]
		for _, neighbour := range Ways[visiting] {
			if !visited[neighbour] {
				visited[neighbour] = true
				parent[neighbour] = visiting
				queu = append(queu, neighbour)
			}
			if neighbour == End {
				newPath := Findway(parent)
				if len(solution) > 0 {
					if len(solution) > Ants {
						if len(solution[len(solution)-1]) > len(newPath) {
							solution[len(solution)-1] = newPath
							*solution1 = solution
						}
						return false
					} else if len(newPath) < *ant {
						*ant -= len(solution[len(solution)-1])
						solution = append(solution, newPath)
						*solution1 = solution
						return false
					}
				}
				solution = append(solution, newPath)
				*solution1 = solution
				return false

			}
		}
	}
	return true
}

// func ccc(s string) bool {
// 	for _, v := range badrooms {
// 		if v == s {
// 			return true
// 		}
// 	}
// 	return false
// }

// this function start from the end traces back to the start
func Findway(parent map[string]string) []string {
	curent := End
	visited = make(map[string]bool)
	way := []string{curent}
	for curent != Start {
		way = append(way, parent[curent])
		curent = parent[curent]
	}
	return Flip(way)
}

// this function marks all the places in the current solution set as visited
func Close(sol [][]string) {
	visited = make(map[string]bool)
	for _, s := range sol {
		for _, v := range s[1 : len(s)-1] {
			visited[v] = true
		}
	}
}


// this function reverse the order of elements in a slice
func Flip(s []string) []string {
	r := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
