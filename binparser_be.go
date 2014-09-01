package binp

import (
	"unsafe"
)

// Parser type. Don't touch the internals.
type Parser struct {
	R   []byte
	Off int
}

func ntohl(uint32) uint32
func ntohs(uint16) uint16


// Create a new parser with the given buffer.
func NewParser(b []byte) *Parser {
	return &Parser{b, 0}
}

// Parse a byte from the buffer.
func (p *Parser) Byte(d *byte) *Parser {
	*d = p.R[p.Off]
	p.Off++
	return p
}

// Parse 4 bigendian bytes from the buffer.
func (p *Parser) U32(d *uint32) *Parser {
	*d = ntohl(*(*uint32)((unsafe.Pointer(&p.R[p.Off]))))
	p.Off += 4
	return p
}
// Parse 8 bigendian bytes from the buffer.
func (p *Parser) U64(d *uint64) *Parser {
	*d = ntohq(*(*uint64)((unsafe.Pointer(&p.R[p.Off]))))
	p.Off += 8
	return p
}

// Parse 2 bigendian bytes from the buffer.
func (p *Parser) U16(d *uint16) *Parser {
	*d = ntohs(*(*uint16)((unsafe.Pointer(&p.R[p.Off]))))
	p.Off += 2
	return p
}

// Parse n bytes from the buffer to a []byte pointer.
func (p *Parser) NBytes(n int, d *[]byte) *Parser {
	if n > len(p.R[p.Off:]) {
		panic("binparser: overflowing length")
	}
	*d = make([]byte, n)
	copy(*d, p.R[p.Off:])
	p.Off += n
	return p
}

// Parse a string with a 4 byte bigendian length prefix to a []byte pointer.
func (p *Parser) U32Bytes(d *[]byte) *Parser {
	var v uint32
	return p.U32(&v).NBytes(int(v), d)
}

// Parse n bytes from the buffer to a string pointer.
func (p *Parser) NString(n int, d *string) *Parser {
	if n > len(p.R[p.Off:]) {
		panic("binparser: overflowing length")
	}
	b := make([]byte, n)
	copy(b, p.R[p.Off:])
	p.Off += n
	*(*string)(d) = *(*string)(unsafe.Pointer(&b))
	return p
}

// Parse a string with a 4 byte bigendian length prefix to a string pointer.
func (p *Parser) U32String(d *string) *Parser {
	var v uint32
	return p.U32(&v).NString(int(v), d)
}

// Parse a string with a 2 byte bigendian length prefix to a string pointer.
func (p *Parser) U16String(d *string) *Parser {
	var v uint16
	return p.U16(&v).NString(int(v), d)
}

// Parse a string with a 1 byte length prefix to a string pointer.
func (p *Parser) U8String(d *string) *Parser {
	var v uint8
	return p.Byte(&v).NString(int(v), d)
}

// Parse a null terminated string.
func (p *Parser) String0(d *string) *Parser {
	for i,ch := range p.R[p.Off:] {
		if ch == 0 {
			p.NString(i, d)
			p.Off++
			return p
		}
	}
	panic("String0: null byte not found")
}

// Check that we are at the end of input.
func (p *Parser) End() {
	if p.Off != len(p.R) {
		panic("binparser: overlong packet")
	}
}

// Check that we are at the end of input.
func (p *Parser) AtEnd() bool {
	return p.Off == len(p.R)
}

// Peek the rest of input as raw bytes.
func (p *Parser) PeekRest(d *[]byte) *Parser {
	*d = p.R[p.Off:]
	return p
}
