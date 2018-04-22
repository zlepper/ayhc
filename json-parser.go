package yahc

import (
	"encoding/json"
	"io"
)

func init() {
	RegisterRequestParser(ApplicationJsonContentType, JsonRequestParser{})
	RegisterResponseParser(ApplicationJsonContentType, JsonResponseParser{})
	RegisterResponseParser("application/json; charset=utf-8", JsonResponseParser{})
}

type JsonRequestParser struct {
}

func (JsonRequestParser) Convert(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

type JsonResponseParser struct {
}

func (JsonResponseParser) Convert(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
