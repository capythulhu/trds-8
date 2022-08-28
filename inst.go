package trds16

// Build instruction
func I(op byte, arg ...int8) uint16 {
	if len(arg) > 0 {
		return uint16(op)<<8 | uint16(byte(arg[0]))
	}
	return uint16(op) << 8
}

// Get operation from instruction
func Op(inst uint16) byte {
	return byte(inst >> 8)
}

// Get value from instruction
func Val(inst uint16) int8 {
	return int8(inst)
}
