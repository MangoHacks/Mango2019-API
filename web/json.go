package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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
