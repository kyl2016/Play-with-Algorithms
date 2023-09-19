package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestResort(t *testing.T) {
	assert.Equal(t, resort("", 1), "")
	assert.Equal(t, resort("1234567890123", 0), "")
	assert.Equal(t, resort("1234567890123", 1), "1234567890123")
	assert.Equal(t, resort("1234567890123", 2), "135791246802")
	assert.Equal(t, resort("1234567890123", 100), "1234567890123")
	assert.Equal(t, resort("123", 3), "123")

	assert.Equal(t, resort("1234567890123", 3), "1593246802371")
	assert.Equal(t, resort("1234567890123", 4), "1732682359140")
}
