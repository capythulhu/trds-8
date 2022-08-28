package trds16

func I(op byte, arg ...int8) int16 {
	if len(arg) > 0 {
		return int16(op)<<8 | int16(arg[0])
	}
	return int16(op) << 8
}

func Op(inst int16) byte {
	return byte(inst >> 8)
}

func Val(inst int16) int8 {
	return int8(inst)
}
