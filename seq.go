package bio

type Seq interface {
	ID() string
	String() string
	Length() int
}
