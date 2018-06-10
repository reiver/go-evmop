package evmop

import (
	"io"
)

// Dup13 represents the "DUP13" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup13{}.WriteTo(w)
type Dup13 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup13) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup13)
}
