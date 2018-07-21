// Package web contains many functions and structs that are common
// to applications and services that interact via the internet.
package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

// ReadJSONBodyIntoStruct will take the given JSON ReadCloser and read it into a matching, given struct.
func ReadJSONBodyIntoStruct(body io.ReadCloser, v interface{}) *HTTPError {
	defer body.Close()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return &BadRequestError
	}
	log.Printf("received the following payload: %s", string(b))
	if err := json.Unmarshal(b, &v); err != nil {
		return &BadRequestError
	}
	return nil
}
