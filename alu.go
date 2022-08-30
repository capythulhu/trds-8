package trds8

const (
	flag_Z = iota
	flag_N
)

func setALUFlag(flags *byte, flag byte, val bool) {
	*flags &= ^(1 << flag)
	if val {
		*flags |= 1 << flag
	}
}

func getALUFlag(flags, flag byte) bool {
	return (flags & 1 << flag) > 0
}

func setALUFlags(flags *byte, reg int8) {
	setALUFlag(flags, flag_Z, reg == 0)
	setALUFlag(flags, flag_N, reg < 0)
}
