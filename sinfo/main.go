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

//go:embed outputs/**/*
var f embed.FS

func main() {

	output_path := os.Getenv("OUTPUT_DIR")
	if output_path == "" {
		output_path = "outputs"
	}

	scenario := os.Getenv("SCENARIO")
	if scenario == "" {
		scenario = "default"
	}

	var file reader.ScenarioOutput
	var err error
	arguments := os.Args[1:]
	hash := hasher.Hasher(arguments)

	hash_path := fmt.Sprintf("%s/%s/%s/", output_path, scenario, hash)

	if os.Getenv("OUTPUT_DIR") == "" {
		file, err = reader.GetEmbeddedScenarioOutput(hash_path, f)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = reader.GetScenarioOutput(hash_path)

		if err != nil {
			log.Fatal(err)
		}
	}
	
	fmt.Fprintf(os.Stdout, "%s", file.Stdout)
	fmt.Fprintf(os.Stderr, "%s", file.Stderr)
	os.Exit(file.ExitCode)
}
