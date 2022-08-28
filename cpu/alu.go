package cpu

const (
	FLAG_Z = iota
	FLAG_N
)

func SetALUFlag(flags *byte, flag byte, val bool) {
	*flags &= ^(1 << flag)
	if val {
		*flags |= 1 << flag
	}
}

func GetALUFlag(flags, flag byte) bool {
	return (flags & 1 << flag) > 0
}

func SetALUFlags(flags *byte, a int8) {
	SetALUFlag(flags, FLAG_Z, a == 0)
	SetALUFlag(flags, FLAG_N, a < 0)
}
