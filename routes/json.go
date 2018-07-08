package routes

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// sendJSONResponse will send a JSON response back to the caller with the appropriate body and status code.
func sendJSONResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.WriteHeader(statusCode)
	w.Write(body)
}

// readJSONBodyIntoStruct will read the body of a JSON payload and unmarshal it into a struct.
func readJSONBodyIntoStruct(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return errors.New("Invalid JSON")
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return errors.New("Invalid JSON")
	}
	return nil
}
