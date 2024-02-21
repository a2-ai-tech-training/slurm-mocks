package reader_test

import (
	"testing"

	"github.com/a2-ai-tech-training/slurm-mocks/internal/reader"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	res, err := reader.GetScenarioOutput("./testdata/scenario1")
	assert.Nil(t, err)
	assert.Equal(t, []byte("hello from stdout!\n"), res.Stdout)
	assert.Equal(t, []byte("hello from stderr!\n"), res.Stderr)
	assert.Equal(t, 0, res.ExitCode)
}

func TestReadNoStderr(t *testing.T) {
	res, err := reader.GetScenarioOutput("./testdata/scenario2")
	var noBytes []byte
	assert.Nil(t, err)
	assert.Equal(t, []byte("hello from stdout!\n"), res.Stdout)
	assert.Equal(t, noBytes, res.Stderr)
	assert.Equal(t, 0, res.ExitCode)
}

func TestReadNoStdout(t *testing.T) {
	res, err := reader.GetScenarioOutput("./testdata/scenario3")
	var noBytes []byte
	assert.Nil(t, err)
	assert.Equal(t, []byte("hello from stderr!\n"), res.Stderr)
	assert.Equal(t, noBytes, res.Stdout)
	assert.Equal(t, 1, res.ExitCode)
}
