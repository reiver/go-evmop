package evmop

import (
	"io"
)

// Div represents the "DIV" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Div{}.WriteTo(w)
type Div struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Div) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, code_DIV)
}
