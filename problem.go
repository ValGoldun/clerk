package clerk

type Problem struct {
	Error    string   `json:"error"`
	Fields   []Field  `json:"fields,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type Field struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}
