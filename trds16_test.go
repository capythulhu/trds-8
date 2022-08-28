package trds16

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thzoid/trds-16/cpu"
	"github.com/thzoid/trds-16/op"
)

func TestSimpleTemporalBranchProgram(t *testing.T) {
	program := []uint16{
		cpu.I(op.NOOP, 0),
		cpu.I(op.NOOP, 1),
		cpu.I(op.NOOP, 0),
		cpu.I(op.OPEN_U),
		cpu.I(op.STORE_U, 0x2),
		cpu.I(op.LOAD_A, 0x2),
		cpu.I(op.LOAD_B, 0x0),
		cpu.I(op.SUB),
		cpu.I(op.JUMP_Z, 0x9),
		cpu.I(op.HALT, 0x1),
		cpu.I(op.LOAD_U, 0x1),
		cpu.I(op.CLOSE_U),
		cpu.I(op.HALT, 0x0),
	}
	results, _ := RunTemporal(program, 2)
	assert.Equal(t, results, []int8{0, 1})
}

func TestParadoxTemporaProgram(t *testing.T) {
	program := []uint16{
		cpu.I(op.NOOP, 0),
		cpu.I(op.NOOP, 1),
		cpu.I(op.NOOP, 0),
		cpu.I(op.OPEN_U),
		cpu.I(op.STORE_U, 0x2),
		cpu.I(op.LOAD_A, 0x2),
		cpu.I(op.LOAD_B, 0x0),
		cpu.I(op.SUB),
		cpu.I(op.JUMP_Z, 0xB),
		cpu.I(op.LOAD_U, 0x0),
		cpu.I(op.CLOSE_U),
		cpu.I(op.HALT, 0x0),
		cpu.I(op.LOAD_U, 0x1),
		cpu.I(op.CLOSE_U),
		cpu.I(op.HALT, 0x1),
	}
	results, _ := RunTemporal(program, 10)
	assert.Equal(t, results, []int8{1, 0, 1, 0, 1, 0, 1, 0, 1, 0})
}
