TEXT ·ntohl(SB),7,$0
	MOVL arg+0(FP), AX
	// 0f c8                	bswap  %eax
	BYTE $0x0F
	BYTE $0xC8
	MOVL AX, ret+4(FP)
	RET
TEXT ·ntohs(SB),7,$0
	MOVL arg+0(FP), AX
	RORW $0x8, AX
	MOVL AX, ret+4(FP)
	RET

