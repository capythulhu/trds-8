package mem

// Random Access Memory
type RAM struct {
	memory      []byte
	dataBusSize uint
}

// Create new RAM
func NewRAM(size, dataBusSize uint) *RAM {
	return &RAM{
		memory:      make([]byte, size),
		dataBusSize: dataBusSize,
	}
}

// Read bytes from RAM
func (r *RAM) Read(addr uint) (data []byte) {
	if addr+r.dataBusSize >= uint(len(r.memory)) {
		panic("ram: attempt to read invalid address from random access memory")
	}

	return r.memory[addr : addr+r.dataBusSize]
}

// Write bytes to RAM
func (r *RAM) Write(addr uint, data []byte) {
	if len(data) != int(r.dataBusSize) {
		panic("ram: attempt to write invalid quantity of data")
	}

	if addr+r.dataBusSize >= uint(len(r.memory)) {
		panic("ram: attempt to write to invalid address")
	}

	for i := addr; i < addr+r.dataBusSize; i++ {
		r.memory[i] = data[i]
	}
}
