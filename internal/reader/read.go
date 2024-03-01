package reader

import (
	"embed"
	"encoding/json"
	"errors"
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
		// we are ok if stdout is not present and would just expect to print nothing
		if !errors.Is(err, os.ErrNotExist) {
			return ScenarioOutput{}, err
		}
	}
	stderr, err := os.ReadFile(filepath.Join(path, "stderr.txt"))
	if err != nil {
		// we are ok if stderr is not present and would just expect to print nothing
		if !errors.Is(err, os.ErrNotExist) {
			return ScenarioOutput{}, err
		}
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

func GetEmbeddedScenarioOutput(path string, fs embed.FS) (ScenarioOutput, error) {
	stdout, err := fs.ReadFile(filepath.Join(path, "stdout.txt"))
	if err != nil {
		// we are ok if stdout is not present and would just expect to print nothing
		if !errors.Is(err, os.ErrNotExist) {
			return ScenarioOutput{}, err
		}
	}
	stderr, err := fs.ReadFile(filepath.Join(path, "stderr.txt"))
	if err != nil {
		// we are ok if stderr is not present and would just expect to print nothing
		if !errors.Is(err, os.ErrNotExist) {
			return ScenarioOutput{}, err
		}
	}
	metadata, err := fs.ReadFile(filepath.Join(path, "metadata.json"))
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