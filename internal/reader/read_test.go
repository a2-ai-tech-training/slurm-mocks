package reader_test

import (
	"testing"

	"github.com/a2-ai-tech-training/slurm-mocks/internal/reader"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {

	res, err := reader.GetScenarioOutput("./testdata/scenario1")
	t.Log(string(res.Stderr))
	assert.Nil(t, err)
	assert.Equal(t, []byte("hello from stdout!\n"), res.Stdout)
	assert.Equal(t, []byte("hello from stderr!\n"), res.Stderr)
}
