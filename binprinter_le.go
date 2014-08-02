// +build 386 amd64 arm

package native

// Printer type. Don't touch the internals.
type Printer struct {
	w []byte
}

// Create a new printer with empty output.
func Out() *Printer {
	return &Printer{[]byte{}}
}

// Create a new printer with output prefixed with the given byte slice.
func Outwith(b []byte) *Printer {
	return &Printer{b}
}

// Create a new printer with an empty slice with the capacity given below.
func OutCap(initialcap int) *Printer {
	return &Printer{make([]byte, 0, initialcap)}
}

// Output a byte.
func (p *Printer) Byte(d byte) *Printer {
	p.w = append(p.w, d)
	return p
}

// Output 2 native endian bytes.
func (p *Printer) U16(d uint16) *Printer {
	p.w = append(p.w, byte(d), byte(d>>8))
	return p
}

// Output 4 native endian bytes.
func (p *Printer) U32(d uint32) *Printer {
	p.w = append(p.w, byte(d), byte(d>>8), byte(d>>16), byte(d>>24))
	return p
}

// Output 4 native endian bytes.
func (p *Printer) U64(d uint64) *Printer {
	p.w = append(p.w, byte(d), byte(d>>8), byte(d>>16), byte(d>>24), byte(d>>32), byte(d>>40), byte(d>>48), byte(d>>56))
	return p
}

var z16 = make([]byte, 16)

// Align to boundary
func (p *Printer) Align(n int) *Printer {
	r := len(p.w) % n
	if r == 0 {
		return p
	}
	r = n - r
	for r > 0 {
		cur := r
		if cur > 16 {
			cur = 16
		}
		p.w = append(p.w, z16[:cur]...)
		r -= cur
	}

	return p
}

// Skip (zero-fill) some bytes.
func (p *Printer) Skip(n int) *Printer {
	for n > 0 {
		cur := n
		if cur > 16 {
			cur = 16
		}
		p.w = append(p.w, z16[:cur]...)
		n -= cur
	}

	return p
}

// Output a raw byte slice with no length prefix.
func (p *Printer) Bytes(d []byte) *Printer {
	p.w = append(p.w, d...)
	return p
}

// Output a raw string with no length prefix.
func (p *Printer) String(d string) *Printer {
	p.w = append(p.w, []byte(d)...)
	return p
}

// Output a string with a 4 byte native endian length prefix and no trailing null.
func (p *Printer) U32String(d string) *Printer {
	return p.U32(uint32(len(d))).String(d)
}

// Output bytes with a 4 byte native endian length prefix and no trailing null.
func (p *Printer) U32Bytes(d []byte) *Printer {
	return p.U32(uint32(len(d))).Bytes(d)
}

// Output a string with a 2 byte native endian length prefix and no trailing null.
func (p *Printer) U16String(d string) *Printer {
	if len(d) > 0xffff {
		panic("binprinter: string too long")
	}
	return p.U16(uint16(len(d))).String(d)
}

// Output a string with a 1 byte native endian length prefix and no trailing null.
func (p *Printer) U8String(d string) *Printer {
	if len(d) > 0xff {
		panic("binprinter: string too long")
	}
	return p.Byte(byte(len(d))).String(d)
}

// Output a string terminated by a null-byte
func (p *Printer) String0(d string) *Printer {
	return p.String(d).Byte(0)
}

// Get the output as a byte slice.
func (p *Printer) Out() []byte {
	return p.w
}

// Start counting bytes for the length field in question.
func (p *Printer) LenStart(l *Len) *Printer {
	l.start = len(p.w)
	return p
}

// Add a 16 bit field at the current location that will be filled with the length.
func (p *Printer) LenU16(l *Len) *Printer {
	l.ls = append(l.ls, ls{uint32(len(p.w)), 2})
	return p.U16(0)
}

// Call LenDone for all the arguments
func (p *Printer) LensDone(ls ...*Len) *Printer {
	for _, l := range ls {
		p.LenDone(l)
	}
	return p
}

// Fill fields associated with this length with the current offset.
func (p *Printer) LenDone(l *Len) *Printer {
	plen := len(p.w) - l.start
	for _, ls := range l.ls {
		switch ls.size {
		case 2:
			PutU16(p.w[ls.offset:], uint16(plen))
		}
	}
	return p
}

// Type for handling length fields.
type Len struct {
	ls    []ls
	start int
}

type ls struct {
	offset uint32
	size   uint32
}
