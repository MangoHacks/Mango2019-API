package routes

import "net/http"

// sendJSONResponse will send a JSON response back to the caller with the appropriate body and status code.
func sendJSONResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.WriteHeader(statusCode)
	w.Write(body)
}
