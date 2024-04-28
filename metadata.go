package clerk

import "github.com/ValGoldun/logger"

type Metadata map[string]string

func (m Metadata) LoggerFields() logger.Fields {
	var fields = make(logger.Fields, len(m))

	var i int
	for key, value := range m {
		fields[i] = logger.Field{
			Key:   key,
			Value: value,
		}

		i++
	}

	return fields
}
