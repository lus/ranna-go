package snippets

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

// Error represents a ranna snippets API error
type Error struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Status  int    `json:"status"`
	TraceID string `json:"traceId"`
}

// Error stringifies a ranna API error
func (err Error) Error() string {
	return fmt.Sprintf("{ TYPE: %s ;; TITLE: %s ;; STATUS: %d ;; TID: %s }", err.Type, err.Title, err.Status, err.TraceID)
}

func parseErrorFromResponse(response *fasthttp.Response) error {
	status := response.StatusCode()

	if status >= 200 && status < 300 {
		return nil
	}

	body := response.Body()

	snippetsError := new(Error)
	if err := json.Unmarshal(body, snippetsError); err != nil {
		return errors.New(string(body))
	}
	return *snippetsError
}
