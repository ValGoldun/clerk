package problem

type Field struct {
	Key   string
	Value string
}

type Fields []Field

func (f Fields) metadata() Metadata {
	var metadata = make(Metadata)

	for _, field := range f {
		metadata[field.Key] = field.Value
	}

	return metadata
}
