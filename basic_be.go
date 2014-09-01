package binp

import "unsafe"

func ntohq(x uint64) uint64 {
	return ((uint64(ntohl(uint32((x) >> 32)))) | (uint64(ntohl(uint32(x))) << 32))
}

// Take the first 8 bigendian bytes of a buffer and return them.
func B64(b []byte) uint64 {
	if len(b) < 8 {
		panic("B64: too short buffer")
	}
	return ntohq(*(*uint64)(unsafe.Pointer(&b[0])))
}

// Poke 8 bigendian bytes into the buffer.
func PutB64(b []byte, v uint64) {
	if len(b) < 8 {
		panic("PutB64: too short buffer")
	}
	*(*uint64)(unsafe.Pointer(&b[0])) = ntohq(v)
}

// Take the first 4 bigendian bytes of a buffer and return them.
func B32(b []byte) uint32 {
	if len(b) < 4 {
		panic("B32: too short buffer")
	}
	return ntohl(*(*uint32)(unsafe.Pointer(&b[0])))
}

// Poke 4 bigendian bytes into the buffer.
func PutB32(b []byte, v uint32) {
	if len(b) < 4 {
		panic("PutB32: too short buffer")
	}
	*(*uint32)(unsafe.Pointer(&b[0])) = ntohl(v)
}

// Take the first 2 bigendian bytes of a buffer and return them.
func B16(b []byte) uint16 {
	if len(b) < 2 {
		panic("B16: too short buffer")
	}
	return ntohs(*(*uint16)(unsafe.Pointer(&b[0])))
}

// Poke 2 bigendian bytes into the buffer.
func PutB16(b []byte, v uint16) {
	if len(b) < 2 {
		panic("PutB16: too short buffer")
	}
	*(*uint16)(unsafe.Pointer(&b[0])) = ntohs(v)
}
