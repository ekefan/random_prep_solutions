package main

import (
	"log"
	"os"
	"solve-algo/solver"
)

func main() {
	pathToFiles := os.Args[1:]

	if len(pathToFiles) != 1 {
		log.Fatal("must pass directory to files")
	}

	err := solver.RunSolver(pathToFiles[0])
	if err != nil {
		log.Fatal(err)
	}
}
