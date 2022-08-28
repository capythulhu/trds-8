package alu

import (
	"bytes"
	"encoding/binary"
	"math"
)

var (
	SIMPLE = Setup{
		Operations: []func(*ArithmeticLogicUnit, []byte, []byte) ([]byte, bool, bool, bool){
			// ADD
			func(aLU *ArithmeticLogicUnit, v1, v2 []byte) (result []byte, overflow, zero, negative bool) {
				i1 := binary.BigEndian.Uint32(v1)
				i2 := binary.BigEndian.Uint32(v2)

				res := int(i1) + int(i2)

				zero = res == 0
				negative = res < 0
				overflow = int(i2) > (math.MaxInt-int(i1)) || int(i2) < (math.MaxInt-int(i1))

				buf := new(bytes.Buffer)
				binary.Write(buf, binary.BigEndian, res)
				result = buf.Bytes()

				return
			},
		},
	}
)

// Arithmetic-Logic Unit Setup
type Setup struct {
	Size       uint
	Operations []func(aLU *ArithmeticLogicUnit, v1, v2 []byte) (result []byte, overflow, zero, negative bool)
}
