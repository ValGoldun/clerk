package problem

type Problem struct {
	text       string
	isCritical bool
	metadata   Metadata
}

func New(text string, fields ...Field) error {
	return Problem{text: text, metadata: Fields(fields).metadata()}
}

func NewCritical(text string, fields ...Field) error {
	return Problem{text: text, isCritical: true, metadata: Fields(fields).metadata()}
}

func (p Problem) Error() string {
	return p.text
}

func (p Problem) IsCritical() bool {
	return p.isCritical
}

func (p Problem) Metadata() Metadata {
	return p.metadata
}
