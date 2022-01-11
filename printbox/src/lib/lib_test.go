package lib_test

import (
	"start/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBoard(t *testing.T) {
	assert := assert.New(t)

	lib.CpuInfoFile = "./testdata/cpuinfo_pi3"

	resp, err := lib.GetBoard()
	assert.Nil(err)
	assert.NotNil(resp)

	assert.Equal("BCM2835", resp.Hardware)
	assert.Equal("a020d3", resp.Revision)
	assert.Equal("00000000fffffff", resp.Serial)
	assert.Equal("Raspberry Pi 3 Model B Plus Rev 1.3", resp.Model)
}
