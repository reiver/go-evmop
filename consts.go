package evmop

const (
	CodeStop           = byte(0x00)
	CodeAdd            = byte(0x01)
	CodeMul            = byte(0x02)
	CodeSub            = byte(0x03)
	CodeDiv            = byte(0x04)
	CodeSDiv           = byte(0x05)
	CodeMod            = byte(0x06)
	CodeSMod           = byte(0x07)
	CodeAddMod         = byte(0x08)
	CodeMulMod         = byte(0x09)
	CodeExp            = byte(0x0a)
	CodeSignExtend     = byte(0x0b)

	CodeLT             = byte(0x10)
	CodeGT             = byte(0x11)
	CodeSLT            = byte(0x12)
	CodeSGT            = byte(0x13)
	CodeEq             = byte(0x14)
	CodeIsZero         = byte(0x15)
	CodeAnd            = byte(0x16)
	CodeOr             = byte(0x17)
	CodeXor            = byte(0x18)
	CodeNot            = byte(0x19)
	CodeByte           = byte(0x1a)

	CodeSHA3           = byte(0x20)

	CodeAddress        = byte(0x30)
	CodeBalance        = byte(0x31)
	CodeOrigin         = byte(0x32)
	CodeCaller         = byte(0x33)
	CodeCallValue      = byte(0x34)
	CodeCallDataLoad   = byte(0x35)
	CodeCallDataSize   = byte(0x36)
	CodeCallDataCopy   = byte(0x37)
	CodeCodeSize       = byte(0x38)
	CodeCodeCopy       = byte(0x39)
	CodeGasPrice       = byte(0x3a)
	CodeExtCodeSize    = byte(0x3b)
	CodeExtCodeCopy    = byte(0x3c)
	CodeReturnDataSize = byte(0x3d)
	CodeReturnDataCopy = byte(0x3e)

	code_BLOCKHASH      = byte(0x40)
	code_COINBASE       = byte(0x41)
	code_TIMESTAMP      = byte(0x42)
	code_NUMBER         = byte(0x43)
	code_DIFFICULTY     = byte(0x44)
	code_GASLIMIT       = byte(0x45)

	code_POP            = byte(0x50)
	code_MLOAD          = byte(0x51)
	code_MSTORE         = byte(0x52)
	code_MSTORE8        = byte(0x53)
	code_SLOAD          = byte(0x54)
	code_SSTORE         = byte(0x55)
	code_JUMP           = byte(0x56)
	code_JUMPI          = byte(0x57)
	code_PC             = byte(0x58)
	code_MSIZE          = byte(0x59)
	code_GAS            = byte(0x5a)
	code_JUMPDEST       = byte(0x5b)

	code_PUSH1          = byte(0x5f +  1)
	code_PUSH2          = byte(0x5f +  2)
	code_PUSH3          = byte(0x5f +  3)
	code_PUSH4          = byte(0x5f +  4)
	code_PUSH5          = byte(0x5f +  5)
	code_PUSH6          = byte(0x5f +  6)
	code_PUSH7          = byte(0x5f +  7)
	code_PUSH8          = byte(0x5f +  8)
	code_PUSH9          = byte(0x5f +  9)
	code_PUSH10         = byte(0x5f + 10)
	code_PUSH11         = byte(0x5f + 11)
	code_PUSH12         = byte(0x5f + 12)
	code_PUSH13         = byte(0x5f + 13)
	code_PUSH14         = byte(0x5f + 14)
	code_PUSH15         = byte(0x5f + 15)
	code_PUSH16         = byte(0x5f + 16)
	code_PUSH17         = byte(0x5f + 17)
	code_PUSH18         = byte(0x5f + 18)
	code_PUSH19         = byte(0x5f + 19)
	code_PUSH20         = byte(0x5f + 20)
	code_PUSH21         = byte(0x5f + 21)
	code_PUSH22         = byte(0x5f + 22)
	code_PUSH23         = byte(0x5f + 23)
	code_PUSH24         = byte(0x5f + 24)
	code_PUSH25         = byte(0x5f + 25)
	code_PUSH26         = byte(0x5f + 26)
	code_PUSH27         = byte(0x5f + 27)
	code_PUSH28         = byte(0x5f + 28)
	code_PUSH29         = byte(0x5f + 29)
	code_PUSH30         = byte(0x5f + 30)
	code_PUSH31         = byte(0x5f + 31)
	code_PUSH32         = byte(0x5f + 32)
	code_PUSH33         = byte(0x5f + 33)

	code_LOG0           = byte(0xa0)
	code_LOG1           = byte(0xa1)
	code_LOG2           = byte(0xa2)
	code_LOG3           = byte(0xa3)
	code_LOG4           = byte(0xa4)

	code_CREATE         = byte(0xf0)
	code_CALL           = byte(0xf1)
	code_CALLCODE       = byte(0xf2)
	code_RETURN         = byte(0xf3)
	code_DELEGATECALL   = byte(0xf4)
	code_CALLBLACKBOX   = byte(0xf5)

	code_STATICCALL     = byte(0xfa)

	code_REVERT         = byte(0xfd)

	code_SUICIDE        = byte(0xff)
)
