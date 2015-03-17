package binp

// Parse 4 native endian bytes from the buffer.
func (p *Parser) N32(d *uint32) *Parser {
	if p == nil || len(p.r) < p.off+4 {
		return nil
	}
	*d = NativeEndian.Uint32(p.r[p.off:])
	p.off += 4
	return p
}

// Parse 8 native endian bytes from the buffer.
func (p *Parser) N64(d *uint64) *Parser {
	if p == nil || len(p.r) < p.off+8 {
		return nil
	}
	*d = NativeEndian.Uint64(p.r[p.off:])
	p.off += 8
	return p
}

// Parse 2 native endian bytes from the buffer.
func (p *Parser) N16(d *uint16) *Parser {
	if p == nil || len(p.r) < p.off+2 {
		return nil
	}
	*d = NativeEndian.Uint16(p.r[p.off:])
	p.off += 2
	return p
}

// Parse a string with a 4 byte native endian length prefix to a []byte pointer.
func (p *Parser) N32Bytes(d *[]byte) *Parser {
	var v uint32
	return p.N32(&v).NBytes(int(v), d)
}

// Parse n bytes from the buffer to a string pointer.
func (p *Parser) NString(n int, d *string) *Parser {
	if p == nil || n > len(p.r[p.off:]) {
		return nil
	}
	bs := p.r[p.off : p.off+n]
	*d = string(bs)
	p.off += n
	return p
}

// Parse a string with a 4 byte native endian length prefix to a string pointer.
func (p *Parser) N32String(d *string) *Parser {
	var v uint32
	return p.N32(&v).NString(int(v), d)
}

// Parse a string with a 2 byte native endian length prefix to a string pointer.
func (p *Parser) N16String(d *string) *Parser {
	var v uint16
	return p.N16(&v).NString(int(v), d)
}

// Parse a string with a 1 byte length prefix to a string pointer.
func (p *Parser) N8String(d *string) *Parser {
	var v uint8
	return p.Byte(&v).NString(int(v), d)
}
