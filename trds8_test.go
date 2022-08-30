package trds8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thzoid/trds-8/in"
)

func TestSimpleTemporalBranchProgram(t *testing.T) {
	program := []byte{
		in.O(in.OPEN_U),
		in.O(in.LOAD_A),
		in.V(0xE),
		in.O(in.SUB),
		in.R2(in.REG_U, in.REG_A),
		in.O(in.JUMP_Z),
		in.V(0x9),
		in.O(in.HALT),
		in.V(0xF),
		in.O(in.LOAD_U),
		in.V(0xF),
		in.O(in.CLOSE_U),
		in.O(in.HALT),
		in.V(0xE),
		in.S(0),
		in.S(1),
	}
	results, _ := RunTemporal(program, 2)
	assert.Equal(t, []int8{0, 1}, results)
}

func TestTemporalParadoxicalProgram(t *testing.T) {
	program := []byte{
		in.O(in.OPEN_U),
		in.O(in.LOAD_A),
		in.V(0x11),
		in.O(in.SUB),
		in.R2(in.REG_U, in.REG_A),
		in.O(in.JUMP_Z),
		in.V(0xC),
		in.O(in.LOAD_U),
		in.V(0x11),
		in.O(in.CLOSE_U),
		in.O(in.HALT),
		in.V(0x11),
		in.O(in.LOAD_U),
		in.V(0x12),
		in.O(in.CLOSE_U),
		in.O(in.HALT),
		in.V(0x12),
		in.S(0),
		in.S(1),
	}
	results, _ := RunTemporal(program, 4)
	assert.Equal(t, []int8{1, 0, 1, 0}, results)
}

func TestTemporalIterationReductionProgram(t *testing.T) {
	// Program
	program := []byte{
		in.O(in.OPEN_U),
		in.O(in.LOAD_A),
		in.V(0x22),
		in.O(in.SUB),
		in.R2(in.REG_U, in.REG_A),
		in.O(in.OPEN_V),
		in.O(in.JUMP_Z),
		in.V(0xC),
		in.O(in.STORE_V),
		in.V(0x26),
		in.O(in.HALT),
		in.V(0x26),
		in.O(in.LOAD_A),
		in.V(0x24),
		in.O(in.LOAD_B),
		in.V(0x25),
		in.O(in.MUL),
		in.R2(in.REG_A, in.REG_B),
		in.O(in.ADD),
		in.R2(in.REG_A, in.REG_B),
		in.O(in.MUL),
		in.R2(in.REG_A, in.REG_B),
		in.O(in.XOR),
		in.R2(in.REG_A, in.REG_B),
		in.O(in.STORE_A),
		in.V(0x26),
		in.O(in.LOAD_V),
		in.V(0x26),
		in.O(in.CLOSE_V),
		in.O(in.LOAD_U),
		in.V(0x23),
		in.O(in.CLOSE_U),
		in.O(in.HALT),
		in.V(0x26),
		in.S(0),
		in.S(1),
		in.S(-2),
		in.S(3),
		in.S(0),
	}

	results, iterations := RunTemporal(program, 2)
	assert.Equal(t, []int8{-12, -12}, results)
	assert.Equal(t, []int{16, 6}, iterations)
}
