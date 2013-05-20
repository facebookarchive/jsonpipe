// Package jsonpipe provides a io.Reader for Encoding JSON. This is useful to
// encode JSON without buffering.
package jsonpipe

import (
	"encoding/json"
	"io"
)

// Returns a Reader which will return the Encoded JSON. Make sure to read the
// data else the writer goroutine will hang around.
func Encode(v interface{}) io.Reader {
	reader, writer := io.Pipe()
	go func() {
		err := json.NewEncoder(writer).Encode(v)
		if err != nil {
			writer.CloseWithError(err)
		} else {
			writer.Close()
		}
	}()
	return reader
}
