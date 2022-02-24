package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sumitroajiprabowo/go-dependency-injection/example"
)

func TestInizializedCleanup(t *testing.T) {

	filesystem, cleanup := example.InizializedCleanup("test.txt")

	if filesystem == nil {
		t.Errorf("InizializedCleanup() returned nil")
	}

	if cleanup == nil {
		t.Errorf("InizializedCleanup() returned nil")
	}

	assert.NotNil(t, filesystem)
	assert.NotNil(t, cleanup)

	cleanup()

}
