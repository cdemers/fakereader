package fakereader

import (
	"io"
	"testing"
)

func TestLessThan256Range(t *testing.T) {
	bufSize := 5
	readerSize := 10

	buf := make([]byte, bufSize)
	r := NewFakeReader(0, readerSize)
	n, err := r.Read(buf)

	if err != nil {
		t.Errorf("[TestLessThan256Range] Read(buf) should not return an error, got %s", err)
	}

	if n != bufSize {
		t.Errorf("[TestLessThan256Range] expected n:%d, got %d", bufSize, n)
	}

	// Check if all returned bytes match what is expected (a range from 0 to 4)
	ctr := byte(0)
	for _, v := range buf {
		if v != ctr {
			t.Errorf("[TestLessThan256Range] returned bytes expected to be %d at position %d, got %d", ctr, ctr, v)
		}
		ctr++
	}

	// Check that we get io.EOF if we read more
	_, err = r.Read(buf) // Read the remaining 5 bytes

	if err != nil {
		t.Errorf("[TestLessThan256Range] Read(buf) should not return an error, got %s", err)
	}

	n, err = r.Read(buf) // Read one too many

	if err == nil {
		t.Errorf("[TestLessThan256Range] Read(buf) should return an io.EOF error, got nil and n:%d", n)
	}

}

func TestMoreThan256Range(t *testing.T) {
	bufSize := 270
	readerSize := 300

	buf := make([]byte, bufSize)
	r := NewFakeReader(0, readerSize)
	n, err := r.Read(buf)

	if err != nil {
		t.Errorf("[TestMoreThan256Range] Read(buf) should not return an error, got %s", err)
	}

	if n != bufSize {
		t.Errorf("[TestMoreThan256Range] expected n:%d, got %d", bufSize, n)
	}

	// Check if all returned bytes match what is expected (a range from 0 to 269 (overflow 8 bits))
	ctr := byte(0)
	for _, v := range buf {
		if v != ctr {
			t.Errorf("[TestMoreThan256Range] returned bytes expected to be %d at position %d, got %d", ctr, ctr, v)
		}
		ctr++
	}

}

func TestNewFakeReaderOverByteRange(t *testing.T) {
	buf := make([]byte, 257)
	r := NewFakeReader(0, 257)
	n, err := r.Read(buf)

	if err != nil {
		t.Errorf("did not expect error, got %s", err)
	}

	if n != 257 {
		t.Errorf("expected 257, got %d", n)
	}

	ctr := byte(0)
	for _, v := range buf {
		if v != ctr {
			t.Errorf("expected %d, got %d", ctr, v)
		}
		ctr++
	}
}

func TestNewFakeReaderEOF(t *testing.T) {
	buf := make([]byte, 5)
	r := NewFakeReader(0, 7)

	n, err := r.Read(buf)

	if err != nil {
		t.Errorf("did not expect error, got %s", err)
	}

	if n != 5 {
		t.Errorf("expected %d, got %d", len(buf), n)
	}

	n, err = r.Read(buf)

	if err != io.EOF {
		t.Errorf("expected io.EOF, got %v", err)
	}

	if n != 2 {
		t.Errorf("expected 2, got %d", n)
	}
}

func TestFakeReaderReset(t *testing.T) {

	bufSize := 2
	readerSize := 10

	buf := make([]byte, bufSize)
	r := NewFakeReader(10, readerSize)

	_, err := r.Read(buf)
	if err != nil {
		t.Errorf("[TestFakeReaderReset] Read(buf) should not return an error, got %s", err)
	}

	if buf[0] != 10 {
		t.Errorf("Expected buf[0] to be 10, got %d", buf[0])
	}

	r.Reset()
	_, err = r.Read(buf)
	if err != nil {
		t.Errorf("[TestFakeReaderReset] Read(buf) should not return an error, got %s", err)
	}

	if buf[0] != 10 {
		t.Errorf("Expected buf[0] to be 10, got %d", buf[0])
	}

}
