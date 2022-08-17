package mem

// Register
type Register struct {
	memory []byte
}

// Create new Register
func NewRegister(size uint) *Register {
	return &Register{memory: make([]byte, size)}
}

// Write to Register
func (r *Register) Write(data []byte) {
	if len(data) != len(r.memory) {
		panic("reg: attempt to write invalid quantity of data")
	}

	copy(r.memory, data)
}

// Read from Register
func (r *Register) Read() (data []byte) {
	data = make([]byte, len(r.memory))
	copy(data, r.memory)

	return
}
