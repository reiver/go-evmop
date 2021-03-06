package evmop

import (
	"io"
)

// Gas represents the "GAS" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Gas{}.WriteTo(w)
type Gas struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Gas) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeGas)
}
