package yahc

import (
	"errors"
	"fmt"
)

// A request parser will take to the request body struct, and turn it into
// a byte array that can be passed along to the actual request
type RequestParser interface {
	// Should convert the given argument into a byte array
	Convert(v interface{}) ([]byte, error)
}

var requestParsers = make(map[string]RequestParser)

// Registers a new parser for the given contentType
// If a parser is already registered, it's overwritten
func RegisterRequestParser(contentType string, parser RequestParser) {
	requestParsers[contentType] = parser
}

// Convert the body into bytes
// Will attempt to convert into the requested content type
// If no matching parser is registered, will return an error
func ConvertRequestBody(contentType string, body interface{}) (data []byte, err error) {
	// Try to get a registered parser
	parser, exists := requestParsers[contentType]
	if !exists {
		return nil, errors.New(fmt.Sprintf("no request parser found for content type '%s'", contentType))
	}

	// Actually convert the data
	data, err = parser.Convert(body)

	return
}
