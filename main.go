// main.go
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"lemin/helpers"
)

// add somthig to test in my life

func main() {
	farm := &helpers.Farm{
		LinksCheck: make(map[string]any),
		Links:      make(map[string][]string),
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	err = farm.ReadFile(file)
	file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	farm.FindPaths() // Call BFS to find all paths
	// lets filter some bad paths
	farm.Filter()
	// defer fmt.Println("All paths from start to end:", farm.ValidPaths)
	// // fmt.Println("number of paths:", len(farm.ValidPaths))
	// defer fmt.Println(farm.BadRooms)
	// defer fmt.Println(farm.Badpaths, "badpaths")
	if len(farm.ValidPaths) == 0 {
		log.Fatal("ERROR: invalid data format, no path found")
	}
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// lets prait the output
	output, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(output)
	fmt.Println()
	fmt.Println()
	farm.Sendants()
}