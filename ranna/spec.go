package ranna

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// Specs represents a map of specs
type Specs map[string]Spec

// Spec represents a language execution spec
type Spec struct {
	Use        string `json:"use"`
	Image      string `json:"image"`
	Entrypoint string `json:"entrypoint"`
	Filename   string `json:"filename"`
}

// Specs requests the specified specs from the ranna instance
func (client *Client) Specs() (Specs, error) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI(client.instance + EndpointSpec)

	if err := client.http.Do(request, response); err != nil {
		return nil, err
	}

	if err := parseErrorFromResponse(response); err != nil {
		return nil, err
	}

	specs := new(Specs)
	if err := json.Unmarshal(response.Body(), specs); err != nil {
		return nil, err
	}
	return *specs, nil
}
