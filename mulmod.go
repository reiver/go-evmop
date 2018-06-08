package evmop

import (
	"io"
)

// MulMod represents the "MULMOD" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.MulMod{}.WriteTo(w)
type MulMod struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (MulMod) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, code_MULMOD)
}
