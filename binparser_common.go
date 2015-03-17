package binp

import (
	"errors"
	"fmt"
)

// Parser type. Don't touch the internals.
type Parser struct {
	r   []byte
	off int
}

// Catch a panic into an error.
func Catch(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()
	f()
	return
}

// Catch a panic into an error.
func CatchErr(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()
	err = f()
	return
}

// Create a new parser with the given buffer. Panics on error.
func NewParser(b []byte) *Parser {
	return &Parser{b, 0}
}

// Parse a byte from the buffer.
func (p *Parser) Byte(d *byte) *Parser {
	*d = p.r[p.off]
	p.off++
	return p
}

// Parse a byte from the buffer, synonym for .Byte.
func (p *Parser) B8(d *byte) *Parser {
	*d = p.r[p.off]
	p.off++
	return p
}

// Parse a byte from the buffer, synonym for .Byte.
func (p *Parser) N8(d *byte) *Parser {
	*d = p.r[p.off]
	p.off++
	return p
}

// Parse n bytes from the buffer and copy to a []byte pointer that is allocated.
func (p *Parser) NBytes(n int, d *[]byte) *Parser {
	if n > len(p.r[p.off:]) {
		panic("binparser: overflowing length")
	}
	*d = make([]byte, n)
	copy(*d, p.r[p.off:])
	p.off += n
	return p
}

// Parse n bytes from the buffer and copy to the supplied []byte.
func (p *Parser) NBytesCopy(n int, d []byte) *Parser {
	if n > len(p.r[p.off:]) {
		panic("binparser: overflowing length")
	}
	copy(d, p.r[p.off:p.off+n])
	p.off += n
	return p
}

// Parse n bytes from the buffer to a []byte pointer that refers to the parser internal buffer.
func (p *Parser) NBytesPeek(n int, d *[]byte) *Parser {
	if n > len(p.r[p.off:]) {
		panic("binparser: overflowing length")
	}
	*d = p.r[p.off : p.off+n]
	p.off += n
	return p
}

// Ensure the input is aligned possibly skipping bytes.
func (p *Parser) Align(n int) *Parser {
	r := p.off % n
	if r == 0 {
		return p
	}
	p.off += n - r
	return p
}

// Skip bytes.
func (p *Parser) Skip(n int) *Parser {
	p.off += n
	return p
}

// Parse a null terminated string.
func (p *Parser) String0(d *string) *Parser {
	for i, ch := range p.r[p.off:] {
		if ch == 0 {
			p.NString(i, d)
			p.off++
			return p
		}
	}
	panic("String0: null byte not found")
}

// Check that we are at the end of input.
func (p *Parser) End() {
	if p.off != len(p.r) {
		panic("binparser: overlong packet")
	}
}

// Check that we are at the end of input.
func (p *Parser) AtEnd() bool {
	return p.off == len(p.r)
}

// Peek the rest of input as raw bytes.
func (p *Parser) PeekRest(d *[]byte) *Parser {
	*d = p.r[p.off:]
	return p
}
