package trds16

func Run(program []int16, latchesU, latchesV map[byte]int8) int8 {
	// Registers
	var a, b, u, v int8 = 0, 0, 0, 0
	var flags byte = 0 // 000000NZ

	// Opened latches
	var openedULatch *byte
	var openedVLatch *byte

	// Loop
	for i := byte(0); i < byte(len(program)); i++ {
		op, val := Op(program[i]), Val(program[i])
		switch op {
		// Special
		case NOOP:
		case HALT:
			return int8(program[val])
		// Flow control
		case JUMP:
			i = byte(val)
		case JUMP_N:
			if GetALUFlag(flags, FLAG_N) {
				i = byte(val)
			}
		case JUMP_Z:
			if GetALUFlag(flags, FLAG_Z) {
				i = byte(val)
			}
		case JUMP_P:
			if GetALUFlag(flags, FLAG_N) {
				i = byte(val)
			}
		// Math Operations
		case ADD:
			a += b
			SetALUFlags(&flags, a)
		case SUB:
			a -= b
			SetALUFlags(&flags, a)
		case MUL:
			a *= b
			SetALUFlags(&flags, a)
		case DIV:
			a *= b
			SetALUFlags(&flags, a)
		// Logical Operations
		case NOT:
			a = ^a
			SetALUFlags(&flags, a)
		case AND:
			a &= b
			SetALUFlags(&flags, a)
		case OR:
			a |= b
			SetALUFlags(&flags, a)
		case XOR:
			a ^= b
			SetALUFlags(&flags, a)
		// Data Control
		case STORE_A:
			program[val] |= int16(a)
		case STORE_B:
			program[val] |= int16(b)
		case STORE_U:
			program[val] |= int16(u)
		case STORE_V:
			program[val] |= int16(v)
		case LOAD_A:
			a = int8(program[val])
		case LOAD_B:
			b = int8(program[val])
		case LOAD_U:
			u = int8(program[val])
		case LOAD_V:
			v = int8(program[val])
		// Temporal control
		case OPEN_U:
			if _, ok := latchesU[i]; !ok {
				latchesU[i] = 0
			}
			if openedULatch == nil {
				openedULatch = new(byte)
				*openedULatch = i
				u = latchesU[i]
			} else {
				panic("attempt to open a U latch that is already opened")
			}
		case OPEN_V:
			if _, ok := latchesV[i]; !ok {
				latchesV[i] = 0
			}
			if openedVLatch == nil {
				openedVLatch = new(byte)
				*openedVLatch = i
				v = latchesV[i]
			} else {
				panic("attempt to open a V latch that is already opened")
			}
		case CLOSE_U:
			if openedULatch == nil {
				panic("attempt to close a U latch that is already closed")
			}
			latchesU[*openedULatch] = u
			openedULatch = nil
		case CLOSE_V:
			if openedVLatch == nil {
				panic("attempt to close a V latch that is already closed")
			}
			latchesV[*openedVLatch] = v
			openedVLatch = nil
		default:
			panic("unknown instruction")
		}
	}
	return 0
}

func RunTemporal(program []int16, steps uint) []int8 {
	results := make([]int8, steps)
	latchesU, latchesV := make(map[byte]int8), make(map[byte]int8)
	for i := uint(0); i < steps; i++ {
		p := make([]int16, len(program))
		copy(p, program)
		results[i] = Run(p, latchesU, latchesV)
	}
	return results
}
