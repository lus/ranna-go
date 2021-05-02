package snippets

import (
	"encoding/json"
	"time"

	"github.com/valyala/fasthttp"
)

// Client represents a ranna snippets client
type Client struct {
	instance string
	http     *fasthttp.Client
}

// NewClient creates a new ranna snippets client
func NewClient(instance string) *Client {
	return &Client{
		instance: instance,
		http: &fasthttp.Client{
			Name: "lus/ranna-go",
		},
	}
}

// Snippet represents a ranna code snippet
type Snippet struct {
	Ident        string `json:"ident"`
	Language     string `json:"language"`
	Code         string `json:"code"`
	ID           string `json:"id"`
	RawTimestamp string `json:"timestamp"`

	Timestamp time.Time `json:"-"`
}

// Snippet requests a code snippet from the ranna snippet service
func (client *Client) Snippet(ident string) (*Snippet, error) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI(client.instance + EndpointSnippets + "/" + ident)

	if err := client.http.Do(request, response); err != nil {
		return nil, err
	}

	if err := parseErrorFromResponse(response); err != nil {
		return nil, err
	}

	snippet := new(Snippet)
	if err := json.Unmarshal(response.Body(), snippet); err != nil {
		return nil, err
	}
	snippet.Timestamp = parseTimestamp(snippet.RawTimestamp)
	return snippet, nil
}

// Create creates a new ranna code snippet
func (client *Client) Create(snippet *Snippet) (*Snippet, error) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.SetMethod(fasthttp.MethodPost)
	request.SetRequestURI(client.instance + EndpointSnippets)

	jsonBytes, err := json.Marshal(snippet)
	if err != nil {
		return nil, err
	}
	request.Header.SetContentType("application/json")
	request.SetBody(jsonBytes)

	if err := client.http.Do(request, response); err != nil {
		return nil, err
	}

	if err := parseErrorFromResponse(response); err != nil {
		return nil, err
	}

	createdSnippet := new(Snippet)
	if err := json.Unmarshal(response.Body(), createdSnippet); err != nil {
		return nil, err
	}
	createdSnippet.Timestamp = parseTimestamp(createdSnippet.RawTimestamp)
	return createdSnippet, nil
}
