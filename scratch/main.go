package main

import (
	"fmt"
	"os"

	"github.com/a2-ai-tech-training/slurm-mocks/internal/hasher"
)

func main() {

	scenario := os.Getenv("SCENARIO")
	fmt.Printf("Before:\nos $SCENARIO: %s\n", scenario)
	if scenario == "" {
		scenario = "default"
	}
	fmt.Printf("After:        %s\n", scenario)
	//fmt.Printf("%s\n", os.Args)
	fmt.Printf("%s\n", os.Args[1:])
	fmt.Printf("hash: %s\n", hasher.Hasher(os.Args[1:]))
}
