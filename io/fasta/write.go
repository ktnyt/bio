package fasta

import (
	"bufio"
	"io"

	"github.com/ktnyt/bio"
)

func wrapString(s string, l int) string {
	t := ""
	for len(s) > l {
		t += s[:l] + "\n"
		s = s[l:]
	}
	return t + s
}

// Writer implements writing sequences to an io.Writer object in FASTA format.
type Writer struct {
	writer *bufio.Writer
	length int
}

// NewWriter returns a new Writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		writer: bufio.NewWriter(w),
		length: 80,
	}
}

// NewWriterWrap returns a new Writer with wrapping width.
func NewWriterWrap(w io.Writer, l int) *Writer {
	return &Writer{
		writer: bufio.NewWriter(w),
		length: l,
	}
}

// Write a slice of sequences to the buffer.
func (w *Writer) Write(ss []bio.Seq) (err error) {
	for _, s := range ss {
		_, err = w.writer.Write([]byte(">" + s.ID() + "\n"))
		if err != nil {
			return err
		}
		_, err = w.writer.Write([]byte(wrapString(s.Seq(), w.length) + "\n"))
		if err != nil {
			return err
		}
	}

	return w.writer.Flush()
}

// WriteSeq writes a single sequence to the buffer.
func (w *Writer) WriteSeq(s bio.Seq) error {
	return w.Write([]bio.Seq{s})
}
