// Package web contains many functions and structs that are common
// to applications and services that interact via the internet.
package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPErrors
//
// These which arise from HTTP requests and responses.
// It's useful to provide the user with these errors so they can know
// when they've messed up (or when you have).
var (
	// 4xx Client Errors
	// These are for when the client does something
	// incorrect.
	BadRequestError = HTTPError{
		err:        "Bad Request",
		statusCode: 400,
	}
	ForbiddenError = HTTPError{
		err:        "Forbidden",
		statusCode: 403,
	}
	NotFoundError = HTTPError{
		err:        "Not Found",
		statusCode: 404,
	}
	MethodNotAllowedError = HTTPError{
		err:        "Method Not Allowed",
		statusCode: 405,
	}

	// 5xx Server Errors
	// These are for when the server does something
	// incorrect.
	InternalServerError = HTTPError{
		err:        "Internal Server Error",
		statusCode: 500,
	}
)

// HTTPError represents an HTTP response error code.
type HTTPError struct {
	err        string
	statusCode int
}

// Error returns a string representation of the error.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d %s", e.statusCode, e.err)
}

// JSONResponse represents a JSON response to send back to the client.
type JSONResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SendHTTPResponse sends back a response to the client.
// It exhibits some polymorphism, depending on the type of interface{}
// provided, it sends an appropriate response to the user.
//
// The response it sends to the user follows a specific JSON pattern:
//  {
//  	"success": true/false,
//		"message": "message"
//  }
func SendHTTPResponse(w http.ResponseWriter, v interface{}) error {
	var rsp JSONResponse
	switch v.(type) {
	// case string provides shorthand for sending an OK response with
	// a custom message.
	case string:
		if s, ok := v.(string); ok {
			w.WriteHeader(200)
			rsp = JSONResponse{
				Success: true,
				Message: s,
			}
		}
	case HTTPError:
		if e, ok := v.(HTTPError); ok {
			w.WriteHeader(e.statusCode)
			rsp = JSONResponse{
				Success: false,
				Message: e.Error(),
			}
		}
	case error:
		if e, ok := v.(error); ok {
			w.WriteHeader(400)
			rsp = JSONResponse{
				Success: false,
				Message: e.Error(),
			}
		}
	case []byte:
		if j, ok := v.([]byte); ok {
			w.WriteHeader(200)
			w.Write(j)
		}
		return nil
	default:
		w.WriteHeader(InternalServerError.statusCode)
		rsp = JSONResponse{
			Success: false,
			Message: InternalServerError.Error(),
		}
	}
	b, err := json.Marshal(rsp)
	if err != nil {
		SendHTTPResponse(w, InternalServerError)
		return err
	}
	w.Write(b)
	return nil
}
