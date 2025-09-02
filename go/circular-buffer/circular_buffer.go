package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

type Buffer struct {
	size  int
	bytes []byte
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		size:  size,
		bytes: make([]byte, 0, size),
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.bytes) == 0 {
		return 0, errors.New("buffer is empty")
	}
	c := b.bytes[0]
	b.bytes = b.bytes[1:]
	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.bytes) < b.size {
		b.bytes = append(b.bytes, c)
		return nil
	}
	return errors.New("buffer is full")
}

func (b *Buffer) Overwrite(c byte) {
	if len(b.bytes) < b.size {
		b.bytes = append(b.bytes, c)
	} else {
		b.bytes = append(b.bytes[1:], c)
	}
}

func (b *Buffer) Reset() {
	b.bytes = []byte{}
}
