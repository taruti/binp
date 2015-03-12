/*
The binp package provides parsing and printing of binary values.
It is particularly suitable for parsing irregular data structures provided by e.g. kernel.

*/
package binp

import "encoding/binary"

var NativeEndian = func() binary.ByteOrder {
	return binary.LittleEndian
}()
