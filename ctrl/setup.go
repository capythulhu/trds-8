package ctrl

import "encoding/binary"

var (
	SIMPLE16 = Setup{
		Size:      16,
		Registers: 4,
		Operations: []func(cU *ControlUnit, data []byte){
			// NOOP
			func(cU *ControlUnit, data []byte) {},
			// HALT
			func(cU *ControlUnit, data []byte) {},
			// JUMP
			func(cU *ControlUnit, data []byte) {
				cU.PC.Write(data)
			},
			// LOAD_A
			func(cU *ControlUnit, data []byte) {
				cU.Registers[0].Write(data)
			},
			// STORE_A
			func(cU *ControlUnit, data []byte) {
				cU.RAM.Write(
					uint(binary.BigEndian.Uint32(data)),
					cU.Registers[0].Read(),
				)
			},
			// ADD
			func(cU *ControlUnit, data []byte) {
			},
		},
	}
	TRDS16 = Setup{}
)

// Control Unit Setup
type Setup struct {
	Size       uint
	Registers  uint
	Operations []func(cU *ControlUnit, data []byte)
}
