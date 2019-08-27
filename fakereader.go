package fakereader

import "io"

// FakeReader represents a _Fake_ reader to be used as a controlled stub for
// Go tests.
type FakeReader struct {
	q      *int
	start  int
	length int
}

// NewFakeReader returns a _Fake_ reader to be used as a controlled stub for
// Go tests.
// It will return a reader that will, in turn, fill your buffer with up to the
// configured length of bytes, those bytes will cycle from 0 to 255 in loop, starting at the given number, up
// to the said length.  When the length is reached, the reader will return
// an io.EOF error along with the remaining bytes, as expected of a reader.
func NewFakeReader(start, length int) (r *FakeReader) {
	return &FakeReader{
		q:      &start,
		start:  start,
		length: start + length,
	}
}

func (r FakeReader) Read(p []byte) (n int, err error) {
	var ctr int
	for ctr = 0; ctr < (len(p)) && *r.q < r.length; ctr++ {
		p[ctr] = byte(*r.q)
		*r.q++
	}
	if ctr < len(p) {
		return ctr, io.EOF
	}
	return ctr, nil
}

// Reset will reset the reader byte stream to it's begining, which mean that
// when Read() will next be called, the first byte returned will have restarted
// at the start number initially chosen, than start+1, start+2, start+3, etc.
func (r *FakeReader) Reset() {
	*r.q = r.start
}
