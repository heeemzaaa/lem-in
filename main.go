package main

import (
	"fmt"
	"io"
	"log"
	"os"

	lem "lem/internal"
)

func main() {
	err := lem.ReadFile(lem.ValidArgs(os.Args))
	if err != "" {
		log.Fatal(err)
	}
	validways := lem.Search()
	if len(validways) == 0 {
		log.Fatal("ERROR: no way found ")
	}
	file, er := os.Open(os.Args[1])
	if er != nil {
		log.Fatal(err)
	}
	graph, er := io.ReadAll(file)
	if er != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(graph)
	fmt.Print("\n\n")
	lem.Sendants()
}
