/*
Package evmop provides tools for turning Ethereum Virtual Machine (EVM) OpCodes into bytecodes.

This might be useful to someone writing an AOT or JIT compiler targeting the Ethereum Virtual Machine (EVM).


Usage

The following stack-oriented style of assembly for the Ethereum Virtual Machine (EVM):

	push1 3 push1 2 add push1 1 mul

Would be the equivalent to the following Go code:

	var writer io.Writer
	
	// ...
	
	evmop.Push1{3}.WriteTo(writer)
	
	evmop.Push1{2}.WriteTo(writer)
	
	evmop.Add{}.WriteTo(writer)
	
	evmop.Push1{1}.WriteTo(writer)
	
	evmop.Mul{}.WriteTo(writer)

*/
package evmop
