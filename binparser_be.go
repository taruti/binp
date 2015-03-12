package binp

import "encoding/binary"

func ntohl(arg uint32) uint32
func ntohs(arg uint16) uint16

// Parse 4 bigendian bytes from the buffer.
func (p *Parser) B32(d *uint32) *Parser {
	*d = binary.BigEndian.Uint32(p.r[p.off:])
	p.off += 4
	return p
}

// Parse 8 bigendian bytes from the buffer.
func (p *Parser) B64(d *uint64) *Parser {
	*d = binary.BigEndian.Uint64(p.r[p.off:])
	p.off += 8
	return p
}

// Parse 2 bigendian bytes from the buffer.
func (p *Parser) B16(d *uint16) *Parser {
	*d = binary.BigEndian.Uint16(p.r[p.off:])
	p.off += 2
	return p
}

// Parse a string with a 4 byte bigendian length prefix to a []byte pointer.
func (p *Parser) B32Bytes(d *[]byte) *Parser {
	var v uint32
	return p.B32(&v).NBytes(int(v), d)
}

// Parse a string with a 4 byte bigendian length prefix to a string pointer.
func (p *Parser) B32String(d *string) *Parser {
	var v uint32
	return p.B32(&v).NString(int(v), d)
}

// Parse a string with a 2 byte bigendian length prefix to a string pointer.
func (p *Parser) B16String(d *string) *Parser {
	var v uint16
	return p.B16(&v).NString(int(v), d)
}

// Parse a string with a 1 byte length prefix to a string pointer.
func (p *Parser) B8String(d *string) *Parser {
	var v uint8
	return p.Byte(&v).NString(int(v), d)
}

func ntohq(x uint64) uint64 {
	return ((uint64(ntohl(uint32((x) >> 32)))) | (uint64(ntohl(uint32(x))) << 32))
}
