package yahc

import (
	"errors"
	"fmt"
	"io"
)

// A ResponseParser should convert a response body into an object
type ResponseParser interface {
	// Should convert the r into v
	// Closing the reader is handled by yahc
	Convert(r io.Reader, v interface{}) error
}

var responseParsers = make(map[string]ResponseParser)

// Registers a new parser for the given content type
// If a parser is already registered for the content type, it
// is overwritten
func RegisterResponseParser(contentType string, parser ResponseParser) {
	responseParsers[contentType] = parser
}

// Parses the response body into the v object
// v should be a pointer
func ParseResponse(contentType string, body io.Reader, v interface{}) error {
	parser, exists := responseParsers[contentType]
	if !exists {
		return errors.New(fmt.Sprintf("no request parser found for content type '%s'", contentType))
	}

	return parser.Convert(body, v)
}
