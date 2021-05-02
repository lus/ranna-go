package ranna

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

// Error represents a ranna API error
type Error struct {
	Message string `json:"error"`
	Code    int    `json:"code"`
	Context string `json:"context"`
}

// Error stringifies a ranna API error
func (err Error) Error() string {
	return fmt.Sprintf("{ MSG: %s ;; CODE: %d ;; CTX: %s }", err.Message, err.Code, err.Context)
}

func parseErrorFromResponse(response *fasthttp.Response) error {
	status := response.StatusCode()

	if status >= 200 && status < 300 {
		return nil
	}

	body := response.Body()

	rannaError := new(Error)
	if err := json.Unmarshal(body, rannaError); err != nil {
		return errors.New(string(body))
	}
	return *rannaError
}
