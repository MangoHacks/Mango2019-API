package web

import (
	"encoding/json"
	"io"
)

// ReadJSONBodyIntoStruct will take the given io.ReadCloser of a JSON and read it into a matching, given struct.
//
// eg:
//  type preregisterRequest struct {
//  	Email string `json:"email"`
//  }
//  var prr preregisterRequest
//  if err := web.ReadJSONBodyIntoStruct(r.Body, &prr); err != nil {
//  	// TODO: Handle error
//  }
//  fmt.Println(prr.Email)
func ReadJSONBodyIntoStruct(body io.ReadCloser, v interface{}) error {
	if err := json.NewDecoder(body).Decode(&v); err != nil {
		return &BadRequestError
	}
	return nil
}
