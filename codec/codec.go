package codec

import (
	"encoding/json"
)

// DefaultCodec is the default codec used by booby
var DefaultCodec Codec = &JSONCodec{}

// Codec is the interface that wraps the booby Message data encoding method.
//
// # Marshal returns the JSON encoding of v
//
// Unmarshal parses the Message data and stores the result
// in the value pointed to by v
type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// JSONCodec wraps std json
type JSONCodec struct{}

// Marshal wraps std json.Marshal
func (j *JSONCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal wraps std json.Unmarshal
func (j *JSONCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// SetCodec sets default codec instance
func SetCodec(c Codec) {
	DefaultCodec = c
}
