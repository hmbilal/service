package http

import (
	"fmt"
	"io"
)

type RequestError struct {
	StatusCode int
	Body       io.ReadCloser
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("request error code: %d", r.StatusCode)
}
