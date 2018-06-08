package evmop

import (
	"io"
)

// Stop represents the "STOP" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Stop{}.WriteTo(w)
type Stop struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Stop) WriteTo(writer io.Writer) (int64, error) {

	const opcode = code_STOP

	if nil == writer {
		return 0, errNilWriter
	}

	var n64 int64

	{
		var b [1]byte = [1]byte{opcode}
		p := b[:]

		n, err := writer.Write(p)
		if nil != err {
			return n64, err
		}
		if 0 > n {
			return n64, errNegativeNumberOfBytesWritten
		}
		n64 += int64(n)
	}

	return n64, nil
}
