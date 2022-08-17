package main

import (
	"math"

	"github.com/thzoid/trds-16/ctrl"
	"github.com/thzoid/trds-16/mem"
)

func main() {
	// TRDS-16 has a 16-bit architecture
	const bits = 16

	// Create Classical Registers
	A := mem.NewRegister(bits / 2)
	B := mem.NewRegister(bits / 2)

	// Create RAM with address bus size = data bus size / 2
	ram := mem.NewRAM(bits, uint(math.Pow(2, bits/2)))

	// Create Control Unit with specified data bus size
	controlUnit := ctrl.NewControlUnit(ctrl.Setup{})
	// Connect Control Unit to RAM
	controlUnit.RAM = ram
}
