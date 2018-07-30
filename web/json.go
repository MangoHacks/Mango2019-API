package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

// ReadJSONBodyIntoStruct will take the given io.ReadCloser of a JSON and read it into a matching, given struct.
func ReadJSONBodyIntoStruct(body io.ReadCloser, v interface{}) *HTTPError {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return &BadRequestError
	}
	defer body.Close()
	log.Printf("received the following payload: %s", string(b))
	if err := json.Unmarshal(b, &v); err != nil {
		return &BadRequestError
	}
	return nil
}
