package main

import (
	"fmt"
	"os"
)

func main() {
	scenario := os.Getenv("SCENARIO")
	fmt.Printf("Before:\nos $SCENARIO: %s\n", scenario)
	if scenario == "" {
		scenario = "default"
	}
	fmt.Printf("After:        %s\n", scenario)
}
