package fasta

import (
	"bufio"
	"io"

	"github.com/ktnyt/bio"
)

// Writer implements writing sequences to an io.Writer object in FASTA format.
type Writer struct {
	writer *bufio.Writer
}

// NewWriter returns a new Writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		writer: bufio.NewWriter(w),
	}
}

// Write a slice of sequences to the buffer.
func (w *Writer) Write(ss []bio.Seq) (err error) {
	for _, s := range ss {
		_, err = w.writer.Write([]byte(">" + s.ID() + "\n"))
		if err != nil {
			return err
		}
		_, err = w.writer.Write([]byte(s.String() + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

// Write a single sequence to the buffer.
func (w *Writer) WriteSeq(s bio.Seq) error {
	return w.Write([]bio.Seq{s})
}
