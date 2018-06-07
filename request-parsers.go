package yahc

import (
	"errors"
	"fmt"
	"strings"
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
	requestParsers[strings.ToLower(contentType)] = parser
}

// Convert the body into bytes
// Will attempt to convert into the requested content type
// If no matching parser is registered, will return an error
func ConvertRequestBody(contentType string, body interface{}) (data []byte, err error) {
	contentType = strings.ToLower(contentType)

	// Try to get a registered parser
	// First attempt a full match
	parser, exists := requestParsers[contentType]
	if !exists {
		// next attempt a match on the content type only
		rawType := strings.Split(contentType, ";")[0]
		parser, exists = requestParsers[rawType]
		if !exists {
			// If that didn't work either, then fail full out
			return nil, errors.New(fmt.Sprintf("no request parser found for content type '%s'", contentType))
		}
	}

	// Actually convert the data
	data, err = parser.Convert(body)

	return
}
