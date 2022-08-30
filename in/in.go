package in

// Operation
func O(o Operation) byte {
	return byte(o)
}

// One register argument
func R1(reg Register) byte {
	return byte(reg)
}

// Two register argument
func R2(reg1, reg2 Register) byte {
	return byte(reg1)<<4 | byte(reg2)
}

// Signed Value
func S(value int8) byte {
	return byte(value)
}

// Unsigned value
func V(value uint8) byte {
	return byte(value)
}
