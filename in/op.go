package in

type Operation byte

const (
	// Special
	NOOP Operation = iota
	HALT

	// Flow Control
	JUMP
	JUMP_N
	JUMP_Z
	JUMP_P

	// Mathematical Operations
	ADD
	SUB
	MUL
	DIV

	// Bitwise Operations
	NOT
	AND
	OR
	XOR

	// Data Control
	LOAD_A
	LOAD_B
	LOAD_U
	LOAD_V
	STORE_A
	STORE_B
	STORE_U
	STORE_V

	// Temporal Control
	OPEN_U
	OPEN_V
	CLOSE_U
	CLOSE_V
)
