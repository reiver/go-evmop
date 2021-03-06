package evmop

import (
	"io"
)

// Call represents the "CALL" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Call{}.WriteTo(w)
type Call struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Call) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCall)
}
