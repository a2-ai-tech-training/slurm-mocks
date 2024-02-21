// https://slurm.schedmd.com/sinfo.html
package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/a2-ai-tech-training/slurm-mocks/internal/hasher"
	"github.com/a2-ai-tech-training/slurm-mocks/internal/reader"
)

//go:embed outputs/*
var f embed.FS

func main() {

	arguments := os.Args[1:]
	hash := hasher.Hasher(arguments)

	scenario := os.Getenv("SCENARIO")
	if scenario == "" {
		scenario = "default"
	}

	hash_path := fmt.Sprintf("outputs/%s/%s/", scenario, hash)

	file, err := reader.GetScenarioOutput(hash_path)

	if err != nil {
		log.Fatal(err)
	}
	/*
		fmt.Printf("Stderr:\n%s\n", file.Stderr)
		fmt.Printf("Stdout:\n%s\n", file.Stdout)
		fmt.Printf("Exit Code:\n%d\n", file.ExitCode)*/
	fmt.Fprintf(os.Stdout, "%s", file.Stdout)
	fmt.Fprintf(os.Stderr, "%s", file.Stderr)
	os.Exit(file.ExitCode)
}
