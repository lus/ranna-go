package ranna

import (
	"encoding/json"
	"time"

	"github.com/valyala/fasthttp"
)

// ExecutionRequest is being sent to ranna to execute a code snippet
type ExecutionRequest struct {
	Language    string            `json:"language"`
	Code        string            `json:"code"`
	Arguments   []string          `json:"arguments"`
	Environment map[string]string `json:"environment"`
}

// ExecutionResponse represents a response of a ranna code execution
type ExecutionResponse struct {
	StdOut string `json:"stdout"`
	StdErr string `json:"stderr"`

	// Duration represents the duration ranna took to execute the requested snippet
	// Please note that this is not 100% accurate as this is calculated on the client side
	Duration time.Duration `json:"-"`
}

// Execute executes a code snippet using ranna
func (client *Client) Execute(executionRequest *ExecutionRequest) (*ExecutionResponse, error) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.SetMethod(fasthttp.MethodPost)
	request.SetRequestURI(client.instance + EndpointExecute)

	jsonBytes, err := json.Marshal(executionRequest)
	if err != nil {
		return nil, err
	}
	request.Header.SetContentType("application/json")
	request.SetBody(jsonBytes)

	start := time.Now()
	if err := client.http.Do(request, response); err != nil {
		return nil, err
	}
	duration := time.Since(start)

	if err := parseErrorFromResponse(response); err != nil {
		return nil, err
	}

	executionResponse := new(ExecutionResponse)
	if err := json.Unmarshal(response.Body(), executionResponse); err != nil {
		return nil, err
	}
	executionResponse.Duration = duration
	return executionResponse, nil
}
