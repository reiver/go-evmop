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

	CodeBlockhash      = byte(0x40)
	CodeCoinbase       = byte(0x41)
	CodeTimestamp      = byte(0x42)
	CodeNumber         = byte(0x43)
	CodeDifficulty     = byte(0x44)
	CodeGasLimit       = byte(0x45)

	CodePop            = byte(0x50)
	CodeMLoad          = byte(0x51)
	CodeMStore         = byte(0x52)
	CodeMStore8        = byte(0x53)
	CodeSLoad          = byte(0x54)
	CodeSStore         = byte(0x55)
	CodeJump           = byte(0x56)
	CodeJumpI          = byte(0x57)
	CodePC             = byte(0x58)
	CodeMSize          = byte(0x59)
	CodeGas            = byte(0x5a)
	CodeJumpDest       = byte(0x5b)

	CodePush1          = byte(0x5f +  1)
	CodePush2          = byte(0x5f +  2)
	CodePush3          = byte(0x5f +  3)
	CodePush4          = byte(0x5f +  4)
	CodePush5          = byte(0x5f +  5)
	CodePush6          = byte(0x5f +  6)
	CodePush7          = byte(0x5f +  7)
	CodePush8          = byte(0x5f +  8)
	CodePush9          = byte(0x5f +  9)
	CodePush10         = byte(0x5f + 10)
	CodePush11         = byte(0x5f + 11)
	CodePush12         = byte(0x5f + 12)
	CodePush13         = byte(0x5f + 13)
	CodePush14         = byte(0x5f + 14)
	CodePush15         = byte(0x5f + 15)
	CodePush16         = byte(0x5f + 16)
	CodePush17         = byte(0x5f + 17)
	CodePush18         = byte(0x5f + 18)
	CodePush19         = byte(0x5f + 19)
	CodePush20         = byte(0x5f + 20)
	CodePush21         = byte(0x5f + 21)
	CodePush22         = byte(0x5f + 22)
	CodePush23         = byte(0x5f + 23)
	CodePush24         = byte(0x5f + 24)
	CodePush25         = byte(0x5f + 25)
	CodePush26         = byte(0x5f + 26)
	CodePush27         = byte(0x5f + 27)
	CodePush28         = byte(0x5f + 28)
	CodePush29         = byte(0x5f + 29)
	CodePush30         = byte(0x5f + 30)
	CodePush31         = byte(0x5f + 31)
	CodePush32         = byte(0x5f + 32)

	CodeDup1           = byte(0x7f +  1)
	CodeDup2           = byte(0x7f +  2)
	CodeDup3           = byte(0x7f +  3)
	CodeDup4           = byte(0x7f +  4)
	CodeDup5           = byte(0x7f +  5)
	CodeDup6           = byte(0x7f +  6)
	CodeDup7           = byte(0x7f +  7)
	CodeDup8           = byte(0x7f +  8)
	CodeDup9           = byte(0x7f +  9)
	CodeDup10          = byte(0x7f + 10)
	CodeDup11          = byte(0x7f + 11)
	CodeDup12          = byte(0x7f + 12)
	CodeDup13          = byte(0x7f + 13)
	CodeDup14          = byte(0x7f + 14)
	CodeDup15          = byte(0x7f + 15)
	CodeDup16          = byte(0x7f + 16)

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
