package rpi_test

import (
	"printbox/internal/rpi"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBoard(t *testing.T) {
	assert := assert.New(t)

	rpi.CpuInfoFile = "./testdata/cpuinfo_pi3b"

	resp, err := rpi.Identify()
	assert.Nil(err)
	assert.NotNil(resp)
	assert.Equal("BCM2835", resp.Hardware)
	assert.Equal("a020d3", resp.Revision)
	assert.Equal("00000000fffffff", resp.Serial)
	assert.Equal("Raspberry Pi 3 Model B Plus Rev 1.3", resp.Model)

	rpi.CpuInfoFile = "./testdata/cpuinfo_pi4b"

	resp, err = rpi.Identify()
	assert.Nil(err)
	assert.NotNil(resp)
	assert.Equal("BCM2711", resp.Hardware)
	assert.Equal("c03111", resp.Revision)
	assert.Equal("fffffffffpi4pi4", resp.Serial)
	assert.Equal("Raspberry Pi 4 Model B Rev 1.1", resp.Model)

}
