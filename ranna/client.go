package ranna

import "github.com/valyala/fasthttp"

// Client represents a ranna client
type Client struct {
	instance string
	http     *fasthttp.Client
}

// NewClient creates a new ranna client
func NewClient(instance string) *Client {
	return &Client{
		instance: instance,
		http: &fasthttp.Client{
			Name: "lus/ranna-go",
		},
	}
}
