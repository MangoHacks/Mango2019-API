package web

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTPError types
var (
	// 4xx Client Errors
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

func (e *HTTPError) getStatusCode() int {
	return e.statusCode
}

// Error returns a string representation of the error.
func (e *HTTPError) Error() string {
	return strconv.Itoa(e.statusCode) + " " + e.err
}

// JSONResponse represents a JSON response to send back to the client.
type JSONResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SendHTTPResponse sends back a response to the client.
func SendHTTPResponse(w http.ResponseWriter, v interface{}) error {
	var rsp JSONResponse
	switch v.(type) {
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
			w.WriteHeader(e.getStatusCode())
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
			return nil
		}
	default:
		w.WriteHeader(InternalServerError.getStatusCode())
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
