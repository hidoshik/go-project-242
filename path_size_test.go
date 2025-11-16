package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize_File(t *testing.T) {
	pathSize, err := GetSize("./testdata")
	
	assert.NoError(t, err)
	assert.Equal(t, pathSize, "48 ./testdata")
}