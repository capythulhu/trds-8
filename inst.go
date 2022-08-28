package trds16

// Int8 to byte
func abs(n int8) byte {
	if n < 0 {
		return byte(-n)
	}
	return byte(n)
}

// Build instruction
func I(op byte, arg ...int8) uint16 {
	if len(arg) > 0 {
		return uint16(op)<<8 | uint16(arg[0])
	}
	return uint16(op) << 8
}

// Get operation from instruction
func Op(inst uint16) byte {
	return abs(int8(inst >> 8))
}

// Get value from instruction
func Val(inst uint16) int8 {
	return int8(inst)
}
