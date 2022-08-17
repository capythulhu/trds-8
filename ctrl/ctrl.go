package ctrl

import "github.com/thzoid/trds-16/mem"

var (
	SIMPLE16 = Setup{}
	TRDS16   = Setup{}
)

// Control Unit Setup
type Setup struct {
	Operations []func(cU *ControlUnit)
}

// Control Unit
type ControlUnit struct {
	// Random Access Memory
	RAM *mem.RAM
	// Registers
	PC, ACC, IC *mem.Register
}

// Create new Control Unit
func NewControlUnit(setup Setup) *ControlUnit {
	return &ControlUnit{}
}

func (cU *ControlUnit) Tick() {

}
