package problem

type Metadata map[string]string

type Problem struct {
	text       string
	isCritical bool
	metadata   Metadata
}

func New(text string, fields ...Field) error {
	return Problem{text: text, metadata: Fields(fields...).metadata()}
}

func NewCritical(text string, fields ...Field) error {
	return Problem{text: text, isCritical: true, metadata: Fields(fields...).metadata()}
}

func (e Problem) Error() string {
	return e.text
}

func (e Problem) IsCritical() bool {
	return e.isCritical
}
