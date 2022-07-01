package model

import (
	"bytes"
	"net/url"
)

type Route struct {
	Path   string
	Method string

	Input  Input
	Output Output
}

type Input struct {
	Headers    []string
	Body       bytes.Buffer
	Parameters []string
}

type Output struct {
	Headers    url.Values
	StatusCode int
	Body       bytes.Buffer
}
