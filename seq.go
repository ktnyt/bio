package bio

type Seq interface {
	ID() string
	Seq() string
	Length() int
}
