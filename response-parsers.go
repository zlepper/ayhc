package yahc

import (
	"errors"
	"fmt"
	"io"
	"strings"
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
	contentType = strings.ToLower(contentType)

	// Try to get a registered parser
	// First attempt a full match
	parser, exists := responseParsers[contentType]
	if !exists {
		// next attempt a match on the content type only
		rawType := strings.Split(contentType, ";")[0]
		parser, exists = responseParsers[rawType]
		if !exists {
			// If that didn't work either, then fail full out
			return errors.New(fmt.Sprintf("no response parser found for content type '%s'", contentType))
		}
	}

	return parser.Convert(body, v)
}
