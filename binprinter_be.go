package binp

// Output 2 bigendian bytes.
func (p *Printer) B16(d uint16) *Printer {
	p.w = append(p.w, byte(d>>8), byte(d))
	return p
}

// Output 4 bigendian bytes.
func (p *Printer) B32(d uint32) *Printer {
	p.w = append(p.w, byte(d>>24), byte(d>>16), byte(d>>8), byte(d))
	return p
}

// Output 4 bigendian bytes.
func (p *Printer) B64(d uint64) *Printer {
	p.w = append(p.w, byte(d>>56), byte(d>>48), byte(d>>40), byte(d>>32), byte(d>>24), byte(d>>16), byte(d>>8), byte(d))
	return p
}

// Output a string with a 4 byte bigendian length prefix and no trailing null.
func (p *Printer) B32String(d string) *Printer {
	return p.B32(uint32(len(d))).String(d)
}

// Output bytes with a 4 byte bigendian length prefix and no trailing null.
func (p *Printer) B32Bytes(d []byte) *Printer {
	return p.B32(uint32(len(d))).Bytes(d)
}

// Output a string with a 2 byte bigendian length prefix and no trailing null.
func (p *Printer) B16String(d string) *Printer {
	if len(d) > 0xffff {
		panic("binprinter: string too long")
	}
	return p.B16(uint16(len(d))).String(d)
}

// Output a string with a 1 byte bigendian length prefix and no trailing null.
func (p *Printer) B8String(d string) *Printer {
	if len(d) > 0xff {
		panic("binprinter: string too long")
	}
	return p.Byte(byte(len(d))).String(d)
}

// Add a 16 bit field at the current location that will be filled with the length.
func (p *Printer) LenB16(l *Len) *Printer {
	l.ls = append(l.ls, ls{uint32(len(p.w)), 2 | lenMaskBE})
	return p.N16(0)
}

// Add a 32 bit field at the current location that will be filled with the length.
func (p *Printer) LenB32(l *Len) *Printer {
	l.ls = append(l.ls, ls{uint32(len(p.w)), 4 | lenMaskBE})
	return p.N32(0)
}
