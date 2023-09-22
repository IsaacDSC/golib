package l4g

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogError(t *testing.T) {
	exist := ManageError(
		errors.New("Generate error"),
		EnableTracing, RepoLayer, RecService,
	)
	assert.True(t, exist)
}
