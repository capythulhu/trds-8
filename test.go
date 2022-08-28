package trds16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTemporalBranchProgram(t *testing.T) {
	program := []int16{
		Inst(NOOP, 0),
		Inst(NOOP, 1),
		Inst(NOOP, 0),
		Inst(OPEN_U),
		Inst(STORE_U, 0x2),
		Inst(LOAD_A, 0x2),
		Inst(LOAD_B, 0x0),
		Inst(SUB),
		Inst(JUMP_Z, 0x9),
		Inst(HALT, 0x1),
		Inst(LOAD_U, 0x1),
		Inst(CLOSE_U),
		Inst(HALT, 0x0),
	}
	assert.Equal(t, RunTemporal(program, 2), []int8{0, 1})
}

func TestParadoxTemporaProgram(t *testing.T) {
	program := []int16{
		Inst(NOOP, 0),
		Inst(NOOP, 1),
		Inst(NOOP, 0),
		Inst(OPEN_U),
		Inst(STORE_U, 0x2),
		Inst(LOAD_A, 0x2),
		Inst(LOAD_B, 0x0),
		Inst(SUB),
		Inst(JUMP_Z, 0xB),
		Inst(LOAD_U, 0x0),
		Inst(CLOSE_U),
		Inst(HALT, 0x0),
		Inst(LOAD_U, 0x1),
		Inst(CLOSE_U),
		Inst(HALT, 0x1),
	}
	assert.Equal(t, RunTemporal(program, 10), []int8{1, 0, 1, 0, 1, 0, 1, 0, 1, 0})
}
