package trds16

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thzoid/trds-16/op"
)

func TestSimpleTemporalBranchProgram(t *testing.T) {
	program := []int16{
		I(op.NOOP, 0),
		I(op.NOOP, 1),
		I(op.NOOP, 0),
		I(op.OPEN_U),
		I(op.STORE_U, 0x2),
		I(op.LOAD_A, 0x2),
		I(op.LOAD_B, 0x0),
		I(op.SUB),
		I(op.JUMP_Z, 0x9),
		I(op.HALT, 0x1),
		I(op.LOAD_U, 0x1),
		I(op.CLOSE_U),
		I(op.HALT, 0x0),
	}
	results, _ := RunTemporal(program, 2)
	assert.Equal(t, results, []int8{0, 1})
}

func TestParadoxTemporaProgram(t *testing.T) {
	program := []int16{
		I(op.NOOP, 0),
		I(op.NOOP, 1),
		I(op.NOOP, 0),
		I(op.OPEN_U),
		I(op.STORE_U, 0x2),
		I(op.LOAD_A, 0x2),
		I(op.LOAD_B, 0x0),
		I(op.SUB),
		I(op.JUMP_Z, 0xB),
		I(op.LOAD_U, 0x0),
		I(op.CLOSE_U),
		I(op.HALT, 0x0),
		I(op.LOAD_U, 0x1),
		I(op.CLOSE_U),
		I(op.HALT, 0x1),
	}
	results, _ := RunTemporal(program, 10)
	assert.Equal(t, results, []int8{1, 0, 1, 0, 1, 0, 1, 0, 1, 0})
}
