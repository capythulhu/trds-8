package trds8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thzoid/trds-8/in"
)

func TestSimpleTemporalBranchProgram(t *testing.T) {
	program := []byte{
		in.O(in.OPEN_U),
		in.O(in.STORE_U),
		in.V(0x14),
		in.O(in.LOAD_A),
		in.V(0x14),
		in.O(in.LOAD_B),
		in.V(0x12),
		in.O(in.SUB),
		in.R2(in.REG_A, in.REG_B),
		in.O(in.JUMP_Z),
		in.V(0xD),
		in.O(in.HALT),
		in.V(0x13),
		in.O(in.LOAD_U),
		in.V(0x13),
		in.O(in.CLOSE_U),
		in.O(in.HALT),
		in.V(0x12),
		in.S(0),
		in.S(1),
		in.S(0),
	}
	results, _ := RunTemporal(program, 2)
	assert.Equal(t, []int8{0, 1}, results)
}

func TestTemporalParadoxicalProgram(t *testing.T) {
	program := []byte{
		in.O(in.OPEN_U),
		in.O(in.STORE_U),
		in.V(0x17),
		in.O(in.LOAD_A),
		in.V(0x17),
		in.O(in.LOAD_B),
		in.V(0x15),
		in.O(in.SUB),
		in.R2(in.REG_A, in.REG_B),
		in.O(in.JUMP_Z),
		in.V(0x10),
		in.O(in.LOAD_U),
		in.V(0x15),
		in.O(in.CLOSE_U),
		in.O(in.HALT),
		in.V(0x15),
		in.O(in.LOAD_U),
		in.V(0x16),
		in.O(in.CLOSE_U),
		in.O(in.HALT),
		in.V(0x16),
		in.S(0),
		in.S(1),
		in.S(0),
	}
	results, _ := RunTemporal(program, 4)
	assert.Equal(t, []int8{1, 0, 1, 0}, results)
}
