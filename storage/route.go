package storage

import (
	"bytes"
	"net/url"
	"strings"

	"github.com/victorhsb/go-mocker/model"
)

type Route struct {
	Path   string `json:"path" yaml:"-"`
	Method string `json:"method" yaml:"-"`

	Input  Input  `json:"input" yaml:"input"`
	Output Output `json:"output" yaml:"output"`
}

type Input struct {
	Headers    []string `json:"headers" yaml:"headers"`
	Body       string   `json:"body" yaml:"body"`
	Parameters []string `json:"parameters" yaml:"parameters"`
}

type Output struct {
	Headers    map[string]string `json:"headers" yaml:"headers"`
	StatusCode int               `json:"statusCode" yaml:"statusCode"`
	Body       string            `json:"body" yaml:"body"`
}

func (r *Route) ToModel() (*model.Route, error) {
	out := bytes.NewBuffer([]byte(r.Output.Body))
	in := bytes.NewBuffer([]byte(r.Input.Body))

	route := &model.Route{
		Path:   r.Path,
		Method: r.Method,
		Input: model.Input{
			Body:       *in,
			Parameters: r.Input.Parameters,
			Headers:    r.Input.Headers,
		},
		Output: model.Output{
			StatusCode: r.Output.StatusCode,
			Body:       *out,
		},
	}

	route.Output.Headers = url.Values{}
	for k, v := range r.Output.Headers {
		route.Output.Headers.Add(k, v)
	}

	return route, nil
}

func ToStorage(model *model.Route) *Route {
	if model == nil {
		return nil
	}

	route := &Route{
		Path:   model.Path,
		Method: model.Method,
		Input: Input{
			Headers:    model.Input.Headers,
			Body:       model.Input.Body.String(),
			Parameters: model.Input.Parameters,
		},
		Output: Output{
			StatusCode: model.Output.StatusCode,
			Body:       model.Output.Body.String(),
			Headers:    map[string]string{},
		},
	}

	for k, v := range model.Output.Headers {
		route.Output.Headers[k] = strings.Join(v, ",")
	}

	return route
}
