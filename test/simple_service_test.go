package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sumitroajiprabowo/go-dependency-injection/example"
)

func TestInizializeSimpleService(t *testing.T) {
	simpleService := example.InizializeSimpleService()
	if simpleService == nil {
		t.Errorf("InizializeSimpleService() returned nil")
	}
	assert.NotNil(t, simpleService)
	assert.Equal(t, simpleService.SimpleRepository, &example.SimpleRepository{})
}

func TestInizializeSimpleServiceWithError_Error(t *testing.T) {
	simpleService, err := example.InizializeSimpleServiceWithError(true)
	if simpleService != nil {
		t.Errorf("InizializeSimpleServiceWithError() returned %v", simpleService)
	}
	if err == nil {
		t.Errorf("InizializeSimpleServiceWithError() returned nil")
	}
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestInizializeSimpleServiceWithError_NoError(t *testing.T) {
	simpleService, err := example.InizializeSimpleServiceWithError(false)
	if simpleService == nil {
		t.Errorf("InizializeSimpleServiceWithError() returned nil")
	}
	if err != nil {
		t.Errorf("InizializeSimpleServiceWithError() returned %v", err)
	}
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}
