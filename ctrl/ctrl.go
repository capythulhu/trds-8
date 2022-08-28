package ctrl

import (
	"github.com/thzoid/trds-16/alu"
	"github.com/thzoid/trds-16/mem"
)

// Control Unit
type ControlUnit struct {
	// Random Access Memory
	RAM *mem.RAM
	// Inner Registers
	PC, ACC, CI *mem.Register
	// Outer Registers
	Registers []*mem.Register
	// Arithmetic Logic Unit
	ALU *alu.ArithmeticLogicUnit
}

// Create new Control Unit
func NewControlUnit(setup Setup) *ControlUnit {
	return &ControlUnit{}
}

func (cU *ControlUnit) Tick() {

}
