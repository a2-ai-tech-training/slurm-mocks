package reader

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ScenarioOutput struct {
	Stderr   []byte
	Stdout   []byte
	ExitCode int
}

type ScenarioMetadata struct {
	Command  string `json:"command"`
	ExitCode int    `json:"exit_code"`
}

// GetScenarioOutput reads in the outputs from a given scenario
// it expects to find a stderr.txt, stdout.txt, and metadata.json file
func GetScenarioOutput(path string) (ScenarioOutput, error) {
	stdout, err := os.ReadFile(filepath.Join(path, "stdout.txt"))
	if err != nil {
		return ScenarioOutput{}, err
	}
	stderr, err := os.ReadFile(filepath.Join(path, "stderr.txt"))
	if err != nil {
		return ScenarioOutput{}, err
	}
	metadata, err := os.ReadFile(filepath.Join(path, "metadata.json"))
	if err != nil {
		return ScenarioOutput{}, err
	}
	var scenarioMetadata ScenarioMetadata
	err = json.Unmarshal(metadata, &scenarioMetadata)
	if err != nil {
		return ScenarioOutput{}, err
	}

	return ScenarioOutput{
		Stderr:   stderr,
		Stdout:   stdout,
		ExitCode: scenarioMetadata.ExitCode,
	}, nil
}
