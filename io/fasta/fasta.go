package fasta

// FASTA is a Seq implementation for FASTA format sequences
type FASTA struct {
	Description string
	Sequence    string
}

// ID returns the description of the sequnece.
func (f FASTA) ID() string {
	return f.Description
}

// String returns the textual representation of the sequence.
func (f FASTA) String() string {
	return f.Sequence
}

// Length returns the length of the sequence
func (f FASTA) Length() int {
	return len(f.Sequence)
}
