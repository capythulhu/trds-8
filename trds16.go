package trds16

import (
	"fmt"

	"github.com/thzoid/trds-16/in"
)

func fetch(program []byte, pc *byte, it *int) byte {
	*pc++
	return program[*pc]
}

func regPair(b byte) (reg1, reg2 in.Register) {
	reg1 = in.Register(b >> 4)
	reg2 = in.Register(b)
	return
}

func Run(program []byte, latchesU, latchesV map[byte]int8) (code int8, it int) {
	var regs = make([]int8, 4)
	var flags byte = 0 // 000000NZ

	// Opened latches
	var openedULatch *byte
	var openedVLatch *byte

	// Loin
	for i := byte(0); i < byte(len(program)); i, it = i+1, it+1 {
		switch in.Operation(program[i]) {
		// Special
		case in.NOOP:
		case in.HALT:
			return int8(program[fetch(program, &i, &it)]), it
		// Flow control
		case in.JUMP:
			i = fetch(program, &i, &it)
		case in.JUMP_N:
			n := fetch(program, &i, &it)
			if getALUFlag(flags, flag_N) {
				i = n - 1
			}
		case in.JUMP_Z:
			n := fetch(program, &i, &it)
			if getALUFlag(flags, flag_Z) {
				i = n - 1
			}
		case in.JUMP_P:
			n := fetch(program, &i, &it)
			if !getALUFlag(flags, flag_N) {
				i = n - 1
			}
		// Math Operations
		case in.ADD:
			reg1, reg2 := regPair(fetch(program, &i, &it))
			regs[reg1] += regs[reg2]
			setALUFlags(&flags, regs[reg1])
		case in.SUB:
			reg1, reg2 := regPair(fetch(program, &i, &it))
			regs[reg1] -= regs[reg2]
			setALUFlags(&flags, regs[reg1])
		case in.MUL:
			reg1, reg2 := regPair(fetch(program, &i, &it))
			regs[reg1] *= regs[reg2]
			setALUFlags(&flags, regs[reg1])
		case in.DIV:
			reg1, reg2 := regPair(fetch(program, &i, &it))
			regs[reg1] /= regs[reg2]
			setALUFlags(&flags, regs[reg1])
		// Logical Operations
		case in.NOT:
			reg := fetch(program, &i, &it)
			regs[reg] = ^regs[reg]
			setALUFlags(&flags, regs[reg])
		case in.AND:
			reg1, reg2 := regPair(fetch(program, &i, &it))
			regs[reg1] &= regs[reg2]
			setALUFlags(&flags, regs[reg1])
		case in.OR:
			reg1, reg2 := regPair(fetch(program, &i, &it))
			regs[reg1] |= regs[reg2]
			setALUFlags(&flags, regs[reg1])
		case in.XOR:
			reg1, reg2 := regPair(fetch(program, &i, &it))
			regs[reg1] ^= regs[reg2]
			setALUFlags(&flags, regs[reg1])
		// Data Control
		case in.STORE_A:
			program[fetch(program, &i, &it)] |= byte(regs[in.REG_A])
		case in.STORE_B:
			program[fetch(program, &i, &it)] |= byte(regs[in.REG_B])
		case in.STORE_U:
			program[fetch(program, &i, &it)] |= byte(regs[in.REG_U])
		case in.STORE_V:
			program[fetch(program, &i, &it)] |= byte(regs[in.REG_V])
		case in.LOAD_A:
			regs[in.REG_A] = int8(program[fetch(program, &i, &it)])
		case in.LOAD_B:
			regs[in.REG_B] = int8(program[fetch(program, &i, &it)])
		case in.LOAD_U:
			a := fetch(program, &i, &it)
			regs[in.REG_U] = int8(program[a])
		case in.LOAD_V:
			regs[in.REG_V] = int8(program[fetch(program, &i, &it)])
		// Temporal control
		case in.OPEN_U:
			if _, ok := latchesU[i]; !ok {
				latchesU[i] = 0
			}
			if openedULatch == nil {
				openedULatch = new(byte)
				*openedULatch = i
				regs[in.REG_U] = latchesU[i]
			} else {
				panic(fmt.Errorf("attempt to open a U latch that is already opened. inruction: %d, iteration: %d", i, it))
			}
		case in.OPEN_V:
			if _, ok := latchesV[i]; !ok {
				latchesV[i] = 0
			}
			if openedVLatch == nil {
				openedVLatch = new(byte)
				*openedVLatch = i
				regs[in.REG_V] = latchesV[i]
			} else {
				panic(fmt.Errorf("attempt to open a V latch that is already opened. inruction: %d, iteration: %d", i, it))
			}
		case in.CLOSE_U:
			if openedULatch == nil {
				panic(fmt.Errorf("attempt to close a U latch that is already closed. inruction: %d, iteration: %d", i, it))
			}
			latchesU[*openedULatch] = regs[in.REG_U]
			openedULatch = nil
		case in.CLOSE_V:
			if openedVLatch == nil {
				panic(fmt.Errorf("attempt to close a V latch that is already closed. inruction: %d, iteration: %d", i, it))
			}
			latchesV[*openedVLatch] = regs[in.REG_V]
			openedVLatch = nil
		default:
			panic(fmt.Errorf("unknown inruction: %d. inruction: %d, iteration: %d", program[i], i, it))
		}
	}
	return 0, it
}

func RunTemporal(program []byte, steps uint) (results []int8, it []int) {
	results = make([]int8, steps)
	it = make([]int, steps)
	latchesU, latchesV := make(map[byte]int8), make(map[byte]int8)
	for i := uint(0); i < steps; i++ {
		p := make([]byte, len(program))
		copy(p, program)
		results[i], it[i] = Run(p, latchesU, latchesV)
	}
	return results, it
}
