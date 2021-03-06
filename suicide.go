package evmop

import (
	"io"
)

// Suicide represents the "SUICIDE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Suicide{}.WriteTo(w)
type Suicide struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Suicide) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSuicide)
}
